package request

import (
	"github.com/gin-gonic/gin"
	resp "go-web-core/core/util"
	"net/http"
	"reflect"
)

type Requester interface {
	Request(ctx *gin.Context) (*resp.Resp, error)
}

// 返回参数绑定的高阶函数
func Handle(r Requester) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 创建一个新的 Requester, 避免将上一次的参数带到下一次当中
		if reflect.TypeOf(r).Kind() != reflect.Ptr {
			panic("must be a pointer")
		}

		req := reflect.New(reflect.ValueOf(r).Elem().Type()).Interface().(Requester)
		r, err := request(req, ctx)

		// TODO ERROR 处理
		if err != nil {
			resp.Error(ctx, 400, "服务器错误")
			return
		}
		ctx.JSON(http.StatusOK, r)
	}
}

// request 参数绑定
func request(r Requester, ctx *gin.Context) (*resp.Resp, error) {
	if err := ctx.ShouldBind(r); err != nil {
		return nil, err
	}
	return r.Request(ctx)
}
