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

func (service ProjectService) UpdateOutputs(uid int64, params *model.ProjectOutputQueryParam) (*model.Outputslist, int) {
	if len(params.Oids) == 0 {
		return nil, response.ErrCodeSuccess
	}
	if params.Id != 0 && !service.dao.IsIdExists(uid, params.Id) {
		return nil, response.ErrCodeFailed
	}

	// before, err := service.dao.GetProjectById(params.Id)
	// if err != nil {
	// 	return nil, response.ErrCodeFailed
	// }

	tmplist := make(model.Outputslist, 0)
	newlist, err := service.output.GetListByIds(params.Oids)
	for i := 0; i < len(newlist); i++ {
		tmp := model.Outputs{
			Id:     newlist[i].Id,
			Name:   newlist[i].Name,
			Type:   newlist[i].Type,
			Weight: newlist[i].Weight,
			Status: newlist[i].Status,
		}
		tmplist = append(tmplist, &tmp)
	}
	// before.Outputslist = &tmplist
	
	if params.Id != 0 {
		// when create new project, project id is absent
		// return the outputlist corresponding to Oid list
		p := model.ProjectUpdateParam{
			Id:          params.Id,
			Outputslist: &tmplist,
		}
		err = service.dao.Update(&p)
		if err != nil {
			return nil, response.ErrCodeFailed
		}
	}
	return &tmplist, response.ErrCodeSuccess
}

func (service ProjectService) Delete(params *model.ProjectDeleteParam) int {
	if err := service.dao.Delete(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service ProjectService) Create(uid int64, params *model.ProjectCreateParam) int {
	if service.dao.IsExists(uid, params.Name) {
		return response.ErrCodeFailed
	}
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
		row.Finishrate = c.Finishrate
		row.Remark = c.Remark
		excelRows = append(excelRows, row)
	}
	sheet := "项目信息"
	columns := []string{"项目名称", "开始时间", "截止时间", "完成度", "状态", "备注"}
	fileName := "project_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.ErrCodeFailed
	}
	return file, response.ErrCodeSuccess
}
