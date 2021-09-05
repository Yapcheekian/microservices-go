package users

import (
	"net/http"
	"strconv"

	"github.com/Yapcheekian/microservices-go/users/domain/users"
	"github.com/Yapcheekian/microservices-go/users/services"
	"github.com/Yapcheekian/microservices-go/users/utils/errors"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Update(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))

	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Create(c *gin.Context) {
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

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))

	if err != nil {
		c.JSON(err.Code, err)
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

func getUserId(userParam string) (int64, *errors.RestErr) {
	userId, err := strconv.ParseInt(userParam, 10, 64)

	if err != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}

	return userId, nil
}
