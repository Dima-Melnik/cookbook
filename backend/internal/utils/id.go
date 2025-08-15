package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckCorrectID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			SendResponseError(c, http.StatusNotFound, "", err.Error())
			return uint(id), err
		}

		SendLog("ID", "CheckCorrectID", err)
		SendResponseError(c, http.StatusBadRequest, "", err.Error())
		return uint(id), err
	}

	return uint(id), nil
}
