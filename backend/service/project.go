package service

import (
	"atm/common"
	"atm/dao"
	"atm/model"
	"atm/response"
	"strconv"
)

type ProjectService struct {
	dao    dao.ProjectDao
	output dao.OutputDao
}

func NewProjectService() ProjectService {
	return ProjectService{dao.NewProjectDao(), dao.NewOutputDao()}
}

func (service ProjectService) GetList(uid int64, params *model.ProjectQueryParam) ([]*model.ProjectList, int64, int) {
	projectList, rows, err := service.dao.GetList(uid, params)
	if err != nil {
		return nil, NumberNull, response.ErrCodeFailed
	}
	return projectList, rows, response.ErrCodeSuccess
}

func (service ProjectService) GetInfo(params *model.ProjectQueryParam) (*model.ProjectInfo, int) {
	projectInfo, err := service.dao.GetInfo(params)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return projectInfo, response.ErrCodeSuccess
}

func (service ProjectService) Delete(params *model.ProjectDeleteParam) int {
	if err := service.dao.Delete(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service ProjectService) Create(uid int64, params *model.ProjectCreateParam) int {
	if err := service.dao.Create(uid, params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service ProjectService) Update(params *model.ProjectUpdateParam) int {
	if err := service.dao.Update(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service ProjectService) Export(uid int64) (string, int) {
	projects, err := service.dao.GetListByUid(uid)
	if err != nil {
		return StringNull, response.ErrCodeFailed
	}
	excelRows := make([]model.ProjectExcelRecord, 0)
	var row model.ProjectExcelRecord
	for _, c := range projects {
		row.Name = c.Name
		row.BeginTime = c.BeginTime
		row.OverTime = c.OverTime
		row.Remarks = c.Remarks
		if c.Status == 1 {
			row.Status = "已完成"
		}
		if c.Status == 2 {
			row.Status = "未完成"
		}
		excelRows = append(excelRows, row)
	}
	sheet := "项目信息"
	columns := []string{"项目名称", "开始时间", "截止时间", "备注", "状态"}
	fileName := "project_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.ErrCodeFailed
	}
	return file, response.ErrCodeSuccess
}

// func (service ProjectService) GetOutputList(params *model.ProjectQueryParam) ([]*model.Outputs, int) {
// 	if params.Id == 0 {
// 		outputs, err := service.output.GetListByIds(params.Pids)
// 		if err != nil {
// 			return nil, response.ErrCodeFailed
// 		}
// 		return outputs, response.ErrCodeSuccess
// 	}

// 	// 默认已添加的产品列表
// 	project, err := service.dao.GetAddedPList(params.Id)
// 	if err != nil {
// 		return nil, response.ErrCodeFailed
// 	}

// 	// 最终已添加的产品列表
// 	addedOutputList := make([]*model.Outputs, 0)

// 	if len(params.Pids) == 0 {
// 		return addedOutputList, response.ErrCodeSuccess
// 	}

// 	addedPids := make([]int64, 0)
// 	for _, pid := range params.Pids {
// 		if len(*project.Outputlist) == 0 {
// 			addedPids = params.Pids
// 			break
// 		}
// 		for _, Output := range *project.Outputlist {
// 			if pid == Output.Id {
// 				addedOutputList = append(addedOutputList, Output)
// 				continue
// 			}
// 			addedPids = append(addedPids, pid)
// 		}
// 	}
// 	outputs, err := service.output.GetListByIds(addedPids)
// 	if err != nil {
// 		return nil, response.ErrCodeFailed
// 	}
// 	addedOutputList = append(addedOutputList, outputs...)
// 	return addedOutputList, response.ErrCodeSuccess
// }
