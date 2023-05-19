package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", FetchAllUsers)
		v1.POST("/users", CreateUser)
		v1.GET("/users/:id", FetchSingleUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	r.Run()
}
