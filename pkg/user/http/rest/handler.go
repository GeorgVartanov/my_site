package rest

import (
	"github.com/GeorgVartanov/my_site/pkg/user/service/creating"
	"github.com/GeorgVartanov/my_site/pkg/user/service/reading"
	"github.com/gin-gonic/gin"
)

func Handler(cr creating.UserCreatingService, rd reading.UserReadingService) {
	r := gin.Default()
	r.POST("/", AddingUser(cr))
	r.GET("/user", ReadingUser(rd))
	r.Run()
}


