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
<<<<<<< HEAD
	err = dao.Orm.Where("num = ? AND password = ?", name, password).Find(&adminUser).Error
	fmt.Println("cat")
=======
	err = dao.Orm.First(&adminUser).Where("num = ? AND password = ？", name, password).Error
>>>>>>> 81203b5 (选择A，B)
	if err != nil{
		log.Println(err)
	}

	return adminUser, err
}