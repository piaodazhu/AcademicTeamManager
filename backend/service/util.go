package service

import (
	"atm/dao"
	"atm/model"
	"atm/response"
	"mime/multipart"
)

type UtilService struct {
	dao dao.UtilDao
}

func NewUtilService() UtilService {
	return UtilService{dao.NewUtilDao()}
}

func (service UtilService) InitDatabase() int {
	if err := service.dao.InitDatabase(); err != nil {
		return response.ErrCodeInitDataFailed
	}
	return response.ErrCodeSuccess
}


func (service UtilService) FileUpload(file *multipart.FileHeader) (*model.FileParam, int) {
	fileInfo, err := service.dao.FileUpload(file)
	if err != nil {
		return nil, response.ErrCodeFileUploadFailed
	}
	return fileInfo, response.ErrCodeSuccess
}


func (service UtilService) FileRemove(param *model.FileParam) int {
	if err := service.dao.FileRemove(param.Name); err != nil {
		return response.ErrCodeFileUploadFailed
	}
	return response.ErrCodeSuccess
}
