package model

type User struct {
	Id               int64  `gorm:"primaryKey"`
	Type             int    `gorm:"type"`
	Name             string `gorm:"name"`
	Email            string `gorm:"email"`
	Password         string `gorm:"password"`
	CreatedTime      int64  `gorm:"created"`
	SubscribeExpired int64  `gorm:"expired"`
}

type UserGetInfoParams struct {
}

type UserDeleteParams struct {
	Email      string `json:"email" binding:"required, email"`
	VerifyCode string `json:"code" binding:"required,len=6"`
}

type UserGetVerifyCodeParams struct {
	Email string `json:"email" binding:"required, email"`
}

type UserLoginParams struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterParams struct {
	Email      string `json:"email" binding:"required, email"`
	Password   string `json:"password" binding:"required"`
	VerifyCode string `json:"code" binding:"required,len=6"`
	Name       string `json:"name" binding:"required"`
}

type UserForgetPassParams struct {
	Email      string `json:"email" binding:"required, email"`
	Password   string
	VerifyCode string `json:"code" binding:"required,len=6"`
}

type UserToken struct {
	Uid   int64  `json:"uid"`
	Token string `json:"token"`
}

type UserInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Type int    `json:"type"`
}