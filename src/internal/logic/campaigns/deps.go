package campaigns

import (
	campaignPort "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/repositories"
	campaignsRepository "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/infrastructure/repositories"
	"gorm.io/gorm"
)

type CampaignsDependencies struct {
	CampaignsRepositories campaignPort.CampaignsRepository
}

func GetCampaignsDependencies(conn *gorm.DB) *CampaignsDependencies {
	return &CampaignsDependencies{
		CampaignsRepositories: campaignsRepository.NewCampaignsRepository(conn),
	}
}
