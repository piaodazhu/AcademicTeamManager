package api

import (
	"atm/response"
	"atm/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SummaryAPI struct {
	service service.SummaryService
}

func NewSummaryAPI() SummaryAPI {
	return SummaryAPI{service.NewSummaryService()}
}

func (api SummaryAPI) Get(ctx *gin.Context) {
	uid, err1 := strconv.ParseInt(ctx.Request.Header.Get("uid"), 10, 64)
	if err1 != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	sum := api.service.Summary(uid)
	response.Result(response.ErrCodeSuccess, sum, ctx)
}
