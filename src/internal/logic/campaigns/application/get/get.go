package get

import (
	"context"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
)

type get struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCampeings(ctx context.Context, svcCtx *svc.ServiceContext) GetCampaignsWrapper {
	return &get{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (g *get) GetAll(branchId string) (res []*models.Campaigns, err error) {
	if branchId == "" {
		res, err = g.svcCtx.Campaings.CampaignsRepositories.FindAll()
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	res, err = g.svcCtx.Campaings.CampaignsRepositories.FindAllByBranchId(branchId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
