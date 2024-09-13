package svc

import (
	mdd "github.com/andres06-hub/loyalty-service/src/internal/middlewares"
)

type Middleware struct {
	Host mdd.MiddlewareOrigin
}

func GetMiddleware() *Middleware {
	return &Middleware{
		Host: *mdd.NewMiddlewareValidationPermissions(),
	}
}
