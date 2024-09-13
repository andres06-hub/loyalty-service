package parsers

// import (
// 	"net/http"

// 	"github.com/pasztorpisti/qs"
// )

// type QueryParser struct{}

// func Query() *QueryParser {
// 	return &QueryParser{}
// }

// func (*QueryParser) Parse(ctx *http.Request, v any) error {
// 	query := ctx.URL.Query().Encode()
// 	return qs.Unmarshal(v, string(query))
// }
