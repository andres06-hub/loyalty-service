package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONResponse(ctx http.ResponseWriter, response *APIResponse) {
	ctx.WriteHeader(response.Code)
	ctx.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		fmt.Println("json encoding error", err)
		http.Error(ctx, err.Error(), http.StatusInternalServerError)
	}
}
