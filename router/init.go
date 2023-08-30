package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "test"})
		})
	}
	return r
}
