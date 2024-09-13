package svc

import (
	"github.com/andres06-hub/loyalty-service/src/internal/config"
	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	ConnDb      *gorm.DB
	Middlewares Middleware
	Campaings   campaigns.CampaignsDependencies
}

func NewServiceContext(c config.Config, conn *gorm.DB) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ConnDb:      conn,
		Middlewares: *GetMiddleware(),
		Campaings:   *campaigns.GetCampaignsDependencies(conn),
	}
}
