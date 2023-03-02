package dao

import (
	"atm/common"
	"atm/global"
	"atm/model"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type UtilDao struct {
}

func NewUtilDao() UtilDao {
	return UtilDao{}
}

func (dao UtilDao) InitDatabase() error {
	env := global.Conf.ServerConf.RunEnv
	if env == Prod {
		dbFile := global.Conf.MysqlConf.InitSQL
		sql, err := os.ReadFile(dbFile)
		if err != nil {
			log.Printf("Common.InitDatabase.Error: read file %s error: %s", dbFile, err)
			return err
		}
		sqls := strings.Split(string(sql), ";")
		for _, sql := range sqls {
			s := strings.TrimSpace(sql)
			if s == StringNull {
				continue
			}
			if err := global.MysqlClient.Exec(s).Error; err != nil {
				log.Printf("Common.InitDatabase.Error: %s", err)
				return err
			}
		}
		return nil
	}
	return nil
}

func (dao UtilDao) FileUpload(file *multipart.FileHeader) (*model.FileInfo, error) {
	dist := global.Conf.PathConf.Path
	name := common.GenUUID() + path.Ext(file.Filename)
	dn := dist + name
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	out, err := os.Create(dn)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return nil, err
	}
	flieInfo := model.FileInfo{
		Url:  dn,
		Name: name,
	}
	return &flieInfo, err
}

func (dao UtilDao) FileRemove(fileName string) error {
	file := global.Conf.PathConf.Path + fileName
	return os.Remove(file)
}

func restPage(page model.Page, name string, query interface{}, dest interface{}) (int64, error) {
	if page.PageNum > 0 && page.PageSize > 0 {
		offset := (page.PageNum - 1) * page.PageSize
		global.MysqlClient.Offset(offset).Limit(page.PageSize).Table(name).Where(query).Find(dest)
	}
	var rows int64
	res := global.MysqlClient.Table(name).Where(query).Count(&rows)
	return rows, res.Error
}
