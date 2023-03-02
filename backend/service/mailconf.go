package service

import (
	"atm/common"
	"atm/dao"
	"atm/model"
	"atm/response"
)

type MailconfService struct {
	dao dao.MailconfDao
}

func NewMailconfService() MailconfService {
	return MailconfService{dao.NewMailconfDao()}
}

func (service MailconfService) Check(uid int64) int {
	mc, err := service.dao.GetInfo(uid)
	if err != nil {
		return response.ErrCodeMailConfigInvalid
	}
	if err = common.DialMail(mc.Stmp, mc.Port, mc.Email, mc.AuthCode); err != nil {
		return response.ErrCodeMailConfigInvalid
	}
	return response.ErrCodeSuccess
}

func (service MailconfService) UpdateStatus(uid int64, params *model.MailConfigStatusParam) int {
	if !service.dao.IsExists(uid) {
		return response.ErrCodeMailConfigUnSet
	}
	if err := service.dao.UpdateStatus(uid, params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service MailconfService) GetInfo(uid int64) (*model.MailConfigInfo, int) {
	if !service.dao.IsExists(uid) {
		return nil, response.ErrCodeMailConfigUnSet
	}
	mc, err := service.dao.GetInfo(uid)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	mailConfig := model.MailConfigInfo{
		Id:       mc.Id,
		Stmp:     mc.Stmp,
		Port:     mc.Port,
		AuthCode: mc.AuthCode,
		Email:    mc.Email,
		Status:   mc.Status,
	}
	return &mailConfig, response.ErrCodeSuccess
}

func (service MailconfService) Delete(params *model.MailConfigDeleteParam) int {
	if err := service.dao.Delete(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service MailconfService) Save(uid int64, params *model.MailConfigSaveParam) int {
	if err := service.dao.Save(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}