package models

type Print struct {
	Id    int `gorm:"id" json:"id" form:"id" binding:"required"`
	Num   int `gorm:"num" json:"num" binding:"required"`
	Money int `gorm:"money" json:"money" binding:"required"`
}

func (Print) TableName() string {
	return "print"
}
