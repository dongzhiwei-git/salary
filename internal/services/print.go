package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type Print struct {
}

func (s *Print) CreateSubInfo(num, money int) (err error) {
	info := models.Print{}
	info.Num = num
	info.Money = money
	err = dao.Orm.Create(&info).Error
	if err != nil {
		log.Println(err)
	}

	return err
}
