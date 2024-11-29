package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	// Initialize the database
	InitDB()
	// Define the routes
	e.GET("/users", GetUsers)
	e.POST("/users", CreateUser)
	e.GET("/users/:id/bank-accounts", GetUserBankAccounts)
	e.POST("/users/:id/bank-accounts", CreateUserBankAccount)
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

// GetUserBankAccounts handler function
func GetUserBankAccounts(c echo.Context) error {
	userID := c.Param("id")
	var bankAccounts []BankAccount
	if err := DB.Where("user_id = ?", userID).Find(&bankAccounts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, bankAccounts)
} // CreateUserBankAccount handler function
func CreateUserBankAccount(c echo.Context) error {
	userID := c.Param("id")
	var bankAccount BankAccount
	if err := c.Bind(&bankAccount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	bankAccount.UserID = uint(id)
	if err := DB.Create(&bankAccount).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, bankAccount)
}
