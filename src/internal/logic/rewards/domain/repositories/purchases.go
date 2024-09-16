package repositories

import "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"

type PurchasesRepository interface {
	Create(data *models.Purchases) (res *models.Purchases, err error)
}
