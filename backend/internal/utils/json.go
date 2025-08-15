package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		SendLog("JSON", "BindJSON [ShouldBindJSON]", err)
		SendResponseError(c, http.StatusBadRequest, "", err)
		return false
	}

	return true
}
