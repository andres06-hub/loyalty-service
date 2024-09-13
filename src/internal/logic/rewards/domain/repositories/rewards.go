package repositories

import "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"

type RewardsRepository interface {
	CreateRewards(data *models.Rewards) (res *models.Rewards, err error)
	FindOneById(id string) (res *models.Rewards, err error)
}
