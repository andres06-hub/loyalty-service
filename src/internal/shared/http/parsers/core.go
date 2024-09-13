package parsers

import "net/http"

type Parser interface {
	Parse(ctx *http.Request, v interface{}) error
}
