package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Link struct{
	Id 			uuid.UUID	`json:"-" gorm:"type:uuid; unique; primaryKey"`
	Title 		string		`json:"title" gorem:"not null"`
	Url 		string		`json:"url" gorem:"not null"`
	UserId 		uuid.UUID	`json:"-"`
	User		User		`json:"" gorm:"foreignKey:UserId"` 
	CreatedAt 	time.Time	`json:"" gorem:"not null;autoCreateTime"`
}

func(l *Link) CreateLink()error{
	newLink := Link{
		Title: "",
		Url:"",
		CreatedAt: time.Now(),
	}
	Manager.DB.Create(newLink)
	return errors.New("")
}