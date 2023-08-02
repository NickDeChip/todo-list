package user

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/NickDeChip/todo-list/database"
	"github.com/NickDeChip/todo-list/middleware/auth"
	"github.com/NickDeChip/todo-list/model"
	"github.com/NickDeChip/todo-list/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RedirectRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Use(auth.New())

	app.Post("/signup", PostSignup)
	app.Post("/signin", PostSignIn)

	user.Get("/username", GetUser)
}

func PostSignup(c *fiber.Ctx) error {
	userPostData := model.Signup{}
	err := c.BodyParser(&userPostData)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	if userPostData.Password == "" {
		c.Response().SetBodyString("Password can't be empty")
		return c.SendStatus(http.StatusBadRequest)
	}
	if len(userPostData.Password) < 7 {
		c.Response().SetBodyString("Password can't be less then 7 characters")
		return c.SendStatus(http.StatusBadRequest)
	}
	if userPostData.Email == "" {
		c.Response().SetBodyString("Email can't be empty")
		return c.SendStatus(http.StatusBadRequest)
	}
	if userPostData.Username == "" {
		c.Response().SetBodyString("Username can't be empty")
		return c.SendStatus(http.StatusBadRequest)
	}

	hasedPassword := utility.HashPassword(userPostData.Password)

	res, err := database.GetConnection().Exec("INSERT INTO User (Username, Email, Password) VALUES ($1, $2, $3)", userPostData.Username, userPostData.Email, hasedPassword)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":  id,
			"name": userPostData.Username,
			"iat":  time.Now().Unix(),
		})
	signedTok, err := tok.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   signedTok,
		Expires: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC),
	})
	return c.SendString(string(id))
}

func PostSignIn(c *fiber.Ctx) error {
	userPostData := model.SignIn{}
	err := c.BodyParser(&userPostData)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}

	hashedPassword := utility.HashPassword(userPostData.Password)
	res, err := database.GetConnection().Query("SELECT ID, Email, Username FROM User WHERE Password=$1 AND Email=$2", hashedPassword, userPostData.Email)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
	}
	defer res.Close()

	result := make([]model.User, 0)
	for res.Next() {
		user := model.User{}
		res.Scan(&user.ID, &user.Username, &user.Email)
		result = append(result, user)
	}

	if len(result) > 1 {
		log.Println("!!!!!!!!Found more than one user with matching login!!!!!")
		return c.SendStatus(http.StatusInternalServerError)
	}

	if len(result) == 0 {
		return c.SendStatus(http.StatusUnauthorized)
	}

	loggedInUser := result[0]
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":  loggedInUser.ID,
			"name": loggedInUser.Username,
			"iat":  time.Now().Unix(),
		})
	signedTok, err := tok.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   signedTok,
		Expires: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC),
	})

	return c.SendString(string(loggedInUser.ID))
}

func GetUser(c *fiber.Ctx) error {
	id, err := auth.GetUserID(c)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}
	names, err := database.GetConnection().Query("SELECT Username FROM User WHERE ID=$1", id)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}
	defer names.Close()

	res := model.Username{}

	for names.Next() {
		name := model.Username{}
		names.Scan(&name.Username)
		res = name
	}
	return c.JSON(res)
}
