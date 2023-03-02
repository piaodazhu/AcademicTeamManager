package dao

import (
	"atm/global"
	"atm/model"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UserDao struct {
}

func NewUserDao() UserDao {
	return UserDao{}
}

func (dao UserDao) IsExists(email string) bool {
	var count int64
	global.MysqlClient.Table(USER).Where("email = ?", email).Count(&count)
	return count == 1
}

func (dao UserDao) GetUser(email string) (*model.User, error) {
	var user model.User
	err := global.MysqlClient.Table(USER).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao UserDao) GetInfo(uid int64) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	err := global.MysqlClient.Table(USER).Where("id = ?", uid).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (dao UserDao) Create(params *model.UserRegisterParams) error {
	user := model.User{
		Email:            params.Email,
		Password:         params.Password,
		Name:             params.Name,
		Type:             1,
		CreatedTime:      time.Now().Unix(),
		SubscribeExpired: 0,
	}
	return global.MysqlClient.Create(&user).Error
}

func (dao UserDao) GetUid(email string) (int64, error) {
	var uid int64
	err := global.MysqlClient.Table(USER).Where("email = ?", email).Select("id").Scan(&uid).Error
	fmt.Println("!!! uid = ", uid)
	return uid, err
}

func (dao UserDao) DeleteAccount(uid int64, params *model.UserDeleteParams) error {
	return global.MysqlClient.Transaction(
		func(tx *gorm.DB) error {
			tables := map[interface{}]string{
				&model.Student{}: "creator = ?",
				&model.Project{}: "creator = ?",
				&model.Output{}:  "creator = ?",
				&model.User{}:    "id = ?",
			}
			for k, v := range tables {
				if err := tx.Where(v, uid).Delete(k).Error; err != nil {
					return err
				}
			}
			return nil
		})

}

func (dao UserDao) ForgetPass(email, password string) error {
	return global.MysqlClient.Model(&model.User{}).Where("email = ?", email).Update("password", password).Error
}

func (dao UserDao) SetCode(email string, code int) error {
	fmt.Println("setcode: ", code)
	return global.RedisClient.Set(ctx, fmt.Sprintf("VERIFYCODE:%s", email), strconv.Itoa(code), 10*time.Minute).Err()
}

func (dao UserDao) GetCode(email string) string {
	fmt.Println("getcode: ", global.RedisClient.Get(ctx, fmt.Sprintf("VERIFYCODE:%s", email)).Val())
	return global.RedisClient.Get(ctx, fmt.Sprintf("VERIFYCODE:%s", email)).Val()
}
