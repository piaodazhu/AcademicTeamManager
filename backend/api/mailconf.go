package api

import (
	"atm/model"
	"atm/response"
	"atm/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MailConfAPI struct {
	service service.MailconfService
}

func NewMailConfAPI() MailConfAPI {
	return MailConfAPI{service.NewMailconfService()}
}

func (api MailConfAPI) CheckValidity(ctx *gin.Context) {
	uid, err := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Check(uid)
	response.Result(errCode, nil, ctx)
}

// func (api MailConfAPI) GetSwitch(ctx *gin.Context) {
// 	uid, err := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
// 	if err != nil {
// 		response.Result(response.ErrCodeParamInvalid, nil, ctx)
// 		return
// 	}
// 	errCode := api.service.Check(uid)
// 	response.Result(errCode, nil, ctx)
// }

func (api MailConfAPI) SetSwitch(ctx *gin.Context) {
	var params model.MailConfigStatusParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.UpdateStatus(uid, &params)
	response.Result(errCode, nil, ctx)
}

func (api MailConfAPI) GetInfo(ctx *gin.Context) {
	uid, err := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	mailConfig, errCode := api.service.GetInfo(int64(uid))
	response.Result(errCode, mailConfig, ctx)
}

func (api MailConfAPI) Delete(ctx *gin.Context) {
	var params model.MailConfigDeleteParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		log.Println(err)
		return
	}
	errCode := api.service.Delete(&params)
	response.Result(errCode, nil, ctx)
}

func (api MailConfAPI) Save(ctx *gin.Context) {
	var params model.MailConfigSaveParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Save(uid, &params)
	if errCode == 0 {
		mailConfig, errCode := api.service.GetInfo(uid)
		response.Result(errCode, mailConfig, ctx)
		return
	}
	response.Result(errCode, nil, ctx)
}
