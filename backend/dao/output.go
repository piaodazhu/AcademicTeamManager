package dao

import (
	"atm/global"
	"atm/model"
)

type OutputDao struct {
}

func NewOutputDao() OutputDao {
	return OutputDao{}
}


func (dao OutputDao) IsExists(uid int64, name string) bool {
	var count int64
	global.MysqlClient.Table(OUTPUT).Where("creator = ? and name = ?", uid, name).Count(&count)
	return count == 1
}

func (dao OutputDao) GetInfo(params *model.OutputQueryParam) (*model.OutputInfo, error) {
	where := model.Output {
		Id: params.Id,
	}
	outputInfo := model.OutputInfo{}
	err := global.MysqlClient.Table(OUTPUT).Where(&where).First(&outputInfo).Error
	if err != nil {
		return nil, err 
	}
	return &outputInfo, nil
}

func (dao OutputDao) GetList(uid int64, params *model.OutputQueryParam) ([]*model.OutputList, int64, error) {
	where := model.Output {
		Id: params.Id,
		Name: params.Name,
		Type: params.Type,
		Status: params.Status,
		CreatorId: uid,
	}
	outputList := []*model.OutputList{}
	rows, err := restPage(params.Page, OUTPUT, &where, &outputList)
	if err != nil {
		return nil, 0, err 
	}
	return outputList, rows, err
}

func (dao OutputDao) GetListByUid(uid int64) ([]*model.OutputList, error) {
	where := model.Output {
		CreatorId: uid,
	}
	outputList := []*model.OutputList{}
	err := global.MysqlClient.Table(OUTPUT).Where(&where).Find(&outputList).Error
	if err != nil {
		return nil, err 
	}
	return outputList, nil
}

func (dao OutputDao) Create(uid int64, params *model.OutputCreateParam) error {
	output := model.Output {
		Name: params.Name,
		Type: params.Type,
		Weight: params.Weight,
		Description: params.Description,
		Status: params.Status,
		CreatorId: uid,
	}
	return global.MysqlClient.Create(&output).Error
}

func (dao OutputDao) Update(params *model.OutputUpdateParam) error {
	output := model.Output {
		Id: params.Id,
		Name: params.Name,
		Type: params.Type,
		Weight: params.Weight,
		Description: params.Description,
		Status: params.Status,
	}
	return global.MysqlClient.Table(OUTPUT).Select("*").Omit("id", "creator").Updates(&output).Error
}

func (dao OutputDao) Delete(params *model.OutputDeleteParam) error {
	return global.MysqlClient.Delete(&model.Output{}, params.Ids).Error
}

// func (dao OutputDao) GetListByUids(uid int64) ([]*model.OutputList, error) {	
// 	where := 
// }