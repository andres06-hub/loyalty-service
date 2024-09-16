package repositories

import (
	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"
)

type CampaignsRepository interface {
	FindAll() (res []*models.Campaigns, err error)
	FindOneById(id string) (res *models.Campaigns, err error)
	FindAllByBranchId(branchId string) (res []*models.Campaigns, err error)
	FindOneByBranchId(branchId string) (res *models.Campaigns, err error)
	FindOneByBranchIdAndDates(branchID, nowDate string) (res *models.Campaigns, err error)
	Update(data *models.Campaigns) (res *models.Campaigns, err error)
}
