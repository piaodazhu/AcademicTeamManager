package model

type Student struct {
	Id             int64  `gorm:"primaryKey"`
	Name           string `gorm:"name"`
	Phone          string `gorm:"phone"`
	Email          string `gorm:"email"`
	CreatorId      int64  `gorm:"creator"`
	LastDiscussion int64  `gorm:"lastdis"`
	NextDiscussion int64  `gorm:"nextdis"`
	Status         int    `gorm:"status"`
	Remark         string `gorm:"remark"`
}


type StudentCreateParam struct {
	Name           string `json:"name" binding:"required"`
	Phone          string `json:"phone" binding:"omitempty,len=11"`
	Email          string `json:"email" binding:"omitempty,email"`
	Remark         string `json:"remark" binding:"-"`
}

type StudentUpdateParam struct {
	Id       int64  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"-"`
	Phone    string `json:"phone" binding:"omitempty,len=11"`
	Email    string `json:"email" binding:"omitempty,email"`
	LastDiscussion int64  `json:"lastdis" binding:"omitempty"`
	NextDiscussion int64  `json:"nextdis" binding:"omitempty"`
	Remark         string `json:"remark" binding:"-"`
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
	Id     int64  `form:"id" binding:"omitempty, gt=0"`
	Name   string `form:"name" binding:"omitempty, gt=0"`
	Status int `form:"status" binding:"omitempty, oneof=1 2"`
	Page   Page
}

type StudentInfo struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	LastDiscussion int64  `json:"lastdis"`
	NextDiscussion int64  `json:"nextdis"`
	Status         int    `json:"status"`
	Remark         string `json:"remark"`
}

type StudentList StudentInfo

type StudentOption struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type StudentExcelRecord struct {
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	LastDiscussion string  `json:"lastdis"`
	NextDiscussion string  `json:"nextdis"`
	Status         string    `json:"status"`
	Remark         string `json:"remark"`
}
