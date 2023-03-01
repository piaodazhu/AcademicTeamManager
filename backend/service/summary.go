package service

import (
	"atm/dao"
	"atm/model"
	"time"
)

type SummaryService struct {
	dao dao.SummaryDao
}

func NewSummaryService() SummaryService {
	return SummaryService{dao.NewSummaryDao()}
}

func (service SummaryService) Summary(uid int64, days int) model.Summary {
	ds := service.dao.Count(uid, days)
	now := time.Date(2023, 2, 1, 0, 0, 0, 0, time.Local).Unix()
	ds.Amount = make([]float64, days)
	ds.Date = make([]string, days)
	for i, dd := 0, days; i < days; i++ {
		day := now - (int64(dd) * 24 * 60 * 60)
		start, end := dayRange(day)
		ds.Amount[i] = service.dao.AmountSum(start, end, uid)
		ds.Date[i] = time.Unix(day, 0).Format("01-02")
		dd--
	}
	return ds
}

func dayRange(day int64) (int64, int64) {
	t := time.Unix(day, 0)
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	end := start.AddDate(0, 0, 1)
	return start.Unix(), end.Unix()
}
