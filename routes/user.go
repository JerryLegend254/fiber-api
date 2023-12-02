package routes

import (
	"errors"

	"github.com/JerryLegend254/fiber-api/database"
	"github.com/JerryLegend254/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResUser(user)

	return c.Status(201).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResUser(user)
		responseUsers = append(responseUsers, responseUser)

	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.First(&user, id)
	if user.ID == 0 {
		return errors.New("User does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err1 := findUser(ID, &user); err1 != nil {
		return c.Status(400).JSON(err1.Error())
	}

	responseUser := CreateResUser(user)

	return c.Status(200).JSON(responseUser)

}
