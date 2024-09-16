package repositories

import "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/models"

type RedemptionTransactionRepository interface {
	Create(data *models.RedemptionTransactions) (res *models.RedemptionTransactions, err error)
}
