package handler

import (
	campaigns "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/infrastructure"
	rewards "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/infrastructure"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

const apiPrefix = "/api"

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	var routes []rest.Route
	routes = append(routes, campaigns.GetCampaignsRoutes(ctx)...)
	routes = append(routes, rewards.GetRewardsRoutes(ctx)...)

	for _, route := range routes {
		route.Path = apiPrefix + route.Path
		server.AddRoute(route)
	}
}
