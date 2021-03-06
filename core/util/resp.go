package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Resp 返回
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Error 错误返回
func Error(ctx *gin.Context, code int, message string, data ...interface{}) {
	resp(ctx, code, message, data...)
}

// Success 成功返回
func Success(ctx *gin.Context, message string, data ...interface{}) {
	resp(ctx, 0, message, data...)
}

func resp(ctx *gin.Context, code int, message string, data ...interface{}) {
	resp := Resp{
		Code: code,
		Msg:  message,
		Data: data,
	}

	// [args1, args2, args3]
	// [args1] 只有一个参数原数据返回
	if len(data) == 1 {
		resp.Data = data[0]
	}

	// 根据code返回Http状态码

	ctx.JSON(http.StatusOK, resp)

}
