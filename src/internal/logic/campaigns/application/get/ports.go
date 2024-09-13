package get

import "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"

type (
	GetCampaignsWrapper interface {
		getCampaigns
	}

	getCampaigns interface {
		GetAll(branchId string) (res []*models.CampaignModel, err error)
	}
)
