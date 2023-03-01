package api

import "github.com/gin-gonic/gin"

type UtilAPI struct {
}

func NewUtilAPI() UtilAPI {
	return UtilAPI{}
}

func (api UtilAPI) InitDB(ctx *gin.Context) {
	
}

func (api UtilAPI) RemoveFile(ctx *gin.Context) {

}

func (api UtilAPI) UploadFile(ctx *gin.Context) {

}
