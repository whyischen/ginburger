package ginburger

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode    = 0
	SuccessMessage = "OK"

	ErrorCode    = 10007
	ErrorMessage = "error"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildSuccessResponse(c *gin.Context) *Response {
	resp := &Response{
		Code:    SuccessCode,
		Message: SuccessMessage,
		Data:    map[string]interface{}{},
	}

	c.JSON(http.StatusOK, resp)
	return resp
}

func BuildErrorResponse(c *gin.Context) *Response {
	resp := &Response{
		Code:    ErrorCode,
		Message: ErrorMessage,
		Data:    map[string]interface{}{},
	}

	c.JSON(http.StatusOK, resp)
	return resp
}

func (receiver *Response) SetData(data interface{}) *Response {
	receiver.Data = data
	return receiver
}
