package repositories

import "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"

type RewardsRepository interface {
	CreateRewards(data *models.Rewards) (res *models.Rewards, err error)
	FindOneById(id string) (res *models.Rewards, err error)
	FindOne(userID, branchID, rewardType string) (res *models.Rewards, err error)
	FindOneByUserIdAndBranchId(userID, branchID string) (res *models.Rewards, err error)
	Update(data *models.Rewards) (res *models.Rewards, err error)
}
