package handler

import (
	"fmt"
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
		httpSwagger.URL("/swagger/swagger.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)

	dir := http.Dir("/home/as06/data/code/golang/loyalty-service/src/docs")
	fmt.Println("DIR::::", dir)

	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/swagger/{any:*}",
			Handler: http.StripPrefix("/swagger", http.FileServer(dir)).ServeHTTP,
		},
	)

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/swagger-ui",
		Handler: swaggerHandler,
	})

	for _, route := range routes {
		if !strings.HasPrefix(route.Path, "/swagger") {
			route.Path = apiPrefix + route.Path
		}
		server.AddRoute(route)
	}
}

func test() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("TEST")

		w.Write([]byte("TEST"))
	})
}
