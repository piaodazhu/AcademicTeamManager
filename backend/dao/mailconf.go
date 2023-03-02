package dao

import (
	"atm/global"
	"atm/model"
	"time"
)

const (
	Closed = 2
)

type MailconfDao struct {
}

func NewMailconfDao() MailconfDao {
	return MailconfDao{}
}

func (dao MailconfDao) IsExists(uid int64) bool {
	var count int64
	global.MysqlClient.Table(MAILCONFIG).Where("creator = ?", uid).Count(&count)
	return count == 1
}

func (dao MailconfDao) GetInfo(uid int64) (*model.MailConfig, error) {
	var mc model.MailConfig
	err := global.MysqlClient.Table(MAILCONFIG).Where("creator = ?", uid).First(&mc).Error
	if err != nil {
		return nil, err
	}
	return &mc, nil
}

func (dao MailconfDao) UpdateStatus(uid int64, params *model.MailConfigStatusParam) error {
	mailConfig := model.MailConfig{
		Status:  params.Status,
		Updated: time.Now().Unix(),
	}
	err := global.MysqlClient.Model(&mailConfig).Where("creator = ?", uid).Select("status", "updated").Updates(&mailConfig).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao MailconfDao) Delete(params *model.MailConfigDeleteParam) error {
	return global.MysqlClient.Delete(&model.MailConfig{}, params.Id).Error
}

func (dao MailconfDao) Save(uid int64, params *model.MailConfigSaveParam) error {
	if !dao.IsExists(uid) {
		return create(uid, params)
	}
	return update(uid, params)
}

func create(uid int64, param *model.MailConfigSaveParam) error {
	mailConfig := model.MailConfig{
		Stmp:     param.Stmp,
		Port:     param.Port,
		AuthCode: param.AuthCode,
		Email:    param.Email,
		Status:   Closed,
		Creator:  uid,
		Created:  time.Now().Unix(),
	}
	err := global.MysqlClient.Create(&mailConfig).Error
	if err != nil {
		return err
	}
	return nil
}

func update(uid int64, params *model.MailConfigSaveParam) error {
	mailConfig := model.MailConfig{
		Stmp:     params.Stmp,
		Port:     params.Port,
		AuthCode: params.AuthCode,
		Email:    params.Email,
		Creator:  uid,
		Updated:  time.Now().Unix(),
	}
	db := global.MysqlClient.Model(&mailConfig).Select("*").Omit("id", "status", "creator", "created").Where("creator = ?", uid)
	if err := db.Updates(&mailConfig).Error; err != nil {
		return err
	}
	return nil
}
