package application

import (
	"context"
	"fmt"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain"
	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
)

type (
	AccumulateRewardWrapper interface {
		accumulateReward
	}

	accumulateReward interface {
		AccumulateReward(data *domain.RewardsDto) (*domain.RewardsAccumulateResponse, error)
	}
)

type accumulate struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccumalateRewards(ctx context.Context, svcCtx *svc.ServiceContext) AccumulateRewardWrapper {
	return &accumulate{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (a *accumulate) AccumulateReward(data *domain.RewardsDto) (*domain.RewardsAccumulateResponse, error) {
	// Validate if campaign is active
	// TOOD: validar y pasarlo mas a siemple responsability
	campaign, err := a.svcCtx.Campaings.CampaignsRepositories.FindOneByBranchId(data.BranchID)
	fmt.Println("#CAMPAIGN:", campaign)
	if err != nil {
		return nil, err
	}

	rewardEarned := 0
	rewardType := "points"

	if campaign != nil {
		if campaign.BonusType == "double" {
			fmt.Println("--- DOUBLE ---")
			rewardEarned = int(data.PurchaseAmount) * 2
		} else if campaign.BonusType == "percentage" {
			fmt.Println("--- PERCENTAGE ---")
			rewardEarned = int(float32(data.PurchaseAmount) * (1 + campaign.BonusValue))
		}
	} else {
		rewardEarned = int(data.PurchaseAmount) / 1000
	}

	newData := &models.Rewards{
		UserID:      data.UserID,
		BranchID:    data.BranchID,
		RewardType:  rewardType,
		RewardValue: rewardEarned,
	}

	res, err := a.svcCtx.Rewards.RewardsRepositories.CreateRewards(newData)
	if err != nil {
		return nil, err
	}

	fmt.Println("#RES:", res.Id)
	return &domain.RewardsAccumulateResponse{
		RewardsEarned:   rewardEarned,
		RewardsType:     rewardType,
		CampaignApplied: campaign != nil,
	}, nil
}
