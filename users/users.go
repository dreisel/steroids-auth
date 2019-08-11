package users

import (
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Username             string `json:"username"`
	SaltedHashedPassword string
}

func (u UserService) GetByUsernameAndPassword(username string, password string) {

}

func (u UserService) Create(username string, password string) (User, error) {
	user := User{Username: username, SaltedHashedPassword: password}
	result := u.db.Create(&user)
	return user, result.Error
}

func (u UserService) Delete(username string) {

}

func NewUserService(db *gorm.DB) UserService {
	db.AutoMigrate(&User{})
	service := UserService{db: db}
	return service
}
