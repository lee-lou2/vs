package insert_speed_1conn

import (
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(30)"`
	Age  int    `gorm:"type:int"`
}

// Init 초기화
func Init(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.Model(&User{}).Where("1 = 1").Delete(&User{})
}
