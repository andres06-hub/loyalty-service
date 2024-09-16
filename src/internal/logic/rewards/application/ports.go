package application

import "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain"

type (
	AccumulateRewardWrapper interface {
		accumulateReward
	}

	accumulateReward interface {
		AccumulateReward(data *domain.RewardsDto) (*domain.RewardsAccumulateResponse, error)
	}
)

type (
	RedeemRewardWrapper interface {
		redeemReward
	}

	redeemReward interface {
		RedeemReward(data *domain.RewardsRedeemDto) (string, error)
	}
)
