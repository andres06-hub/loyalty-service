package repositories

import (
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
