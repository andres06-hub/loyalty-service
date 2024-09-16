package application

import (
	"context"
	"database/sql"
	"errors"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain"
	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"gorm.io/gorm"
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
	_, err := a.svcCtx.Rewards.BranchesRepository.FindOneById(data.BranchID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("branch not found")
		}
		return nil, err
	}

	// TOOD: validar y pasarlo mas a siemple responsability
	campaign, err := a.svcCtx.Campaings.CampaignsRepositories.FindOneByBranchId(data.BranchID)
	if err != nil {
		return nil, err
	}

	// currentDate := time.Now().Format("2006-01-02")
	currentDate := "2024-05-16"

	rewardEarned := 0.0
	rewardType := "points"
	campaignId := sql.NullString{Valid: false}

	if campaign != nil {
		// Validate if campaign is active
		if currentDate < campaign.StartDate || currentDate > campaign.EndDate {
			return nil, errors.New("campaign not active")
		}
		campaignId.Valid = true
		campaignId.String = campaign.Id
		if campaign.BonusType == "double" {
			rewardEarned = (float64(data.PurchaseAmount) / 1000) * 2
		} else if campaign.BonusType == "percentage" {
			rewardEarned = (float64(data.PurchaseAmount) / 1000) * (1 + campaign.BonusValue)
		}
	} else {
		rewardEarned = float64(data.PurchaseAmount) / 1000
	}

	newData := &models.Rewards{
		UserID:      data.UserID,
		BranchID:    data.BranchID,
		RewardType:  rewardType,
		RewardValue: rewardEarned,
	}

	res, err := a.svcCtx.Rewards.RewardsRepositories.FindOneByUserIdAndBranchId(data.UserID, data.BranchID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			_, err = a.svcCtx.Rewards.RewardsRepositories.CreateRewards(newData)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if res != nil {
		res.RewardValue += rewardEarned
		_, err = a.svcCtx.Rewards.RewardsRepositories.Update(res)
		if err != nil {
			return nil, err
		}
	}

	purchaseData := &models.Purchases{
		UserId:         data.UserID,
		BranchId:       data.BranchID,
		PurchaseAmount: float64(data.PurchaseAmount),
		RewardEarned:   rewardEarned,
		RewardType:     rewardType,
		CampaignId:     campaignId,
	}

	_, err = a.svcCtx.Rewards.PurchasesRepository.Create(purchaseData)
	if err != nil {
		return nil, err
	}

	return &domain.RewardsAccumulateResponse{
		RewardsEarned:   rewardEarned,
		RewardsType:     rewardType,
		CampaignApplied: campaign != nil,
	}, nil
}
