package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"
	rwRpt "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RewardsRepository struct {
	db *gorm.DB
}

func NewRewardsRepository(dbConnection *gorm.DB) rwRpt.RewardsRepository {
	return &RewardsRepository{
		db: dbConnection,
	}
}

func (r *RewardsRepository) FindOneByUserIdAndBranchId(userID string, branchID string) (res *models.Rewards, err error) {
	if err = r.db.Where("user_id = ? AND branch_id = ?", userID, branchID).First(&res).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (r *RewardsRepository) CreateRewards(data *models.Rewards) (res *models.Rewards, err error) {
	data.Id = uuid.New().String()
	data.CreatedAt = time.Now().Format("2006-01-02 15:04:05.000 -0700")

	err = r.db.Create(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error creating reward")
	}

	return data, nil
}

func (r *RewardsRepository) FindOneById(id string) (res *models.Rewards, err error) {
	err = r.db.Raw("SELECT * FROM rewards WHERE id = ?", id).Scan(&res).Error
	if err != nil {
		return nil, fmt.Errorf("error finding reward")
	}

	if res == nil {
		return nil, fmt.Errorf("reward not found")
	}

	return res, nil
}

func (r *RewardsRepository) FindOne(userID, branchID, rewardType string) (res *models.Rewards, err error) {
	if err = r.db.Where("user_id = ? AND branch_id = ? AND reward_type = ?", userID, branchID, rewardType).First(&res).Error; err != nil {
		return nil, errors.New("reward balance not found")
	}
	return res, nil
}

func (r *RewardsRepository) Update(data *models.Rewards) (res *models.Rewards, err error) {
	err = r.db.Save(&data).Error
	if err != nil {
		return nil, fmt.Errorf("error updating reward")
	}

	return data, nil
}
