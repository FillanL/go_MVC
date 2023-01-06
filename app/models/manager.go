package models

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type DatabaseManager struct{
	DB *gorm.DB
}
type SessionManager struct{
	Session *session.Store
}

var Manager DatabaseManager
func SetManager(dbReference *gorm.DB){
	Manager = DatabaseManager{
		DB: dbReference,
	}
}

var SessionStore SessionManager
func SetSessionManager(sess *session.Store){
	SessionStore = SessionManager{
		Session: sess,
	}
}