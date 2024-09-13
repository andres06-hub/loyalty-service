package svc

import (
	"github.com/andres06-hub/loyalty-service/src/internal/shared/http/parsers"
	httpResponse "github.com/andres06-hub/loyalty-service/src/internal/shared/http/response"
)

type Http struct {
	Parsers   *Parsers
	Responses *Responses
}

type Parsers struct {
	Body *parsers.BodyParser
}

type Responses struct {
	Success *httpResponse.APIResponse
	Error   *httpResponse.APIResponse
}

func GetHttp() *Http {
	return &Http{
		Parsers: &Parsers{
			Body: &parsers.BodyParser{},
		},
		Responses: &Responses{
			Success: httpResponse.NewAPISuccesResponse(),
			Error:   httpResponse.NewAPIErrorResponse(),
		},
	}
}
