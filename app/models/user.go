package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	Id			uuid.UUID	`json:"id" gorm:"type:uuid;primary_key"`
	Email		string		`json:"email" gorem:"unique;not null"`
	Username	string		`json:"username" gorem:"unique"`
	Password	string		`json:"-" gorem:"not null"`
}

func (u *User) CreateUser(user *User)error{
	// model := User{}
	if doesEmailExist := u.EmailExist(user.Email); doesEmailExist {
		return errors.New("email already exist")
	}
	user.Id = uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	user.Password = string(hashedPassword)
	tx := Manager.DB.FirstOrCreate(&user)
	if tx.Error != nil{
		return tx.Error
	}
	return nil
}

func (u *User)EmailExist(email string) bool{
	var user User

	Manager.DB.Where("email = ?", email).First(&user)
	
	fmt.Println(user.Email)
	return len(user.Email) > 0
}

func (u *User)GetUsers() []User{
	var user[] User

	Manager.DB.Find(&user)
	return user
}