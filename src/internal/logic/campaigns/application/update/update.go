package update

import (
	"context"
	"time"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
)

type update struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdate(ctx context.Context, svcCtx *svc.ServiceContext) Update {
	return &update{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *update) Update(campaignId string, data UpdateCampaignDto) (res *models.Campaigns, err error) {
	res, err = u.svcCtx.Campaings.CampaignsRepositories.FindOneById(campaignId)
	if err != nil {
		return nil, err
	}

	startDate, _ := time.Parse("2006-01-02", data.StartDate)
	endDate, _ := time.Parse("2006-01-02", data.EndDate)

	res.StartDate = startDate
	res.EndDate = endDate
	res.BonusType = data.BonusType
	res.BonusValue = data.BonusValue

	res, err = u.svcCtx.Campaings.CampaignsRepositories.Update(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
