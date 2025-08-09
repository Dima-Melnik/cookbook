package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (uint, error) {
	uid, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("user_id not found")
	}

	userID, ok := uid.(uint)
	if !ok {
		return 0, errors.New("invalid userID type")
	}

	return uint(userID), nil
}
