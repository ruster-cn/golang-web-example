package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/op-server/pkg/server/middleware/requestid"
)

// MiddleWares returns global middleware
func MiddleWares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		//TODO: add global middleware
		requestid.Middleware(),
	}
}
