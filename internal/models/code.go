package models

type Code struct {
	Id int `gorm:"id" json:"id"`
	Code string `gorm:"code" json:"code" binding:"required"`
}

func (Code) TableName() string {
	return "code"
}
