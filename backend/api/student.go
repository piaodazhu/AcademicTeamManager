package api

import (
	"atm/model"
	"atm/response"
	"atm/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentAPI struct {
	service service.StudentService
}

func NewStudentAPI() StudentAPI {
	return StudentAPI{service.NewStudentService()}
}

func (api StudentAPI) GetList(ctx *gin.Context) {
	var params model.StudentQueryParams
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	studentList, rows, errCode := api.service.GetList(uid, &params)
	response.PageResult(errCode, studentList, rows, ctx)
}

func (api StudentAPI) GetOption(ctx *gin.Context) {
	uid, err := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	studentOption, errCode := api.service.GetOption(uid)
	response.Result(errCode, studentOption, ctx)
}

func (api StudentAPI) GetInfo(ctx *gin.Context) {
	var params model.StudentQueryParams
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	studentInfo, errCode := api.service.GetInfo(&params)
	response.Result(errCode, studentInfo, ctx)
}

func (api StudentAPI) Delete(ctx *gin.Context) {
	var params model.StudentDeleteParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Delete(&params)
	response.Result(errCode, nil, ctx)
}

func (api StudentAPI) ExportList(ctx *gin.Context) {
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

func (api StudentAPI) CreateInfo(ctx *gin.Context) {
	var params model.StudentCreateParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Create(uid, &params)
	response.Result(errCode, nil, ctx)
}

func (api StudentAPI) UpdateInfo(ctx *gin.Context) {
	var params model.StudentUpdateParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Update(&params)
	response.Result(errCode, nil, ctx)
}

func (api StudentAPI) SendEmail(ctx *gin.Context) {
	var params model.StudentSendMailParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.SendMail(uid, &params)
	response.Result(errCode, nil, ctx)
}
