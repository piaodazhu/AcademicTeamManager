package api

import (
	"atm/model"
	"atm/response"
	"atm/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectAPI struct {
	service service.ProjectService
}

func NewProjectAPI() ProjectAPI {
	return ProjectAPI{service.NewProjectService()}
}

func (api ProjectAPI) GetList(ctx *gin.Context) {
	var params model.ProjectQueryParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil || params.Page.PageNum <= 0 || params.Page.PageSize <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	projectList, rows, errCode := api.service.GetList(uid, &params)
	response.PageResult(errCode, projectList, rows, ctx)
}

func (api ProjectAPI) GetInfo(ctx *gin.Context) {
	var params model.ProjectQueryParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	projectInfo, errCode := api.service.GetInfo(&params)
	response.Result(errCode, projectInfo, ctx)

}

func (api ProjectAPI) UpdateOutputs(ctx *gin.Context) {
	var params model.ProjectOutputQueryParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	outputList, errCode := api.service.UpdateOutputs(uid, &params)
	response.Result(errCode, outputList, ctx)
}

func (api ProjectAPI) Delete(ctx *gin.Context) {
	var params model.ProjectDeleteParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Delete(&params)
	response.Result(errCode, nil, ctx)
}

func (api ProjectAPI) ExportList(ctx *gin.Context) {
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

func (api ProjectAPI) CreateInfo(ctx *gin.Context) {
	var params model.ProjectCreateParam
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	err2 := ctx.ShouldBind(&params)
	if err1 != nil || err2 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Create(uid, &params)
	response.Result(errCode, nil, ctx)
}

func (api ProjectAPI) UpdateInfo(ctx *gin.Context) {
	var params model.ProjectUpdateParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.Update(&params)
	response.Result(errCode, nil, ctx)
}
