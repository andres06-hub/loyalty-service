package logic

import (
	h "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/infrastructure/handlers"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func GetCampaignsRoutes(ctx *svc.ServiceContext) []rest.Route {
	return []rest.Route{
		{
			Method: "GET",
			Path:   "/campaigns",
			Handler: ctx.Middlewares.Host.HostValidatorMiddleware()(
				h.GetCampaignsHandler(ctx),
			),
		},
	}
}
