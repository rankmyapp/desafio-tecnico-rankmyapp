package response

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func JSONError(c *gin.Context, status int, message string, details string) {
	c.JSON(status, ErrorResponse{
		Message: message,
		Details: details,
	})
}
