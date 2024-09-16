package application

import (
	"context"
	"errors"
	"time"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain"
	models "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"gorm.io/gorm"
)

type reedem struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedeemRewards(ctx context.Context, svcCtx *svc.ServiceContext) RedeemRewardWrapper {
	return &reedem{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (r *reedem) RedeemReward(data *domain.RewardsRedeemDto) (string, error) {
	reward, err := r.svcCtx.Rewards.RewardsRepositories.FindOne(data.UserID, data.BranchID, data.RewardType)
	if err != nil {
		return "", err
	}
	data.RewardValue /= 1000

	if reward.RewardValue < data.RewardValue {
		return "", errors.New("insufficient reward balance")
	}

	// Check if there are active campaigns for the branch office
	// now := time.Now().Format("2006-01-02 15:04:05.000 -0700")
	now := "2024-05-16"
	campaign, err := r.svcCtx.Campaings.CampaignsRepositories.FindOneByBranchIdAndDates(data.BranchID, now)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	if campaign != nil {
		// Aplicar el bono de la campaña si existe
		if campaign.BonusType == "double" {
			data.RewardValue /= 2
		} else if campaign.BonusType == "percentage" {
			data.RewardValue *= (1 + campaign.BonusValue)
		}
	}

	// Deduct the user's balance
	reward.RewardValue -= data.RewardValue
	_, err = r.svcCtx.Rewards.RewardsRepositories.Update(reward)
	if err != nil {
		return "", err
	}

	transaction := &models.RedemptionTransactions{
		UserID:      data.UserID,
		BranchID:    data.BranchID,
		RewardType:  data.RewardType,
		RewardValue: data.RewardValue,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05.000 -0700"),
	}

	_, err = r.svcCtx.Rewards.RedemptionTransactionRepositories.Create(transaction)
	if err != nil {
		return "", err
	}

	// TODO: remaining_points
	// Retornar el saldo restante
	// return fmt.Sprintf("Remaining %s: %f", data.RewardType, reward.RewardValue), nil
	return "", nil
}

type Response struct {
	RewardType         string  `json:"reward_type"`
	CurrentRewardValue float64 `json:"reward_value"`
}
