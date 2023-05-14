package user

import "gorm.io/gorm"

type Repository interface {
	FindByID(ID int) (User, error)
	FindByUsername(username string) (User, error)
	Create(user User) (User, error)
	// Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).First(&user).Error

	return user, err
}

func (r *repository) FindByUsername(username string) (User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error

	return user, err
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

// func (r *repository) Update(user User) (User, error) {
// 	err := r.db.Save(&user).Error

// 	return user, err
// }
