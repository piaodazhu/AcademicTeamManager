package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Project struct {
	Id          int64        `gorm:"primaryKey"`
	Name        string       `gorm:"column:name"`
	BeginTime   string       `gorm:"column:begin_time"`
	OverTime    string       `gorm:"column:over_time"`
	Remarks     string       `gorm:"column:remarks"`
	Cid         int64        `gorm:"column:cid"`
	Outputslist *Outputslist `gorm:"type:json,column:outputlist"`
	Status      int          `gorm:"column:status"`
	Creator     int64        `gorm:"column:creator"`
}

type ProjectCreateParam struct {
	Name        string       `json:"name" binding:"required"`
	BeginTime   string       `json:"beginTime" binding:"-"`
	OverTime    string       `json:"overTime" binding:"-"`
	Remarks     string       `json:"remarks" binding:"-"`
	Cid         int64        `json:"cid" binding:"required,gt=0"`
	Outputslist *Outputslist `json:"productlist"`
}

type ProjectUpdateParam struct {
	Id          int64        `json:"id" binding:"required,gt=0"`
	Name        string       `json:"name" binding:"required"`
	BeginTime   string       `json:"beginTime" binding:"-"`
	OverTime    string       `json:"overTime" binding:"-"`
	Remarks     string       `json:"remarks" binding:"-"`
	Cid         int64        `json:"cid" binding:"required,gt=0"`
	Outputslist *Outputslist `json:"productlist"`
}

type ProjectDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type ProjectQueryParam struct {
	Id     int64  `form:"id" binding:"omitempty,gt=0"`
	Name   string `form:"name" binding:"-"`
	Status int    `form:"status" binding:"omitempty,oneof=1 2"`
	Page   Page
}

type ProjectList struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	BeginTime string `json:"beginTime"`
	OverTime  string `json:"overTime"`
	Remarks   string `json:"remarks"`
	Status    int    `json:"status"`
}

type ProjectInfo struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	Cid         int64        `json:"cid"`
	BeginTime   string       `json:"beginTime"`
	OverTime    string       `json:"overTime"`
	Remarks     string       `json:"remarks"`
	Outputslist *Outputslist `json:"productlist"`
	Status      int          `json:"status"`
}

type ProjectExcelRecord struct {
	Name      string `json:"name"`
	BeginTime string `json:"beginTime"`
	OverTime  string `json:"overTime"`
	Status    string `json:"status"`
	Remarks   string `json:"remarks"`
}

type Outputs struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Type   int     `json:"type"`
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
