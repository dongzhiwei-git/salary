package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
	"log"
)

type Print struct {
}

<<<<<<< HEAD
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
=======
func (s *Print) CreateSubInfo(num, money int) (err error) {
	info := models.Print{}
	info.Num = num
	info.Money = money
	err = dao.Orm.Create(&info).Error
	if err != nil {
		log.Println(err)
>>>>>>> 81203b5 (选择A，B)
	}

	return err
}
<<<<<<< HEAD

func (s *Print) GetSalaryInfo() (salaryInfo []models.Print, err error) {
	err = dao.Orm.Find(&salaryInfo).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return salaryInfo, err
}
=======
>>>>>>> 81203b5 (选择A，B)
