package v1

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	resp "go-web-core/core/util"
	"log"
)

type CreateUser struct {
	UserName string
}

func (u *CreateUser) Request(ctx *gin.Context) (*resp.Resp, error) {

	// service

	log.Println("username1", u.UserName)

	return &resp.Resp{
		Data: base64.StdEncoding.EncodeToString([]byte("asdfa")),
	}, errors.New("出错")
}
