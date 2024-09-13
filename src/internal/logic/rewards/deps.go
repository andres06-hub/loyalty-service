package rewards

import (
	rwPort "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/repositories"
	rewardsRepository "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/infrastructure/repositories"
	"gorm.io/gorm"
)

type RewardsDependencies struct {
	RewardsRepositories rwPort.RewardsRepository
}

func GetRewardsDependencies(conn *gorm.DB) *RewardsDependencies {
	return &RewardsDependencies{
		RewardsRepositories: rewardsRepository.NewRewardsRepository(conn),
	}
}
