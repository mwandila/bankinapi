package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// Initialize the database
	InitDB()
	// Define the routes
	e.GET("/users", GetUsers)
	e.POST("/users", CreateUser)
	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

// GetUsers handler function
func GetUsers(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

// CreateUser handler function
func CreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}
