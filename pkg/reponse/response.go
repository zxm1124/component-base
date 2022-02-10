package reponse

import (
	"github.com/gin-gonic/gin"
	meta "github.com/zxm1124/component-base/pkg/meta/v1"
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

func (resp *Response) ToResponseList(listInterface meta.ListInterface) {
	resp.Ctx.JSON(http.StatusOK, gin.H{
		"items": listInterface, .
	})
}
