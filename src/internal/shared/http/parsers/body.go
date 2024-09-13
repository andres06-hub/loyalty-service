package parsers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BodyParser struct{}

func Body() Parser {
	return &BodyParser{}
}

func (*BodyParser) Parse(ctx *http.Request, v interface{}) error {
	fmt.Println("#### BodyParser")
	return json.NewDecoder(ctx.Body).Decode(&v)
}
