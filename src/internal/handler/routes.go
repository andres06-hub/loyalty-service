package handler

import (
	campaigns "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/infrastructure"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

const apiPrefix = "/api"

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	var routes []rest.Route
	routes = append(routes, campaigns.GetCampaignsRoutes(ctx)...)

	for _, route := range routes {
		route.Path = apiPrefix + route.Path
		server.AddRoute(route)
	}
}
