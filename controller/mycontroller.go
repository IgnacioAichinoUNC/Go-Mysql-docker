package controller

import (
	"Aichino/dockergo/data"
	"Aichino/dockergo/token"
	"Aichino/dockergo/model"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Login(c *gin.Context) {

	var userReq model.User
	if err := c.ShouldBindJSON(&userReq); err != nil {
        print("Error bind REQUEST")
		return
    }

	queryuser := data.GetUser(userReq)

	err := bcrypt.CompareHashAndPassword([]byte(queryuser.Password), []byte(userReq.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Usuario INVALIDO"})
		return
	}

	token, err := tokens.GenerateToken(queryuser.ID)
		if err != nil {
			println("Fail token in Login")
		}

	data := map[string]interface{}{
		"message": "Login success",
		"token":   token,
	}
		

	c.JSON(http.StatusOK, data)
	
}

func CreateUser(c *gin.Context) {
	var newUser model.User

	fmt.Println("Informacion REQUEST: ")
	contentType := c.Request.Header.Get("Content-Type")
	fmt.Println("ContentTYPE: ", contentType)

	if err := c.ShouldBindJSON(&newUser); err != nil {
		println("Error bind REQUEST")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Usuario FAIL"})
		return
	}

	err := data.Insertnewuser(newUser)
	if err != nil {
		print("Error INSERT REQUEST")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Usuario FAIL EN INSERT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario correctamente registrado"})

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

	c.JSON(http.StatusOK, usersResult)
}
