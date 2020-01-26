package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tobypatterson/bookstore_users-api/domain/users"
	"github.com/tobypatterson/bookstore_users-api/services"
	"github.com/tobypatterson/bookstore_users-api/utils/errors"
)

// CreateUser does something
func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser does something
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user)
}

// SearchUser will search for users
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
