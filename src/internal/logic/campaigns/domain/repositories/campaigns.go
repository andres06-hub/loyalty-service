package repositories

import (
	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"
)

type CampaignsRepository interface {
	FindAll() (res []*models.CampaignModel, err error)
	FindAllByBranchId(branchId string) (res []*models.CampaignModel, err error)
}
