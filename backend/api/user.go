package api

import (
	"atm/model"
	"atm/response"
	"atm/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	service service.UserService
}

func NewUserAPI() UserAPI {
	return UserAPI{service.NewUserService()}
}

func (api UserAPI) GetInfo(ctx *gin.Context) {
	uid, err := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	if err != nil {
		fmt.Println("[ERROR] strconv.ParseInt: ", err)
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return 
	}
	userInfo, errCode := api.service.GetInfo(uid)
	response.Result(errCode, userInfo, ctx)
}

func (api UserAPI) Delete(ctx *gin.Context) {
	var params model.UserDeleteParams
	fmt.Println(ctx.Params)
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return 
	}
	errCode := api.service.Delete(uid, &params)
	response.Result(errCode, nil, ctx)
}

func (api UserAPI) GetVerifyCode(ctx *gin.Context) {
	var params model.UserGetVerifyCodeParams
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return 
	}
	errCode := api.service.GetVerifyCode(&params)
	response.Result(errCode, nil, ctx)
}

func (api UserAPI) Login(ctx *gin.Context) {
	var params model.UserLoginParams
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return 
	}
	token, errCode := api.service.Login(&params)
	if token == nil {
		response.Result(errCode, nil, ctx)
		return
	}
	response.Result(errCode, token, ctx)
}

func (api UserAPI) Register(ctx *gin.Context) {
	var params model.UserRegisterParams
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return 
	}
	errCode := api.service.Register(&params)
	response.Result(errCode, nil, ctx)
}

func (api UserAPI) ForgetPass(ctx *gin.Context) {
	var params model.UserForgetPassParams
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return 
	}
	errCode := api.service.ForgetPass(&params)
	response.Result(errCode, nil, ctx)
}
