package auth_utils

import (
	"github.com/gin-gonic/gin"
)

const (
	authUserKey = "auth_user"
)

func SetAuthUser(c *gin.Context, username string) {
	c.Set(authUserKey, username)
}

func GetAuthUser(c *gin.Context) string {
	value, exists := c.Get(authUserKey)
	if !exists {
		return ""
	}

	username, _ := value.(string)
	return username
}