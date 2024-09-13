package svc

import (
	"github.com/andres06-hub/loyalty-service/src/internal/config"
	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns"
	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	ConnDb      *gorm.DB
	Middlewares *Middleware
	Campaings   *campaigns.CampaignsDependencies
	Rewards     *rewards.RewardsDependencies
	Http        *Http
}

func NewServiceContext(c config.Config, conn *gorm.DB) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ConnDb:      conn,
		Middlewares: GetMiddleware(),
		Campaings:   campaigns.GetCampaignsDependencies(conn),
		Rewards:     rewards.GetRewardsDependencies(conn),
		Http:        GetHttp(),
	}
}
