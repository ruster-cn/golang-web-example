package resp

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// WithJSONResp json格式返回
func WithJSONResp(f func(c *gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		data, err := f(c)
		if data == nil {
			data = make(map[string]interface{})
		}

		switch e := err.(type) {
		case nil:
			success := NewCodeError(Success)
			c.JSON(http.StatusOK, response{Code: success.Code(), Msg: success.Error(), Data: data})
		case *BaseError:
			c.JSON(http.StatusOK, response{Code: e.Code(), Msg: e.Error(), Data: data})
		default:
			switch err {
			case sql.ErrNoRows:
				//
			default:
				serverError := NewCodeError(ServerError)
				c.JSON(http.StatusOK, response{Code: serverError.Code(), Msg: serverError.Error(), Data: data})
			}
		}
	}
}
