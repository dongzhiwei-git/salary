package models

type SysUser struct {
<<<<<<< HEAD
	Id       int    `gorm:"id" json:"id" form:"id" `
	Num      int    `gorm:"num" json:"num" binding:"required"`
	Password string `gorm:"password" json:"password" binding:"required"`
	PlayA    int  `gorm:"play_a" json:"play_a" `
	PlayB    int  `gorm:"play_b" json:"play_b" `
=======
	Id       int    `gorm:"id" json:"id" form:"id" binding:"required"`
	Num      int    `gorm:"num" json:"num" binding:"required"`
	Password string `gorm:"password" json:"password" binding:"required"`
	PlayA    int  `gorm:"play_a" json:"play_a" binding:"required"`
	PlayB    int  `gorm:"play_b" json:"play_b" binding:"required"`
>>>>>>> 81203b5 (选择A，B)
}

func (SysUser) TableName() string {
	return "user"
}
