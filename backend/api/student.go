package api

import "github.com/gin-gonic/gin"

type StudentAPI struct {
}

func NewStudentAPI() StudentAPI {
	return StudentAPI{}
}

func (api StudentAPI) GetList(ctx *gin.Context) {

}

func (api StudentAPI) GetOption(ctx *gin.Context) {

}

func (api StudentAPI) GetInfo(ctx *gin.Context) {

}

func (api StudentAPI) Delete(ctx *gin.Context) {

}

func (api StudentAPI) ExportList(ctx *gin.Context) {

}

func (api StudentAPI) CreateInfo(ctx *gin.Context) {

}

func (api StudentAPI) UpdateInfo(ctx *gin.Context) {

}

func (api StudentAPI) SendEmail(ctx *gin.Context) {

}
