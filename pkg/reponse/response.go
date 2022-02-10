package reponse

import (
	"github.com/gin-gonic/gin"
	code "github.com/zxm1124/component-base/pkg/code"
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

// ToResponseList 响应列表信息
func (resp *Response) ToResponseList(listMeta meta.ListMeta, items []interface{}) {
	resp.Ctx.JSON(http.StatusOK, gin.H{
		"items": items,
		"pager": listMeta,
	})
}

// ToErrorResponse 响应错误信息
func (resp *Response) ToErrorResponse(err code.ErrorCode) {
	resp := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}
	if len(err.Details()) > 0 {
		resp["details"] = err.Details()
	}
	r.Ctx.JSON(err.StatusCode(), resp)
}
