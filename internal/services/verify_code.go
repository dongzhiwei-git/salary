package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type Code struct {
}

// 校验code
func (c *Code) VerifyCode(code string) (codeInfo models.Code, err error) {
	info := models.Code{}
	err = dao.Orm.Find(&info).Where("code = ?", code).Error
	if err != nil {
		log.Println(err)
		return info, err
	}

	return info, err
}
