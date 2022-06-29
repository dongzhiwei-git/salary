package models

type Print struct {
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> d73ad65 (基本完成)
	ID       int    `gorm:"id" json:"id"`
	UserName string `gorm:"user_name" json:"user_name" binding:"required"`
	Scheme   string `json:"scheme" json:"scheme" binding:"required"`
	Num      int    `gorm:"num" json:"num" binding:"required"`
	Money    int    `gorm:"money" json:"money" binding:"required"`
<<<<<<< HEAD
=======
	Id    int `gorm:"id" json:"id" form:"id" binding:"required"`
	Num   int `gorm:"num" json:"num" binding:"required"`
	Money int `gorm:"money" json:"money" binding:"required"`
>>>>>>> 81203b5 (选择A，B)
=======
>>>>>>> d73ad65 (基本完成)
}

func (Print) TableName() string {
	return "print"
}
