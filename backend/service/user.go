package service

import (
	"atm/common"
	"atm/dao"
	"atm/model"
	"atm/response"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	dao dao.UserDao
}

func NewUserService() UserService {
	return UserService{dao.NewUserDao()}
}

func (service UserService) GetInfo(uid int64) (*model.UserInfo, int) {
	userInfo, err := service.dao.GetInfo(uid)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return userInfo, response.ErrCodeSuccess
}

func (service UserService) Delete(uid int64, params *model.UserDeleteParams) int {
	code := service.dao.GetCode(params.Email)
	if code != params.VerifyCode {
		return response.ErrCodeVerityCodeInvalid
	}
	// 删除用户所有数据
	if err := service.dao.DeleteData(params); err != nil {
		return response.ErrCodeFailed
	}
	return response.ErrCodeSuccess
}

func (service UserService) GetVerifyCode(params *model.UserGetVerifyCodeParams) int {
	// 生成6位随机数
	code := common.RandInt(100000, 999998)

	// 保存验证码
	if err := service.dao.SetCode(code, params.Email); err != nil {
		return response.ErrCodeFailed
	}

	// 发送验证码
	content := fmt.Sprintf("验证码%v，您正在找回密码，切勿向他人泄露。", code)
	err := common.SendVerifyCode(params.Email, content)
	if err != nil {
		return response.ErrCodeVerityCodeSendFailed
	}
	return response.ErrCodeSuccess
}

func (service UserService) Login(params *model.UserLoginParams) (*model.UserToken, int) {
	// 判断用户是否存在
	if !service.dao.IsExists(params.Email) {
		return nil, response.ErrCodeUserNotExist
	}

	// 获取用户信息
	user, err := service.dao.GetUser(params.Email)
	if err != nil {
		return nil, response.ErrCodeFailed
	}

	// 校验账号密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, response.ErrCodeUserEmailOrPass
	}

	// 生成并保存Token
	token, err := common.GenToken(user.Id)
	if err != nil {
		log.Printf("[error]Login:GenerateToken:%s", err)
		return nil, response.ErrCodeFailed
	}

	UserToken := model.UserToken{
		Uid:   user.Id,
		Token: token,
	}

	return &UserToken, response.ErrCodeSuccess
}

func (service UserService) Register(params *model.UserRegisterParams) int {
	if service.dao.IsExists(params.Email) {
		return response.ErrCodeUserHasExist
	}

	code := service.dao.GetCode(params.Email)
	if code != params.VerifyCode {
		return response.ErrCodeVerityCodeInvalid
	}


	password, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrCodeFailed
	}
	params.Password = string(password)


	if err := service.dao.Create(params); err != nil {
		return response.ErrCodeFailed
	}

	uid, err := service.dao.GetUid(params.Email)
	if err != nil {
		return response.ErrCodeFailed
	}

	return response.ErrCodeSuccess
}

func (service UserService) ForgetPass(params *model.UserForgetPassParams) int {
	if !service.dao.IsExists(params.Email) {
		return response.ErrCodeUserNotExist
	}

	code := service.dao.GetCode(params.Email)
	if code != params.VerifyCode {
		return response.ErrCodeVerityCodeInvalid
	}

	password, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrCodeFailed
	}
	if err := service.dao.UpdatePass(params.Email, string(password)); err != nil {
		return response.ErrCodeUserPassResetFailed
	}
	return response.ErrCodeSuccess
}
