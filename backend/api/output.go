package api

import "github.com/gin-gonic/gin"

type OutputAPI struct {
}

func NewOutputAPI() OutputAPI {
	return OutputAPI{}
}

func (api OutputAPI) GetList(ctx *gin.Context) {

}

func (api OutputAPI) GetInfo(ctx *gin.Context) {

}

func (api OutputAPI) Delete(ctx *gin.Context) {

}

func (api OutputAPI) ExportList(ctx *gin.Context) {

}

func (api OutputAPI) CreateInfo(ctx *gin.Context) {

}

func (api OutputAPI) UpdateInfo(ctx *gin.Context) {

}
