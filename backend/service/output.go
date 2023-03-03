package service

import (
	"atm/common"
	"atm/dao"
	"atm/model"
	"atm/response"
	"strconv"
)

type OutputService struct {
	dao dao.OutputDao
}

func NewOutputService() OutputService {
	return OutputService{dao.NewOutputDao()}
}

func (service OutputService) GetList(uid int64, params *model.OutputQueryParam) ([]*model.OutputList, int64, int) {
	outputList, rows, err := service.dao.GetList(uid, params)
	if err != nil {
		return nil, NumberNull, response.ErrCodeFailed
	}
	return outputList, rows, response.ErrCodeSuccess
}

func (service OutputService) GetInfo(params *model.OutputQueryParam) (*model.OutputInfo, int) {
	outputInfo, err := service.dao.GetInfo(params)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return outputInfo, response.ErrCodeSuccess
}

func (service OutputService) Delete(params *model.OutputDeleteParam) int {
	if err := service.dao.Delete(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service OutputService) Create(uid int64, params *model.OutputCreateParam) int {
	if service.dao.IsExists(uid, params.Name) {
		return response.ErrCodeOutputHasExist
	}
	if err := service.dao.Create(uid, params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service OutputService) Update(params *model.OutputUpdateParam) int {
	if err := service.dao.Update(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service OutputService) Export(uid int64) (string, int) {
	outputs, err := service.dao.GetListByUid(uid)
	if err != nil {
		return StringNull, response.ErrCodeFileExportFailed
	}
	excelRows := make([]model.OutputExcelRecord, 0)
	var row model.OutputExcelRecord
	for _, p := range outputs {
		row.Name = p.Name
		switch p.Type {
		case 1:
			row.Type = "期刊论文"
		case 2:
			row.Type = "会议论文"
		case 3:
			row.Type = "学术专著"
		case 4:
			row.Type = "发明专利"
		}

		switch p.Status {
		case 1:
			row.Status = "推进中"
		case 2:
			row.Status = "已完成"
		}
		row.Weight = p.Weight
		row.Description = p.Description
		excelRows = append(excelRows, row)
	}
	sheet := "成果产出"
	columns := []string{"成果名称", "成果类型", "权重", "产出状态"}
	fileName := "output_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.ErrCodeFileExportFailed
	}
	return file, response.ErrCodeSuccess
}
