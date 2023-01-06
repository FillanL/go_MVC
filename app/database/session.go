package database

import (
	"time"

	"github.com/FillanL/creatturlinks/app/models"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

type Session struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserRole string `json:"user_role"`
	Expiration time.Time `json:"expiration"`
}

func NewStorage() *session.Store {
	store := redis.New(
		redis.Config{
			Host:"localhost",
			Port:6379,
			Password:"",
			URL:"",
			Database:0,
			Reset: false,
			TLSConfig: nil,
		},
	)
	sess := session.New(session.Config{
		Storage: store,
		CookieHTTPOnly: true,
		Expiration: time.Hour * 5,
		KeyLookup: "cookie:sid",
	})
	models.SetSessionManager(sess)
	return sess
}