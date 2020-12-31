package controller

import (
	"betalectic_task_1_secure/models"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// GetDetails Gets Details of a User
func GetDetails(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	user := &models.User{}
	coll := mgm.Coll(user)
	err := coll.First(bson.M{"username": username}, user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendString("Welcome " + username + "\nYour Details are : " + user.Details)
}
