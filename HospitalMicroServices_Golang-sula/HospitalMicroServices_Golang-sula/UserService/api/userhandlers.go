package api

import (
	"fmt"
	"github.com/Fring02/HospitalMicroservices/UserService/core"
	"github.com/Fring02/HospitalMicroservices/UserService/pkg/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

var UserRepository repositories.UserRepository

func RouteUsers(router *gin.Engine) {
	router.GET("/users", GetUsers)
	router.POST("/users", CreateUser)
	router.GET("/users/:id", GetUsersByID)
	router.DELETE("/users/:id", DeleteUser)
	router.PUT("/users/:id", UpdateUser)
	router.POST("/users/login", UserLogin)
}

func GetUsers(c *gin.Context) {
	users := UserRepository.GetUsers()
	c.JSON(200, users)
}

func GetUsersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	user := UserRepository.GetUserByID(id)
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	user := &core.User{}
	err := c.BindJSON(user)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
	}
	id, err := UserRepository.CreateUser(*user)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to create a user"))
	} else {
		token, err := CreateToken(id)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, token)
	}

}

func DeleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	user := UserRepository.GetUserByID(id)
	if user == nil {
		c.Data(400, jsonContentType, []byte("No such user in database"))
		return
	}
	_, err = UserRepository.DeleteUser(user.ID)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to delete the user"))
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The user with id %v has been successfully deleted", id)))
	}

}

func UpdateUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	user := UserRepository.GetUserByID(id)

	if user == nil {
		c.Data(400, jsonContentType, []byte("There is no such user"))
		return
	}

	updatedUser := &core.User{}
	err = c.BindJSON(updatedUser)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	updatedUser.ID = id
	_, err = UserRepository.UpdateUser(*updatedUser)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to update the user"))
		return
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The user with id %v has been successfully updated", id)))
	}

}

func UserLogin(c *gin.Context) {
	userGet := &core.User{}
	err := c.BindJSON(userGet)

	user := UserRepository.GetUser(userGet.Email, userGet.Password)
		if user == nil {
			c.JSON(http.StatusUnauthorized, "Please provide valid login details")
			return
		}

	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}


func CreateToken(id int) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}


