package users

import (
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Username             string `gorm:"unique;not null"`
	SaltedHashedPassword string `gorm:"not null"`
}

func (u UserService) GetByUsernameAndPassword(username string, password string) (User, error) {
	user := User{Username: username, SaltedHashedPassword: hashAndSalt(password)}
	result := u.db.Find(&user)
	return user, result.Error
}

func (u UserService) Create(username string, password string) (User, error) {
	user := User{Username: username, SaltedHashedPassword: hashAndSalt(password)}
	result := u.db.Create(&user)
	return user, result.Error
}

func (u UserService) Delete(id uint) (User, error) {
	user := User{
		Model:                gorm.Model{ID: id},
		Username:             "",
		SaltedHashedPassword: "",
	}
	result := u.db.Delete(&user)
	return user, result.Error
}

func NewUserService(db *gorm.DB) UserService {
	db.AutoMigrate(&User{})
	service := UserService{db: db}
	return service
}
