package handlers

import (
	"cook_book/backend/internal/model"
	"cook_book/backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllRecipes(c *gin.Context) {
	// userID, err := helper.GetUserID(c)
	// if err != nil {
	// 	return
	// }

	result, err := h.cookService.GetAll()
	if err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, c.Request.Method, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, c.Request.Method, result)
}

func (h *Handler) GetRecipeByID(c *gin.Context) {
	id, err := utils.CheckCorrectID(c)
	if err != nil {
		return
	}

	result, err := h.cookService.GetByID(id)
	if err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, c.Request.Method, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, c.Request.Method, result)
}

func (h *Handler) CreateRecipe(c *gin.Context) {
	var createdRecipe model.CreateCookBook

	if !utils.BindJSON(c, &createdRecipe) {
		return
	}

	// userID, err := helper.GetUserID(c)
	// if err != nil {
	// 	return
	// }

	recipe := model.CookBook{
		Title:       createdRecipe.Title,
		Description: createdRecipe.Description,
		// UserID:      userID,
	}

	if err := h.cookService.Create(&recipe); err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, c.Request.Method, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusCreated, c.Request.Method, "Successfully created")
}

func (h *Handler) UpdateRecipe(c *gin.Context) {
	id, err := utils.CheckCorrectID(c)
	if err != nil {
		return
	}

	var updatedRecipe model.UpdateCookBook

	if !utils.BindJSON(c, &updatedRecipe) {
		return
	}

	recipe := model.CookBook{
		Title:       updatedRecipe.Title,
		Description: updatedRecipe.Description,
	}

	if err := h.cookService.Update(&recipe, id); err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, c.Request.Method, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, c.Request.Method, "Seccessfully updated")
}

func (h *Handler) DeleteRecipe(c *gin.Context) {
	id, err := utils.CheckCorrectID(c)
	if err != nil {
		return
	}

	if err := h.cookService.Delete(id); err != nil {
		utils.SendResponseError(c, http.StatusInternalServerError, c.Request.Method, err.Error())
		return
	}

	utils.SendResponseJSON(c, http.StatusOK, c.Request.Method, "Successfully deleted")
}
