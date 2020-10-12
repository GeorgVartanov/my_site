package rest

import (
	"github.com/GeorgVartanov/my_site/pkg/user/service/reading"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadingUser(rd reading.UserReadingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var email reading.UserEmail
		if err := c.ShouldBindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := rd.ReadUserService(email.EmailName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": user})
		return

	}
}

//func ReadingAll(rd read.ServiceReader) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		users, err := rd.ReadAll()
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			json.NewEncoder(w).Encode(err.Error())
//			return
//		}
//		json.NewEncoder(w).Encode(users)
//	}
//
//}
