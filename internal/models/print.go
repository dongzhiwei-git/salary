package models

type Print struct {
	ID       int    `gorm:"id" json:"id"`
	UserName string `gorm:"user_name" json:"user_name" binding:"required"`
	Scheme   string `json:"scheme" json:"scheme" binding:"required"`
	Num      int    `gorm:"num" json:"num" binding:"required"`
	Money    int    `gorm:"money" json:"money" binding:"required"`
}

func (Print) TableName() string {
	return "print"
}
