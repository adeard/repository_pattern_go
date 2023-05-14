package registry

import (
	"gin_web_api/user"

	"gorm.io/gorm"
)

func UserRegistry(db *gorm.DB) user.Service {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	return userService
}
