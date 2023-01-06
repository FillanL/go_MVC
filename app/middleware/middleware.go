package middleware

import (
	"errors"

	"github.com/FillanL/creatturlinks/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Middlewares struct{
	
}
func IsAuthenticated(c *fiber.Ctx) error {
	// Set a custom header on all responses:
	// c.Set("X-Custom-Header", "Hello, World")
  
	// Go to next middleware:
	return c.Next()
}

// func authGenerateAccessClaim(id string) *jwt.Token{
// 	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id": id,
// 		"exp": time.Now().Add(time.Minute * 10).Unix(),
// 		"iat": time.Now().Unix(),
// 	})
// 	return claim
// }
// func authGenerateRefreshClaim(id string) *jwt.Token{
// 	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id": id,
// 		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
// 		"iat": time.Now().Unix(),
// 	})
// 	return claim
// }

// func AuthEncodeAccessToken(id string) (string,  error){
// 	claim := authGenerateAccessClaim(id)
// 	secret := config.Env.JWTSecret
// 	token, err := claim.SignedString([]byte(secret))
// 	if err != nil{
// 		return "", err
// 	}
// 	return token, nil
// }
// func AuthEncodeRefreshToken(id string) (string,  error){
// 	claim := authGenerateRefreshClaim(id)
// 	secret := config.Env.JWTSecret
// 	token, err := claim.SignedString([]byte(secret))
// 	if err != nil{
// 		return "", err
// 	}
// 	return token, nil
// }
/*
Set a new active session in cache
*/
func SetSessionId(userId string, c *fiber.Ctx) (*session.Session, error){

	sess, sessionErr := models.SessionStore.Session.Get(c)
	if sessionErr != nil{
		return nil, errors.New("something went wrong")
	}

	sess.Set("auth", true)
	sess.Set("sid", userId)

	sessionErr = sess.Save()
	if sessionErr != nil{
		return nil, errors.New("something went wrong" + sessionErr.Error())
	}
	return sess, nil
}
/*
remove active session from cache
*/
func RemoveSession(c *fiber.Ctx) error{
	sess, sessErr := models.SessionStore.Session.Get(c)
	if sessErr != nil{
		return sessErr
	}
	sess.Delete("auth")
	sess.Delete("sid")
	sess.Destroy()
	return nil
}
/*
checks if session is still active
*/
func IsSessionActive(c *fiber.Ctx) error{
	
	sess, sessErr := models.SessionStore.Session.Get(c)
	if sessErr != nil{
		return sessErr
	}
	if sess.Get("auth") == nil{
		return errors.New("failed: not auth")
	}
	return nil
}