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

func (dao ProjectDao) IsExists(uid int64, name string) bool {
	var count int64
	global.MysqlClient.Table(PROJECT).Where("creator = ? and name = ?", uid, name).Count(&count)
	return count == 1
}

func (dao ProjectDao) IsIdExists(uid int64, id int64) bool {
	var count int64
	global.MysqlClient.Table(PROJECT).Where("id = ? and creator = ?", id, uid).Count(&count)
	return count == 1
}


func (dao ProjectDao) GetInfo(params *model.ProjectQueryParam) (*model.ProjectInfo, error) {
	where := model.Project{
		Id: params.Id,
	}
	project := model.Project{}
	err := global.MysqlClient.Table(PROJECT).Where(where).First(&project).Error
	if err != nil {
		return nil, err 
	}
	projectInfo := model.ProjectInfo {
		Id: project.Id,
		Name: project.Name,
		// Cid: project.Cid,
		BeginTime: project.BeginTime,
		OverTime: project.OverTime,
		Finishrate: project.Finishrate,
		Remark: project.Remark,
		Outputslist: project.Outputslist,
	}
	return &projectInfo, nil
}

func (dao ProjectDao) GetList(uid int64, params *model.ProjectQueryParam) ([]*model.ProjectList, int64, error) {
	where := model.Project{
		Id: params.Id,
		Name: params.Name,
		Creator: uid,
	}
	projectList := []*model.ProjectList{}
	// rows, err := restPage(params.Page, PROJECT, &where, &projectList)

	var sortOpt string
	switch params.Sorttype {
	case 0:
	case 1: sortOpt = "id ASC"
	case 2: sortOpt = "finishrate DESC"
	case 3: sortOpt = "finishrate ASC"
	}
	if params.Page.PageNum > 0 && params.Page.PageSize > 0 {
		offset := (params.Page.PageNum - 1) * params.Page.PageSize
		global.MysqlClient.Table(PROJECT).Where(&where).Order(sortOpt).Offset(offset).Limit(params.Page.PageSize).Find(&projectList)
	}
	var rows int64
	res := global.MysqlClient.Table(PROJECT).Where(where).Count(&rows)

	return projectList, rows, res.Error
}

func (dao ProjectDao) GetListByUid(uid int64) ([]*model.ProjectList, error) {
	where := model.Project{
		Creator: uid,
	}
	projectList := []*model.ProjectList{}
	err := global.MysqlClient.Debug().Table(PROJECT).Where(&where).Find(&projectList).Error
	if err != nil {
		return nil, err
	}
	return projectList, err
}

func (dao ProjectDao) GetProjectById(id int64) (*model.Project, error) {
	var project model.Project
	err := global.MysqlClient.Debug().Table(PROJECT).First(&project, id).Error
	if err != nil {
		return nil, err 
	}
	return &project, nil
}


func (dao ProjectDao) Create(uid int64, params *model.ProjectCreateParam) error {
	project := model.Project {
		// Cid: params.Cid,
		Name: params.Name,
		BeginTime: params.BeginTime,
		OverTime: params.OverTime,
		Remark: params.Remark,
		Finishrate: params.Finishrate,
		Outputslist: params.Outputslist,
		Creator: uid,
	}
	return global.MysqlClient.Debug().Create(&project).Error
}

func (dao ProjectDao) Update(params *model.ProjectUpdateParam) error {
	project := model.Project {
		Id: params.Id,
		// Cid: params.Cid,
		Name: params.Name,
		BeginTime: params.BeginTime,
		OverTime: params.OverTime,
		Remark: params.Remark,
		Finishrate: params.Finishrate,
		Outputslist: params.Outputslist,
	}
	return global.MysqlClient.Debug().Model(&project).Select("*").Omit("id", "creator").Updates(&project).Error
}

func (dao ProjectDao) Delete(params *model.ProjectDeleteParam) error {
	return global.MysqlClient.Delete(&model.Project{}, params.Ids).Error
}