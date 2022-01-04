package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type SysUser struct {
}

func (s *SysUser) GetSysUser(name int, password string) (userInfo models.SysUser, err error) {
	adminUser := models.SysUser{}
	err = dao.Orm.First(&adminUser).Where("num = ? AND password = ？", name, password).Error
	if err != nil{
		log.Println(err)
	}

	return adminUser, err
}