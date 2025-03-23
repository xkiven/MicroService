package dao

import (
	"MicroService/kitex_gen/user"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser *user.User) error {
	return db.Create(newUser).Error
}

func GetUserByPasswordAndUsername(db *gorm.DB, password string, username string) (*user.User, error) {
	var user user.User
	result := db.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func CheckUsernameExist(db *gorm.DB, username string) (bool, error) {
	var user user.User
	result := db.Where("username = ?", username).First(&user)
	return result.RowsAffected > 0, result.Error
}
