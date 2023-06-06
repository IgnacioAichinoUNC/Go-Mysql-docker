package main

import (
	"Aichino/dockergo/controller"
	"Aichino/dockergo/data"
	"Aichino/dockergo/token"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
)


func main() {

	err := data.ConnectionString()
	if err != nil {
		fmt.Println("Error creating database")
	}

	r := gin.Default()
	
	r.POST("/createuser", controller.CreateUser)
	r.GET("//login", controller.Login)

	r.Use(JwtAuthMiddleware())
	r.GET("/listall", controller.ListAll)

	r.Run(":8080") //listen and serve on 0.0.0.0:8080
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokens.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}