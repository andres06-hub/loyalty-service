package rewards

import (
	rwPort "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain/repositories"
	repositories "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/infrastructure/repositories"
	"gorm.io/gorm"
)

type RewardsDependencies struct {
	RewardsRepositories               rwPort.RewardsRepository
	RedemptionTransactionRepositories rwPort.RedemptionTransactionRepository
	PurchasesRepository               rwPort.PurchasesRepository
	BranchesRepository                rwPort.BranchesRepository
}

func GetRewardsDependencies(conn *gorm.DB) *RewardsDependencies {
	return &RewardsDependencies{
		RewardsRepositories:               repositories.NewRewardsRepository(conn),
		RedemptionTransactionRepositories: repositories.NewRedemptionTransactionRepository(conn),
		PurchasesRepository:               repositories.NewPurchasesRepository(conn),
		BranchesRepository:                repositories.NewBranchesRepository(conn),
	}
}
