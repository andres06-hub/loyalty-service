package repositories

import (
	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"
)

type CampaignsRepository interface {
	FindAll() (res []*models.Campaigns, err error)
	FindAllByBranchId(branchId string) (res []*models.Campaigns, err error)
	FindOneByBranchId(branchId string) (res *models.Campaigns, err error)
}
