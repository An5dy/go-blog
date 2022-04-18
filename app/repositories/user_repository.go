package repositories

import (
	"go-blog/app/models/user"
	"go-blog/pkg/database"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *user.User) (result *gorm.DB)
	FindByUsername(username string) (userModel user.User, result *gorm.DB)
	Find(user *user.User, id uint64) *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: database.DB,
	}
}

func (ur *userRepository) Create(user *user.User) (result *gorm.DB) {
	return ur.db.Create(user)
}

func (ur *userRepository) FindByUsername(username string) (userModel user.User, result *gorm.DB) {
	result = ur.db.Where("username = ?", username).First(&userModel)
	return
}

func (ur *userRepository) Find(user *user.User, id uint64) *gorm.DB {
	return ur.db.Find(user, id)
}
