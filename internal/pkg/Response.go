package pkg

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Ok(c *app.RequestContext, data interface{}) {
	resp := &Response{
		Code:    200,
		Success: true,
		Message: "操作成功",
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func Err(c *app.RequestContext) {
	resp := &Response{
		Code:    500,
		Success: false,
		Message: "操作失败",
	}
	c.JSON(http.StatusBadGateway, resp)
}

func ErrWithMessage(c *app.RequestContext, message string) {
	resp := &Response{
		Code:    500,
		Success: false,
		Message: message,
	}
	c.JSON(http.StatusBadGateway, resp)
}
