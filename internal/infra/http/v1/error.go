package v1

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"exceptions"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}
