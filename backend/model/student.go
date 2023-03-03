package model

type Student struct {
	Id             int64  `gorm:"primaryKey"`
	Name           string `gorm:"column:name"`
	Type           int    `gorm:"column:type"`
	Phone          string `gorm:"column:phone"`
	Email          string `gorm:"column:email"`
	CreatorId      int64  `gorm:"column:creator"`
	LastDiscussion string `gorm:"column:lastdis"`
	NextDiscussion string `gorm:"column:nextdis"`
	Status         int    `gorm:"column:status"`
	Remark         string `gorm:"column:remark"`
}

type StudentCreateParam struct {
	Name    string `json:"name" binding:"required"`
	Type    int    `json:"type" binding:"required,oneof=1 2 3"`
	Phone   string `json:"phone" binding:"omitempty,len=11"`
	Email   string `json:"email" binding:"omitempty,email"`
	Lastdis string `json:"lastdis" binding:"omitempty"`
	Nextdis string `json:"nextdis" binding:"omitempty"`
	Status  int    `json:"status" binding:"omitempty,oneof=0 1 2"`
	Remark  string `json:"remark" binding:"-"`
}

type StudentUpdateParam struct {
	Id      int64  `json:"id" binding:"required"`
	Name    string `json:"name" binding:"-"`
	Type    int    `json:"type" binding:"omitempty,oneof=1 2 3"`
	Phone   string `json:"phone" binding:"omitempty,len=11"`
	Email   string `json:"email" binding:"omitempty,email"`
	Lastdis string `json:"lastdis" binding:"omitempty"`
	Nextdis string `json:"nextdis" binding:"omitempty"`
	Status  int    `json:"status" binding:"omitempty,oneof=0 1 2"`
	Remark  string `json:"remark" binding:"-"`
}

type StudentSendMailParam struct {
	Uid        int64  `json:"uid" binding:"-"`
	Receiver   string `json:"receiver" binding:"required,email"`
	Subject    string `json:"subject" binding:"omitempty,gt=0"`
	Content    string `json:"content" binding:"required,gt=0"`
	Attachment string `json:"attachment" binding:"omitempty,gt=0"`
}

type StudentDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type StudentQueryParams struct {
	Id     int64  `form:"id" binding:"omitempty,gt=0"`
	Name   string `form:"name" binding:"omitempty,gt=0"`
	Type   int    `form:"type" binding:"omitempty,oneof=1 2 3"`
	Status int    `form:"status" binding:"omitempty,oneof=1 2"`
	Page   Page
}

type StudentInfo struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Lastdis string `json:"lastdis"`
	Nextdis string `json:"nextdis"`
	Status  int    `json:"status"`
	Remark  string `json:"remark"`
}

type StudentList StudentInfo

type StudentOption struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

type StudentExcelRecord struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	LastDiscussion string `json:"lastdis"`
	NextDiscussion string `json:"nextdis"`
	Status         string `json:"status"`
	Remark         string `json:"remark"`
}
