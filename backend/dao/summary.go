package dao

import (
	"atm/global"
	"atm/model"
)

type SummaryDao struct {
}

func NewSummaryDao() SummaryDao {
	return SummaryDao{}
}

func (dao SummaryDao) Count(uid int64) *model.Summary{
	summary := model.Summary{}
	var count int64
	global.MysqlClient.Table(STUDENT).Where("creator = ?", uid).Count(&count)
	summary.Students = int(count)
	global.MysqlClient.Table(PROJECT).Where("creator = ?", uid).Count(&count)
	summary.Projects = int(count)
	global.MysqlClient.Table(OUTPUT).Where("creator = ?", uid).Count(&count)
	summary.Outputs = int(count)
	return &summary
}