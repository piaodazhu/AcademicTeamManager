package model

type Output struct {
	Id          int64   `gorm:"primaryKey"`
	Name        string  `gorm:"column:name"`
	Type        int     `gorm:"column:type"`
	Weight      float64 `gorm:"column:weight"`
	Status      int     `gorm:"column:status"`
	Description string  `gorm:"column:description"`
	CreatorId   int64   `gorm:"column:creator"`
}

type OutputCreateParam struct {
	Name        string  `json:"name" binding:"required"`
	Type        int     `json:"type" binding:"required,oneof=1 2 3 4"`
	Weight      float64 `json:"weight" binding:"required,gt=0"`
	Description string  `json:"description" binding:"-"`
	Status      int     `json:"status" binding:"required,oneof=1 2"`
	CreatorId   int64   `json:"creator,omitempty" binding:"-"`
}

type OutputUpdateParam struct {
	Id          int64   `json:"id" binding:"required,gt=0"`
	Name        string  `json:"name" binding:"required"`
	Type        int     `json:"type" binding:"required,oneof=1 2 3 4"`
	Weight      float64 `json:"weight" binding:"required,gt=0"`
	Description string  `json:"description" binding:"-"`
	Status      int     `json:"status" binding:"required,oneof=1 2"`
	CreatorId   int64   `json:"creator,omitempty" binding:"-"`
}

type OutputDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type OutputQueryParam struct {
	Id        int64   `form:"id" binding:"omitempty,gt=0"`
	Ids       []int64 `form:"ids" json:"ids" binding:"-"`
	Name      string  `form:"name" binding:"-"`
	Type      int     `form:"type" binding:"omitempty,oneof=1 2 3 4"`
	Status    int     `form:"status" binding:"omitempty,oneof=1 2"`
	CreatorId int64   `form:"creator,omitempty" binding:"-"`
	Page      Page
}

type OutputList struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Type        int     `json:"type"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
}

type OutputInfo struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Type        int     `json:"type"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
}

type OutputExcelRecord struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}
