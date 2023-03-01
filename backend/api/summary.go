package api

import "github.com/gin-gonic/gin"

type SummaryAPI struct {

}

func NewSummaryAPI() SummaryAPI {
	return SummaryAPI{}
}

func (api SummaryAPI) Get(ctx *gin.Context) {
	
}