package parsers

import (
	"encoding/json"
	"net/http"
)

type BodyParser struct{}

func Body() Parser {
	return &BodyParser{}
}

func (*BodyParser) Parse(ctx *http.Request, v interface{}) error {
	return json.NewDecoder(ctx.Body).Decode(&v)
}
