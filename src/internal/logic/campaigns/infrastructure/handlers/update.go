package handlers

import (
	"net/http"
	"strings"
	"time"

	update "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/application/update"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateCampaignHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body update.UpdateCampaignDto

		err := httpx.Parse(r, &body)
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
