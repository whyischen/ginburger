package ginburger

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode    = 0
	SuccessMessage = "OK"
)

type E map[string]interface{}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// default response
// {"code":0, "message": "OK", "data": {}}
func DefaultResp() *Response {
	resp := &Response{
		Code:    SuccessCode,
		Message: SuccessMessage,
		Data:    map[string]interface{}{},
	}
	return resp
}

func BuildCodeMsgResp(c *gin.Context, code int, msg string) {
	receiver := &Response{
		Code:    code,
		Message: msg,
		Data:    map[string]interface{}{},
	}

	receiver.Build(c)
}

func BuildCustomResp(c *gin.Context, code int, msg string, data interface{}) {
	receiver := &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}

	receiver.Build(c)
}

func (receiver *Response) Build(c *gin.Context) {
	c.JSON(http.StatusOK, receiver)
	c.Abort()
}
