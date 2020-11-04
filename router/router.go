package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-web-core/api/v1"
	"go-web-core/core/request"
	"log"
)

func init() {
	r := gin.Default()
	//u := &v1.User{}
	r.POST("user", request.Handle(&v1.CreateUser{}))
	go func() {
		err := r.Run(":8081")
		if err != nil {
			log.Println(err.Error())
		}
	}()

}
