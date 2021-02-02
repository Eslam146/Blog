package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"eslam" : "asdfgh",
		"mohamed" : "123456",
		"admin" : "abc123",
	})
}