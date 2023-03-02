package api

import (
	"atm/model"
	"atm/response"
	"atm/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OutputAPI struct {
	service service.OutputService
}

func NewOutputAPI() OutputAPI {
	return OutputAPI{service.NewOutputService()}
}

func (api OutputAPI) GetList(ctx *gin.Context) {
	var params model.OutputQueryParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	outputList, rows, errCode := api.service.GetList(uid, &params)
	response.PageResult(errCode, outputList, rows, ctx)
}

func (api OutputAPI) GetInfo(ctx *gin.Context) {
	var params model.OutputQueryParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	outputInfo, errCode := api.service.GetInfo(&params)
	response.Result(errCode, outputInfo, ctx)
}

func (api OutputAPI) Delete(ctx *gin.Context) {
	var params model.OutputDeleteParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Delete(&params)
	response.Result(errCode, nil, ctx)
}

func (api OutputAPI) ExportList(ctx *gin.Context) {
	uid, err := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	file, errCode := api.service.Export(uid)
	if len(file) >= 0 && errCode != 0 {
		response.Result(errCode, nil, ctx)
		return
	}
	ctx.File(file)
}

func (api OutputAPI) CreateInfo(ctx *gin.Context) {
	var params model.OutputCreateParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)

	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		fmt.Println(err2)
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Create(uid, &params)
	response.Result(errCode, nil, ctx)
}

func (api OutputAPI) UpdateInfo(ctx *gin.Context) {
	var params model.OutputUpdateParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Update(&params)
	response.Result(errCode, nil, ctx)
}
