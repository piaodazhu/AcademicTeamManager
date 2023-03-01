package api

import (
	"atm/response"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
}

func NewUserAPI() UserAPI {
	return UserAPI{}
}

func (api UserAPI) GetInfo(ctx *gin.Context) {
	response.Result(response.ErrCodeSuccess, "test", ctx)
}

func (api UserAPI) Delete(ctx *gin.Context) {

}

func (api UserAPI) GetVerifyCode(ctx *gin.Context) {

}

func (api UserAPI) Login(ctx *gin.Context) {

}

func (api UserAPI) Register(ctx *gin.Context) {

}

func (api UserAPI) ForgetPass(ctx *gin.Context) {

}
