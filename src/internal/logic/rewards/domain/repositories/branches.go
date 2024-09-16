package repositories

import "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"

type BranchesRepository interface {
	FindOneById(id string) (res *models.Branches, err error)
}
