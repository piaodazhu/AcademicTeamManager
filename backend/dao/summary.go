package dao

import (
	"atm/global"
	"atm/model"
	"encoding/json"
)

type SummaryDao struct {
}

func NewSummaryDao() SummaryDao {
	return SummaryDao{}
}

func (dao SummaryDao) Count(uid int64) *model.Summary {
	summary := model.Summary{}
	var count int64
	global.MysqlClient.Table(STUDENT).Where("creator = ?", uid).Count(&count)
	summary.StudentNum = int(count)
	global.MysqlClient.Table(PROJECT).Where("creator = ?", uid).Count(&count)
	summary.ProjectNum = int(count)
	global.MysqlClient.Table(OUTPUT).Where("creator = ?", uid).Count(&count)
	summary.OutputNum = int(count)
	return &summary
}

type GroupResult struct {
	Type int
	Cnt  int
}

func (dao SummaryDao) AnalyzeStudent(uid int64) []*model.PanArg {
	var res []GroupResult
	ret := []*model.PanArg{}
	global.MysqlClient.Raw("select type, count(*) as cnt from student where creator = ? group by type", uid).Scan(&res)
	for _, v := range res {
		tmp := model.PanArg{Value: v.Cnt}
		switch v.Type {
		case 1:
			tmp.Name = "博士"
		case 2:
			tmp.Name = "硕士"
		case 3:
			tmp.Name = "本科"
		default:
			tmp.Name = "未知"
		}
		ret = append(ret, &tmp)
	}
	return ret
}

type ScanResult struct {
	Name       string
	Outputlist string
}

func (dao SummaryDao) AnalyzeProject(uid int64) []*model.PanArg {
	res := []ScanResult{}
	ret := []*model.PanArg{}
	global.MysqlClient.Raw("select name, outputlist from project where creator = ?", uid).Scan(&res)
	for _, v := range res {
		p := []interface{}{}
		json.Unmarshal([]byte(v.Outputlist), &p)
		tmp := model.PanArg{Name: v.Name, Value: len(p)}
		ret = append(ret, &tmp)
	}
	return ret
}

func (dao SummaryDao) AnalyzeOutput(uid int64) []*model.PanArg {
	res := []GroupResult{}
	ret := []*model.PanArg{}
	global.MysqlClient.Raw("select type, count(*) as cnt from output where creator = ? group by type", uid).Scan(&res)
	for _, v := range res {
		tmp := model.PanArg{Value: v.Cnt}
		switch v.Type {
		case 1:
			tmp.Name = "期刊论文"
		case 2:
			tmp.Name = "会议论文"
		case 3:
			tmp.Name = "学术专著"
		case 4:
			tmp.Name = "发明专利"
		default:
			tmp.Name = "其他成果"
		}
		ret = append(ret, &tmp)
	}
	return ret
}
