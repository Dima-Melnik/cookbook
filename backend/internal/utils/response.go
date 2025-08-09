package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Type    string      `json:"type"`
	Method  string      `json:"method,omitempty"`
	Message interface{} `json:"message"`
}

func SendResponseError(c *gin.Context, status int, method string, err any) {
	resp := Response{
		Type:    "error",
		Method:  method,
		Message: err,
	}

	c.JSON(status, resp)
	c.Abort()
}

func SendResponseJSON(c *gin.Context, status int, method string, response any) {
	resp := Response{
		Type:    "success",
		Method:  method,
		Message: response,
	}

	c.JSON(status, resp)
}
