package repositories

import (
	"time"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"
	repositories "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchasesRepository struct {
	db *gorm.DB
}

func NewPurchasesRepository(db *gorm.DB) repositories.PurchasesRepository {
	return &PurchasesRepository{
		db: db,
	}
}

func (p *PurchasesRepository) Create(data *models.Purchases) (res *models.Purchases, err error) {
	data.Id = uuid.New().String()
	data.CreatedAt = time.Now().Format("2006-01-02 15:04:05.000 -0700")

	err = p.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
