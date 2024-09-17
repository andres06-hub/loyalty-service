package handlers

import (
	"net/http"
	"strings"
	"time"

	update "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/application/update"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
)

//	@Tags			campaigns
//	@Router			/campaigns/:campaignId [put]
//	@Summary		Get all campaigns
//	@Description	Get all campaigns
//	@Accept			json
//	@Success		200			{object}	string	"Ok"
//	@Param			campaignId	path		string	true	"ID of the campaign to update"
//	@Failure		400			{object}	string	"Bad Request"
//	@Failure		401			{object}	string	"Unauthorized"
//	@Failure		503			{object}	string	"Service Unavailable"
func UpdateCampaignHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body update.UpdateCampaignDto

		err := ctx.Http.Parsers.Body.Parse(r, &body)
		if err != nil {
			ctx.Http.Responses.Error.
				WithTimestamp(int64(time.Now().Unix())).
				WithMessage(err.Error()).
				WithCode(http.StatusBadRequest).
				Build(w)
			return
		}

		campaignId := strings.Split(r.URL.String(), "/")[3]

		l := update.NewUpdate(r.Context(), ctx)
		res, err := l.Update(campaignId, body)
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
			SetData(res).
			WithMessage("campaign update").
			Build(w)
	}
}
