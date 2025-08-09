package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		SendResponseError(c, http.StatusBadRequest, "", err)
		return false
	}

	return true
}
