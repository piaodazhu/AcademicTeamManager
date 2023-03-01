package api

import (
	"atm/model"
	"atm/response"
	"atm/service"

	"github.com/gin-gonic/gin"
)

type UtilAPI struct {
	service service.UtilService
}

func NewUtilAPI() UtilAPI {
	return UtilAPI{service.NewUtilService()}
}

func (api UtilAPI) InitDB(ctx *gin.Context) {
	errCode := api.service.InitDatabase()
	response.Result(errCode, nil, ctx)
}

func (api UtilAPI) RemoveFile(ctx *gin.Context) {
	var params model.FileParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := api.service.FileRemove(&params)
	response.Result(errCode, nil, ctx)
}

func (api UtilAPI) UploadFile(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	fileInfo, errCode := api.service.FileUpload(file)
	response.Result(errCode, fileInfo, ctx)
}
