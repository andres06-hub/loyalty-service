package adapters

import (
	"net/http"

	get "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/application/get"
	httpResponse "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/infrastructure/http/response"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
)

type CampaignsHandler struct{}

func GetCampaignsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		branchId := r.URL.Query().Get("branchId")

		l := get.NewGetCampeings(r.Context(), ctx)
		rp, err := l.GetAll(branchId)
		if err != nil {
			res := httpResponse.NewAPIErrorResponse().
				WithMessage(err.Error()).
				WithCode(http.StatusInternalServerError)
			httpResponse.JSONResponse(w, res)
			return
		}

		res := httpResponse.NewAPISuccesResponse().
			SetData(rp).
			WithMessage("campaigns found")
		httpResponse.JSONResponse(w, res)
	}
}
