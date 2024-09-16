package application

import (
	"context"
	"errors"
	"time"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain"
	models "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
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

func (r *reedem) RedeemReward(data *domain.RewardsRedeemDto) (*domain.RewardRedemptionResponse, error) {
	reward, err := r.svcCtx.Rewards.RewardsRepositories.FindOne(data.UserID, data.BranchID, data.RewardType)
	if err != nil {
		return nil, err
	}
	data.RewardValue /= 1000

	if reward.RewardValue < data.RewardValue {
		return nil, errors.New("insufficient reward balance")
	}

	// Deduct the user's balance
	reward.RewardValue -= data.RewardValue
	_, err = r.svcCtx.Rewards.RewardsRepositories.Update(reward)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &domain.RewardRedemptionResponse{
		RewardType:     data.RewardType,
		RewardRedeemed: data.RewardValue,
		CurrentReward:  reward.RewardValue,
	}, nil
}
