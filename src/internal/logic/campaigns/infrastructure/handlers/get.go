package handlers

import (
	"net/http"
	"time"

	get "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/application/get"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
)

// @Tags			campaigns
// @Router			/campaigns [get]
// @Summary		Get all campaigns
// @Description	Get all campaigns
// @Accept			json
// @Success		200	{object}	string	"Ok"
// @Failure		400	{object}	string	"Bad Request"
// @Failure		401	{object}	string	"Unauthorized"
// @Failure		503	{object}	string	"Service Unavailable"
func GetCampaignsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		branchId := r.URL.Query().Get("branchId")

		l := get.NewGetCampeings(r.Context(), ctx)
		rp, err := l.GetAll(branchId)
		if err != nil {
			ctx.Http.Responses.Error.
				WithTimestamp(int64(time.Now().Unix())).
				WithMessage(err.Error()).
				WithCode(http.StatusInternalServerError).
				Build(w)
			return
		}

		ctx.Http.Responses.Success.
			WithTimestamp(int64(time.Now().Unix())).
			SetData(rp).
			WithMessage("campaigns found").
			Build(w)
	}
}
