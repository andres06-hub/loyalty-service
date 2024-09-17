package handler

import (
	"net/http"
	"strings"

	docs "github.com/andres06-hub/loyalty-service/src/docs"
	campaigns "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/infrastructure"
	rewards "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/infrastructure"
	"github.com/andres06-hub/loyalty-service/src/internal/svc" // Swagger handler
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zeromicro/go-zero/rest"
)

const apiPrefix = "/api"

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	var routes []rest.Route
	routes = append(routes, campaigns.GetCampaignsRoutes(ctx)...)
	routes = append(routes, rewards.GetRewardsRoutes(ctx)...)

	docs.SwaggerInfo.BasePath = apiPrefix

	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/swagger/*any",
				Handler: swaggerHandler,
			},
		},
	)

	for _, route := range routes {
		if !strings.HasPrefix(route.Path, "/swagger") {
			route.Path = apiPrefix + route.Path
		}
		server.AddRoute(route)
	}
}
