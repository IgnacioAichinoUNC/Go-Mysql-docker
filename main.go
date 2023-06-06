package main

import (
	"Aichino/dockergo/controller"
	"Aichino/dockergo/data"
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

	
	r.GET("/listall", controller.ListAll)
	r.POST("/createuser", controller.CreateUser)

	r.Run(":8080") //listen and serve on 0.0.0.0:8080
}

