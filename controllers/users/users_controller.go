package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tobypatterson/bookstore_users-api/domain/users"
	"github.com/tobypatterson/bookstore_users-api/services"
	"github.com/tobypatterson/bookstore_users-api/utils/errors"
)

// func TestServiceInterface() {
// services.UsersService
// }

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return userId, nil
}

// Create does something
func Create(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusCreated, result.Marshall(isPublic))
}

// Delete will delete a user
func Delete(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

// Get does something
func Get(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getError := services.UsersService.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusOK, user.Marshall(isPublic))
}

// Search will search for users
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users)

	isPublic := c.GetHeader("X-Public") == "true"
	result := make([]interface{}, len(users))
	for idx, user := range users {
		result[idx] = user.Marshall(isPublic)
	}
	c.JSON(http.StatusOK, result)
}

// Update will update a user
func Update(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, saveErr := services.UsersService.UpdateUser(isPartial, user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusCreated, result.Marshall(isPublic))
}
