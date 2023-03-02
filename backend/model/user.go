package model

type User struct {
	Id               int64  `gorm:"primaryKey"`
	Type             int    `gorm:"column:type"`
	Name             string `gorm:"column:name"`
	Email            string `gorm:"column:email"`
	Password         string `gorm:"column:password"`
	CreatedTime      int64  `gorm:"column:created"`
	SubscribeExpired int64  `gorm:"column:expired"`
}

type UserDeleteParams struct {
	Email      string `json:"email" binding:"required,email"`
	VerifyCode string `json:"code" binding:"required,len=6"`
}

type UserGetVerifyCodeParams struct {
	Email string `form:"email" binding:"required,email"`
}

type UserLoginParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterParams struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	VerifyCode string `json:"code" binding:"required,len=6"`
	Name       string `json:"name" binding:"required"`
}

type UserForgetPassParams struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	VerifyCode string `json:"code" binding:"required,len=6"`
}

type UserToken struct {
	Uid   int64  `json:"uid"`
	Token string `json:"token"`
}

type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  int    `json:"type"`
}
