package update

import "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"

type (
	Update interface {
		Update(campaignId string, data UpdateCampaignDto) (res *models.Campaigns, err error)
	}
)
