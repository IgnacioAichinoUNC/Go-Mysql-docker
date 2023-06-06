package controller

import (
	"Aichino/dockergo/data"
	"Aichino/dockergo/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser model.User

	fmt.Println("Informacion REQUEST: ")
	contentType := c.Request.Header.Get("Content-Type")
	fmt.Println("ContentTYPE: ", contentType)

	if err := c.ShouldBindJSON(&newUser); err != nil {
		println("Error bind REQUEST")
		c.JSON(400, gin.H{"message": "Usuario FAIL"})
		return
	}

	fmt.Println("Datos del cuerpo de la solicitud:")
	fmt.Println("Username:", newUser.Username)
	fmt.Println("Password:", newUser.Password)


	err := data.Insertnewuser(newUser)
	if err != nil {
		print("Error INSERT REQUEST")
		c.JSON(404, gin.H{"message": "Usuario FAIL EN INSERT"})
		return
	}

	c.JSON(200, gin.H{"message": "Usuario correctamente registrado"})

}

func ListAll(c *gin.Context) {

	users := data.GetAllUsers()

	type UserResponse struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	var usersResult []UserResponse
	for _, user := range users {
		usersResult = append(usersResult, UserResponse{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	c.JSON(200, usersResult)
}
