package logic

import (
	h "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/infrastructure/handlers"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func GetRewardsRoutes(ctx *svc.ServiceContext) []rest.Route {
	return []rest.Route{
		{
			Method: "POST",
			Path:   "/rewards",
			Handler: ctx.Middlewares.Host.HostValidatorMiddleware()(
				h.AccumulateRewardHandler(ctx),
			),
		},
		{
			Method: "POST",
			Path:   "/rewards/redeem",
			Handler: ctx.Middlewares.Host.HostValidatorMiddleware()(
				h.RedeemRewardHandler(ctx),
			),
		},
	}
}
