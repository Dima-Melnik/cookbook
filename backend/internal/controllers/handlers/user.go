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
	if !utils.BindJSON(c, &user) {
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	user.Password = hashedPassword

	if err := db.DB.Create(&user).Error; err != nil {
		utils.SendLog("UserHandlers", "Register", err)
		utils.SendResponseError(c, http.StatusInternalServerError, c.Request.Method, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var user model.User
	var input model.User

	if !utils.BindJSON(c, &input) {
		return
	}

	if err := db.DB.Where("email ILIKE ?", input.Email).First(&user).Error; err != nil {
		utils.SendLog("UserHandlers", "Login", err)
		utils.SendResponseError(c, http.StatusInternalServerError, c.Request.Method, err.Error())
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		utils.SendLog("UserHandlers", "GenerateToken", err)
		utils.SendResponseError(c, http.StatusInternalServerError, "", err.Error())
		fmt.Println("Помилка генерації токена:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetAllUsers(c *gin.Context) {
	var users []model.User

	result := db.DB.Find(&users)
	if result.Error != nil {
		utils.SendLog("UserHandlers", "GetAllUsers", result.Error)
		utils.SendResponseError(c, http.StatusInternalServerError, c.Request.Method, result.Error)
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, "", result)
}
