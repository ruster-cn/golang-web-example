package requestid

import (
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	middleware_lib "github.com/op-server/pkg/server/middleware/lib"
)

// HeaderXRequestID X-Request-ID header
const HeaderXRequestID = "X-Request-ID"

// Middleware middleware which add X-Request-ID header in the http request when not exist
func Middleware(skippers ...middleware_lib.Skipper) func(ctx *gin.Context) {
	return middleware_lib.New(func(ctx *gin.Context) {
		rid := ctx.Request.Header.Get(HeaderXRequestID)
		if rid == "" {
			rid = uuid.New().String()
			ctx.Request.Header.Set(HeaderXRequestID, rid)
		}
		ctx.Writer.Header().Set(HeaderXRequestID, rid)
		ctx.Next()
	}, skippers...)
}
