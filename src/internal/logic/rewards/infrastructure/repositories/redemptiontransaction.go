package repositories

import (
	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"
	rdRpt "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RedemptionTransactionRepository struct {
	db *gorm.DB
}

func NewRedemptionTransactionRepository(db *gorm.DB) rdRpt.RedemptionTransactionRepository {
	return &RedemptionTransactionRepository{
		db: db,
	}
}

func (r *RedemptionTransactionRepository) Create(data *models.RedemptionTransactions) (res *models.RedemptionTransactions, err error) {
	data.Id = uuid.New().String()

	err = r.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
