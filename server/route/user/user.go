package user

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/NickDeChip/todo-list/database"
	"github.com/NickDeChip/todo-list/model"
	"github.com/NickDeChip/todo-list/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RedirectRoutes(app *fiber.App) {
	app.Post("/signup", PostSignup)
	app.Post("/signin", PostSignIn)
}

func PostSignup(c *fiber.Ctx) error {
	userPostData := model.Signup{}
	err := c.BodyParser(&userPostData)
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusBadRequest)
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

	return c.JSON(model.ID{ID: id})
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
	signedTok, err := tok.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		return utility.LogErrorAndSendStatus(c, err, http.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: signedTok,
	})
	return c.Redirect("http://localhost:3000")
}
