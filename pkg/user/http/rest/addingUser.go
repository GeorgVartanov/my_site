package rest

import (
	"github.com/GeorgVartanov/my_site/pkg/user/service/creating"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddingUser(cr creating.UserCreatingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user creating.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := cr.CreateUserService(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "user was created"})
		return

	}
}
