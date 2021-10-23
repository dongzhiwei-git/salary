package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type Print struct {
}

func (s *Print) CreateSubInfo(num, money int, name, scheme string) (err error) {
	info := models.Print{}
	info.Num = num
	info.Money = money
	info.UserName = name
	info.Scheme = scheme
	err = dao.Orm.Create(&info).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func (s *Print) GetSalaryInfo() (salaryInfo []models.Print, err error) {
	err = dao.Orm.Find(&salaryInfo).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return salaryInfo, err
}
