package dao

import (
	"atm/global"
	"atm/model"
)

type ProjectDao struct {
}

func NewProjectDao() ProjectDao {
	return ProjectDao{}
}

func (dao ProjectDao) GetInfo(params *model.ProjectQueryParam) (*model.ProjectInfo, error) {
	where := model.Project{
		Id: params.Id,
	}
	projectInfo := model.ProjectInfo{}
	err := global.MysqlClient.Table(PROJECT).Where(where).First(&projectInfo).Error
	if err != nil {
		return nil, err 
	}
	return &projectInfo, nil
}

func (dao ProjectDao) GetList(uid int64, params *model.ProjectQueryParam) ([]*model.ProjectList, int64, error) {
	where := model.Project{
		Id: params.Id,
		Name: params.Name,
		Status: params.Status,
		Creator: uid,
	}
	projectList := []*model.ProjectList{}
	rows, err := restPage(params.Page, PROJECT, &where, &projectList)
	return projectList, rows, err
}

func (dao ProjectDao) GetListByUid(uid int64) ([]*model.ProjectList, error) {
	where := model.Project{
		Creator: uid,
	}
	projectList := []*model.ProjectList{}
	err := global.MysqlClient.Table(PROJECT).Where(&where).Find(&projectList).Error
	if err != nil {
		return nil, err
	}
	return projectList, err
}

func (dao ProjectDao) Create(uid int64, params *model.ProjectCreateParam) error {
	project := model.Project {
		Cid: params.Cid,
		Name: params.Name,
		BeginTime: params.BeginTime,
		OverTime: params.OverTime,
		Remarks: params.Remarks,
		Outputslist: params.Outputslist,
		Creator: uid,
	}
	return global.MysqlClient.Create(&project).Error
}

func (dao ProjectDao) Update(params *model.ProjectUpdateParam) error {
	project := model.Project {
		Id: params.Id,
		Cid: params.Cid,
		Name: params.Name,
		BeginTime: params.BeginTime,
		OverTime: params.OverTime,
		Remarks: params.Remarks,
		Outputslist: params.Outputslist,
	}
	return global.MysqlClient.Model(&project).Select("*").Omit("id", "creator","cid").Updates(&project).Error
}

func (dao ProjectDao) Delete(params *model.ProjectDeleteParam) error {
	return global.MysqlClient.Delete(&model.Project{}, params.Ids).Error
}