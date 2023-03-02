package model

type MailConfig struct {
	Id       int64  `gorm:"primaryKey"`
	Stmp     string `gorm:"column:stmp"`
	Port     int    `gorm:"column:port"`
	AuthCode string `gorm:"column:auth_code"`
	Email    string `gorm:"column:email"`
	Status   int    `gorm:"column:status"`
	Creator  int64  `gorm:"column:creator"`
	Created  int64  `gorm:"column:created"`
	Updated  int64  `gorm:"column:updated"`
}

type MailConfigSaveParam struct {
	Stmp     string `json:"stmp" binding:"required,ip|hostname"`
	Port     int    `json:"port" binding:"required,gt=0"`
	AuthCode string `json:"authCode" binding:"required,gt=0"`
	Email    string `json:"email" binding:"required,email"`
}

type MailConfigStatusParam struct {
	Status  int   `json:"status" binding:"required,oneof=1 2"`
}

type MailConfigDeleteParam struct {
	Id int64 `json:"id" binding:"required,gt=0"`
}

type MailConfigInfo struct {
	Id        int64  `json:"id"`
	Stmp      string `json:"stmp"`
	Port      int    `json:"port"`
	AuthCode  string `json:"authCode"`
	Email     string `json:"email"`
	Status    int    `json:"status"`
	Usability int    `json:"usability"`
}
