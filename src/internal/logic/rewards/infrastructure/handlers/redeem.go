package handlers

import (
	"fmt"
	"net/http"
	"time"

	rw "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/application"
	"github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/domain"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
)

// @Tags			rewards
// @Router		/rewards/redeem [post]
// @Summary		Get all campaigns
// @Description	redeem rewards
// @Accept			json
// @Success		200			{object}	string	"Ok"
// @Failure		400			{object}	string	"Bad Request"
// @Failure		401			{object}	string	"Unauthorized"
// @Failure		503			{object}	string	"Service Unavailable"
func RedeemRewardHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.RewardsRedeemDto{}

		err := ctx.Http.Parsers.Body.Parse(r, data)
		if err != nil {
			ctx.Http.Responses.Error.
				WithTimestamp(int64(time.Now().Unix())).
				WithMessage(err.Error()).
				WithCode(http.StatusBadGateway).
				Build(w)
			return
		}

		l := rw.NewRedeemRewards(r.Context(), ctx)
		res, err := l.RedeemReward(data)
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
			WithMessage(fmt.Sprintf("%s redeemd", data.RewardType)).
			Build(w)
	}
}
