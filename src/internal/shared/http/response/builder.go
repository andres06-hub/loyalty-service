package response

import (
	"net/http"
	"time"
)

// APIResponse is the response structure for the API
// swagger:model APIResponse
type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
	// Localization l.Localization `json:"localization"`
}

func NewAPISuccesResponse() *APIResponse {
	return &APIResponse{
		Success:   true,
		Code:      http.StatusOK,
		Timestamp: int64(time.Now().Unix()),
	}
}

func NewAPIErrorResponse() *APIResponse {
	return &APIResponse{
		Success:   false,
		Code:      http.StatusInternalServerError,
		Timestamp: int64(time.Now().Unix()),
		Data:      nil,
	}
}

func (a *APIResponse) SetData(data interface{}) *APIResponse {
	a.Data = data
	return a
}

func (a *APIResponse) WithMessage(message string) *APIResponse {
	a.Message = message
	return a
}

func (a *APIResponse) WithCode(code int) *APIResponse {
	a.Code = code
	return a
}

func (a *APIResponse) WithTimestamp(timestamp int64) *APIResponse {
	a.Timestamp = timestamp
	return a
}

func (a *APIResponse) WithSuccess(success bool) *APIResponse {
	a.Success = success
	return a
}

func (a *APIResponse) Build(ctx http.ResponseWriter) {
	JSONResponse(ctx, a)
}
