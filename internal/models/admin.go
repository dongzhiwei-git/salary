package models

type SysUser struct {
	Id       int    `gorm:"id" json:"id" form:"id" binding:"required"`
	Num      int    `gorm:"num" json:"num" binding:"required"`
	Password string `gorm:"password" json:"password" binding:"required"`
	PlayA    int  `gorm:"play_a" json:"play_a" binding:"required"`
	PlayB    int  `gorm:"play_b" json:"play_b" binding:"required"`
}

func (SysUser) TableName() string {
	return "user"
}
