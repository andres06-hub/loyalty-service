package middlewares

import (
	"net/http"
	"os"
	"strings"

	httpResponse "github.com/andres06-hub/loyalty-service/src/internal/shared/http/response"
)

type MiddlewareOrigin struct{}

func NewMiddlewareValidationPermissions() *MiddlewareOrigin {
	return &MiddlewareOrigin{}
}

func (mpv *MiddlewareOrigin) IPValidatorMiddleware() func(next http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedIPs := []string{
				"127.0.0.1:8888",
			}

			clientIP := r.RemoteAddr
			isAllowedIP := false
			for _, ip := range allowedIPs {
				if clientIP == ip {
					isAllowedIP = true
					break
				}
			}
			if !isAllowedIP {
				res := httpResponse.NewAPIErrorResponse().
					WithMessage("Access from this IP is not allowed").
					WithCode(http.StatusForbidden)
				httpResponse.JSONResponse(w, res)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (mpv *MiddlewareOrigin) HostValidatorMiddleware() func(next http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if !mpv.isHostAllowed(r.Host) {
				res := httpResponse.NewAPIErrorResponse().
					WithMessage("Host not allowed").
					WithCode(http.StatusForbidden)
				httpResponse.JSONResponse(w, res)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (mpv *MiddlewareOrigin) isHostAllowed(host string) bool {
	allowedHost := strings.Split(os.Getenv("ALLOWED_HOSTS"), ",")

	for _, h := range allowedHost {
		if host == h {
			return true
		}
	}

	return false
}
