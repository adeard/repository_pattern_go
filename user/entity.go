package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"size:255;not null;unique" json:"username"`
	Password    string `gorm:"size:255;not null;" json:"password"`
	UserTypeId  int    `gorm:"default:1" json:"user_type_id"`
	UserLevelId int    `gorm:"default:1" json:"user_level_id"`
	Name        string `gorm:"size:255;default:null" json:"name"`
	Email       string `gorm:"size:255;default:null" json:"email"`
	Address     string `gorm:"default:null" json:"address"`
	ImgProfile  string `gorm:"size:255;" json:"img_profile"`
	Gender      string `gorm:"size:255;default:male" json:"gender"`
	BirthDate   string `gorm:"type:date;default:null" json:"birth_date"`
	IsActive    bool   `gorm:"default:true;" json:"is_active"`
}
