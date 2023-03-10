package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Project struct {
	Id        int64  `gorm:"primaryKey"`
	Name      string `gorm:"column:name"`
	BeginTime string `gorm:"column:begin_time"`
	OverTime  string `gorm:"column:over_time"`
	Remark    string `gorm:"column:remark"`
	// Cid         int64        `gorm:"column:cid"`
	Outputslist *Outputslist `gorm:"column:outputlist;type:json"`
	Finishrate  float64      `gorm:"column:finishrate"`
	Creator     int64        `gorm:"column:creator"`
}

type ProjectCreateParam struct {
	Name       string  `json:"name" binding:"required"`
	BeginTime  string  `json:"beginTime" binding:"-"`
	OverTime   string  `json:"overTime" binding:"-"`
	Remark     string  `json:"remark" binding:"-"`
	Finishrate float64 `json:"finishrate" binding:"-"`
	// Cid         int64        `json:"cid" binding:"required,gt=0"`
	Outputslist *Outputslist `json:"outputlist"`
}

type ProjectUpdateParam struct {
	Id         int64   `json:"id" binding:"required,gt=0"`
	Name       string  `json:"name" binding:"-"`
	BeginTime  string  `json:"beginTime" binding:"-"`
	OverTime   string  `json:"overTime" binding:"-"`
	Remark     string  `json:"remark" binding:"-"`
	Finishrate float64 `json:"finishrate" binding:"-"`
	// Cid         int64        `json:"cid" binding:"omitempty,gt=0"`
	Outputslist *Outputslist `json:"outputlist"`
}

type ProjectDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type ProjectQueryParam struct {
	Id     int64  `form:"id" binding:"omitempty,gt=0"`
	Name   string `form:"name" binding:"-"`
	Sorttype int    `form:"sorttype" binding:"omitempty,oneof=1 2 3"`
	Page   Page
}

type ProjectOutputQueryParam struct {
	Id   int64   `form:"id" binding:"omitempty,gt=0"`
	Oids []int64 `form:"oids" binding:"-"`
}

type ProjectList struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	BeginTime  string  `json:"beginTime"`
	OverTime   string  `json:"overTime"`
	Remark     string  `json:"remark"`
	Finishrate float64 `json:"finishrate" binding:"-"`
}

type ProjectInfo struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	Cid         int64        `json:"cid"`
	BeginTime   string       `json:"beginTime"`
	OverTime    string       `json:"overTime"`
	Remark      string       `json:"remark"`
	Outputslist *Outputslist `json:"outputlist"`
	Finishrate  float64      `json:"finishrate" binding:"-"`
}

type ProjectExcelRecord struct {
	Name       string  `json:"name"`
	BeginTime  string  `json:"beginTime"`
	OverTime   string  `json:"overTime"`
	Finishrate float64 `json:"finishrate" binding:"-"`
	Remark     string  `json:"remark"`
}

type Outputs struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Type   int     `json:"type"`
	Status int     `json:"status"`
	Weight float64 `json:"weight"`
}

type Outputslist []*Outputs

func (p *Outputslist) Value() (driver.Value, error) {
	b, err := json.Marshal(p)
	return string(b), err
}

func (p *Outputslist) Scan(src any) error {
	return json.Unmarshal(src.([]byte), p)
}
