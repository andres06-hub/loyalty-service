package repositories

import (
	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"
	repositories "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/repositories"
	"gorm.io/gorm"
)

type BranchesRepository struct {
	db *gorm.DB
}

func NewBranchesRepository(dbConnection *gorm.DB) repositories.BranchesRepository {
	return &BranchesRepository{
		db: dbConnection,
	}
}

func (r *BranchesRepository) FindOneById(id string) (res *models.Branches, err error) {
	if err = r.db.Where("id = ?", id).First(&res).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return res, nil
}
