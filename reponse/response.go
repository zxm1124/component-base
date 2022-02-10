package reponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 封装响应实体
type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (resp *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	resp.Ctx.JSON(http.StatusOK, data)
}

func (resp *Response) ToResponseList()
