package services

import (
	"fmt"
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type SysUser struct {
}

func (s *SysUser) GetSysUser(name int, password string) (userInfo models.SysUser, err error) {
	adminUser := models.SysUser{}
	err = dao.Orm.Where("num = ? AND password = ?", name, password).Find(&adminUser).Error
	fmt.Println("cat")
	if err != nil{
		log.Println(err)
	}

	return adminUser, err
}