package controller

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FillanL/creatturlinks/app/config"
	"github.com/FillanL/creatturlinks/app/middleware"
	"github.com/FillanL/creatturlinks/app/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleConfig *oauth2.Config
var randomState string

type GoogleUser struct{
	Id 				string
	Email 			string
	Verified_email 	bool
	Picture 		string 
}
type GoogleHandlers struct{}

// not used
// func (gh GoogleHandlers) AuthSignUpUser(c *fiber.Ctx) error {
// 	var user models.User

// 	if err := c.BodyParser(&user); err != nil{
// 		return gernerateJsonError("invalid",c,fiber.StatusInternalServerError)
// 	}
	
// 	if err := models.CreateUser(&user); err != nil{
// 		return gernerateJsonError(err.Error(), c, fiber.StatusNotAcceptable)
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(
// 		fiber.Map{
// 			"Ok": true,
// 		},
// 	)
// }

func (gh GoogleHandlers)  AuthLogoutUser(c *fiber.Ctx) error{
	
	middleware.RemoveSession(c)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "user logged out",
		},
	)
}
// not in use
// func (gh GoogleHandlers)  AuthisUserAuthenticated(c *fiber.Ctx) error{
// 	err := models.AuthIsSessionActive(c)
// 	if err != nil{
// 		return gernerateJsonError(err.Error(), c, fiber.StatusUnauthorized)
// 	}
	 
// 	userEmail := fmt.Sprintf("%s",c.Locals("user"))
// 	user,err := models.UserGetUserByEmail(userEmail)
// 	c.Locals("user", nil)
// 	if err != nil{
// 		return gernerateJsonError(err.Error(), c, fiber.StatusUnauthorized)
// 	}
// 	return c.Status(fiber.StatusOK).JSON(
// 		fiber.Map{
// 			"status":  "success",
// 			"message": "authorized to access this resource!",
// 			"user": user,
// 		})
// }


func (gh GoogleHandlers) Login(c *fiber.Ctx) error{
	fmt.Println("google login")
	clientId := config.Env.CLIENT_ID
	clientSecret := config.Env.ClIENT_SECRET

	googleConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/api/v1/auth/google/callback",
		Scopes:       []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint:     google.Endpoint,
	}
	randomState = createRandState()
	url := googleConfig.AuthCodeURL(randomState)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}


func (gh GoogleHandlers) Callback(c *fiber.Ctx) error{
	userModel := models.User{}

	const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	code := c.Query("code")
	state := c.Query("state")
	fmt.Println("state is :" +state, randomState)
	if state != randomState{
		return gernerateJsonError("wrong response due to state", c, fiber.StatusBadRequest)
	}
	
	token, err := googleConfig.Exchange(c.UserContext(), code)
	if err != nil{
		return gernerateJsonError("unable to get t", c, fiber.StatusBadRequest)
	}

	url := OauthGoogleUrlAPI+token.AccessToken
	res, err := http.Get(url)
	if err != nil{
		return gernerateJsonError("google services are down", c, fiber.StatusBadGateway)
	}

	defer res.Body.Close()
	var g GoogleUser
	err = json.NewDecoder(res.Body).Decode(&g)
	if err != nil{
		gernerateJsonError("google services are down", c, fiber.StatusBadRequest)
	}
	if !g.Verified_email{
		return gernerateJsonError("only accept verified emails", c, fiber.StatusBadRequest)
	}
	// add user to db
	user := models.User{
		Email: g.Email,
		Username: g.Email,
		Password: g.Id,
	}

	err = userModel.CreateUser(&user)
	if err != nil{
		fmt.Println("user already exist in DB")
	}

	userId := fmt.Sprint(user.Email)

	_, err = middleware.SetSessionId(userId, c)
	if err != nil{
		gernerateJsonError(err.Error(), c, fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message":"ok",
		},
	)
}

func createRandState() string{
	b := make([]byte, 20)
	rand.Read(b)
	randomState := base64.URLEncoding.EncodeToString(b)
	return randomState
}