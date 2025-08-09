package handlers

import (
	"cook_book/backend/internal/auth"
	"cook_book/backend/internal/db"
	"cook_book/backend/internal/model"
	"cook_book/backend/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = hashedPassword

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var user model.User
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DB.Where("email ILIKE ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		fmt.Println("Помилка генерації токена:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("Згенерований токен:", token)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetAllUsers(c *gin.Context) {
	var users []model.User

	result := db.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, users)
}
