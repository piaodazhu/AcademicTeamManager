package service

import (
	"atm/dao"
	"atm/model"
)

type SummaryService struct {
	dao dao.SummaryDao
}

func NewSummaryService() SummaryService {
	return SummaryService{dao.NewSummaryDao()}
}

func (service SummaryService) Summary(uid int64) *model.Summary {
	ms := service.dao.Count(uid)
	ms.StudentPan = service.dao.AnalyzeStudent(uid)
	ms.ProjectPan = service.dao.AnalyzeProject(uid)
	ms.OutputPan = service.dao.AnalyzeOutput(uid)
	return ms
}
