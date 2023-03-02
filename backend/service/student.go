package service

import (
	"atm/dao"
	"atm/model"
	"atm/common"
	"atm/response"
	"strconv"
	"time"
)

const (
	Open  = 1
	Close = 2
)

type StudentService struct {
	dao  dao.StudentDao
	mail dao.MailconfDao
}

func NewStudentService() StudentService {
	return StudentService{dao.NewStudentDao(), dao.NewMailconfDao()}
}

func (service StudentService) GetList(uid int64, params *model.StudentQueryParams) ([]*model.StudentList, int64, int) {
	customerList, rows, err := service.dao.GetList(uid, params)
	if err != nil {
		return nil, NumberNull, response.ErrCodeFailed
	}
	return customerList, rows, response.ErrCodeSuccess
}

func (service StudentService) GetOption(uid int64) ([]*model.StudentOption, int) {
	option, err := service.dao.GetOption(uid)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return option, response.ErrCodeSuccess
}

func (service StudentService) GetInfo(params *model.StudentQueryParams) (*model.StudentInfo, int) {
	customerInfo, err := service.dao.GetInfo(params)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return customerInfo, response.ErrCodeSuccess
}

func (service StudentService) Delete(params *model.StudentDeleteParam) int {
	if err := service.dao.Delete(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service StudentService) Export(uid int64) (string, int) {
	customers, err := service.dao.GetListByUid(uid)
	if err != nil {
		return StringNull, response.ErrCodeFailed
	}
	excelRows := make([]model.StudentExcelRecord, 0)
	var row model.StudentExcelRecord
	for _, c := range customers {
		row.Name = c.Name
		row.Phone = c.Phone
		row.Email = c.Email
		row.LastDiscussion = time.Unix(c.LastDiscussion, 0).Format("2006-01-02")
		row.NextDiscussion = time.Unix(c.NextDiscussion, 0).Format("2006-01-02")
		row.Remark = c.Remark
		if c.Status == 1 {
			row.Status = "暂无安排"
		}
		if c.Status == 2 {
			row.Status = "计划讨论"
		}
		excelRows = append(excelRows, row)
	}
	sheet := "学生信息"
	columns := []string{"姓名", "手机号", "邮箱", "上次讨论时间", "下次讨论时间", "状态", "备注"}
	fileName := "student_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.ErrCodeFailed
	}
	return file, response.ErrCodeSuccess
}

func (service StudentService) Create(uid int64, params *model.StudentCreateParam) int {
	if service.dao.IsExists(uid, params.Name) {
		return response.ErrCodeCustomerHasExist
	}
	if err := service.dao.Create(uid, params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service StudentService) Update(params *model.StudentUpdateParam) int {
	if err := service.dao.Update(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service StudentService) SendMail(uid int64, params *model.StudentSendMailParam) int {
	mc, err := service.mail.GetInfo(uid)
	if err != nil {
		return response.ErrCodeMailSendFailed
	}
	if mc.Status == Close {
		return response.ErrCodeMailSendUnEnable
	}
	mail := model.MailParam{
		Smtp:       mc.Stmp,
		Port:       mc.Port,
		AuthCode:   mc.AuthCode,
		Sender:     mc.Email,
		Subject:    params.Subject,
		Content:    params.Content,
		Attachment: params.Attachment,
		Receiver:   params.Receiver,
	}
	if err := common.SendMail(mail); err != nil {
		return response.ErrCodeMailSendFailed
	}
	return response.ErrCodeSuccess
}
