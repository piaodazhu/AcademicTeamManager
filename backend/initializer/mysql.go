package initializer

import (
	"atm/global"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initMysql() {
	conf := global.Conf.MysqlConf
	var err error
	global.MysqlClient, err = gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.DB)),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		})
	if err != nil {
		panic("[ERROR] failed to initialize mysql")
	}
	sqldb, err := global.MysqlClient.DB()
	if err != nil {
		panic("[ERROR] failed to open mysql client")
	}
	sqldb.SetMaxIdleConns(conf.MaxIdleConns)
	sqldb.SetMaxOpenConns(conf.MaxOpenConns)
	sqldb.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetime))
}
