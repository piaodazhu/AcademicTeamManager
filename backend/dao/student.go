package dao

import (
	"atm/global"
	"atm/model"
)

type StudentDao struct {
}

func NewStudentDao() StudentDao {
	return StudentDao{}
}

func (dao StudentDao) GetInfo(params *model.StudentQueryParams) (*model.StudentInfo, error) {
	where := model.Student{
		Id: params.Id,
	}
	studentInfo := model.StudentInfo{}
	err := global.MysqlClient.Table(STUDENT).Where(&where).First(&studentInfo).Error
	if err != nil {
		return nil, err
	}
	return &studentInfo, nil
}

func (dao StudentDao) GetList(uid int64, params *model.StudentQueryParams) ([]*model.StudentList, int64, error) {
	where := model.Student{
		Id:        params.Id,
		Name:      params.Name,
		Status:    params.Status,
		CreatorId: uid,
	}
	studentList := []*model.StudentList{}
	rows, err := restPage(params.Page, STUDENT, where, &studentList)
	if err != nil {
		return nil, 0, err
	}
	return studentList, rows, nil
}

func (dao StudentDao) GetListByUid(uid int64) ([]*model.StudentList, error) {
	where := model.Student{
		CreatorId: uid,
	}
	list := []*model.StudentList{}
	err := global.MysqlClient.Table(STUDENT).Where(&where).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (dao StudentDao) GetOption(uid int64) ([]*model.StudentOption, error) {
	where := model.Student{
		CreatorId: uid,
	}
	option := []*model.StudentOption{}
	err := global.MysqlClient.Table(STUDENT).Where(&where).Find(&option).Error
	if err != nil {
		return nil, err
	}
	return option, nil
}

func (dao StudentDao) IsExists(uid int64, name string) bool {
	var count int64
	global.MysqlClient.Table(STUDENT).Where("creator = ? and name = ?", name, uid).Count(&count)
	return count == 1
}

func (dao StudentDao) Create(uid int64, params *model.StudentCreateParam) error {
	student := model.Student{
		Name: params.Name,
		Phone: params.Phone,
		Email: params.Email,
		CreatorId: uid,
		LastDiscussion: 0,
		NextDiscussion: 0,
		Status: 0,
		Remark: params.Remark,
	}
	return global.MysqlClient.Create(&student).Error
}

func (dao StudentDao) Update(params *model.StudentUpdateParam) error {
	student := model.Student{
		Id: params.Id,
		Name: params.Name,
		Phone: params.Phone,
		Email: params.Email,
		LastDiscussion: params.LastDiscussion,
		NextDiscussion: params.NextDiscussion,
		Remark: params.Remark,
	}
	return global.MysqlClient.Model(&student).Select("*").Omit("id", "creator", "status").Updates(&student).Error
}

func (dao StudentDao) Delete(params *model.StudentDeleteParam) error {
	return global.MysqlClient.Delete(&model.Student{}, params.Ids).Error
}