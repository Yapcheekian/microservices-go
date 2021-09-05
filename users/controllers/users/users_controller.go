package users

import (
	"net/http"
	"strconv"

	"github.com/Yapcheekian/microservices-go/users/domain/users"
	"github.com/Yapcheekian/microservices-go/users/services"
	"github.com/Yapcheekian/microservices-go/users/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	result, err := services.CreateUser(user)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		restErr := errors.NewBadRequestError("invalid user id")
		c.JSON(restErr.Code, restErr)
		return
	}

	user, userErr := services.GetUser(userId)

	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {

}
