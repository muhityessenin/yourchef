package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"user_service/internal/domain/user"
	"user_service/internal/service/interfaces"
	"user_service/pkg"
)

type UserHandler struct {
	UserService interfaces.UserService
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	var token string
	var err error
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := u.UserService.Login(input.Email, input.Password)
	if err != nil {
		if res == http.StatusNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		if res == http.StatusUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == http.StatusOK {
		output, err := u.UserService.SearchUser("email", input.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		token, err = user.GenerateToken(&output[0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": token})
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	var response pkg.Response
	res, err := u.UserService.GetUsers()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response = pkg.Response{Status: http.StatusNotFound, Data: "", Message: "no users found"}
			c.JSON(http.StatusNotFound, response)
			return
		}
		response = pkg.Response{
			Status:  http.StatusServiceUnavailable,
			Message: err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}
	response = pkg.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	}
	c.JSON(http.StatusOK, res)
}
func (u *UserHandler) CreateUser(c *gin.Context) {
	var input user.InputResponse
	var res pkg.Response
	if err := c.BindJSON(&input); err != nil {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if input.ValidateUserInput(&input) == false {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input1"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := u.UserService.CreateUser(&input)
	if err != nil {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input2"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = pkg.Response{Status: http.StatusCreated, Data: "", Message: "success"}
	c.JSON(http.StatusCreated, res)
}
func (u *UserHandler) GetUserById(c *gin.Context) {
	var response pkg.Response
	res, err := u.UserService.GetUserById(c.Param("id"))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response = pkg.Response{Status: http.StatusNotFound, Data: "", Message: "user not found"}
			c.JSON(http.StatusNotFound, response)
			return
		}
		response = pkg.Response{Status: http.StatusServiceUnavailable, Data: "", Message: "invalid input"}
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}
	response = pkg.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	}
	c.JSON(http.StatusOK, response)
}
func (u *UserHandler) UpdateUser(c *gin.Context) {
	var res pkg.Response
	var input user.InputResponse
	if err := c.BindJSON(&input); err != nil {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if input.ValidateUserInput(&input) == false {
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	_, err := u.UserService.UpdateUser(c.Param("id"), &input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res = pkg.Response{Status: http.StatusNotFound, Data: "", Message: "user not found"}
			c.JSON(http.StatusNotFound, res)
			return
		}
		res = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "invalid input2"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = pkg.Response{Status: http.StatusOK, Data: "", Message: "success"}
	c.JSON(http.StatusOK, res)
}
func (u *UserHandler) DeleteUser(c *gin.Context) {
	var res pkg.Response
	_, err := u.UserService.DeleteUser(c.Param("id"))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res = pkg.Response{Status: http.StatusNotFound, Data: "", Message: "user not found"}
			c.JSON(http.StatusNotFound, res)
			return
		}
		res = pkg.Response{Status: http.StatusServiceUnavailable, Data: "", Message: "invalid input"}
		c.JSON(http.StatusServiceUnavailable, res)
		return
	}
	res = pkg.Response{Status: http.StatusOK, Data: "", Message: "success"}
	c.JSON(http.StatusOK, res)
}
func (u *UserHandler) SearchUser(c *gin.Context) {
	var response pkg.Response
	var res []user.Entity
	var err error
	check := false
	name := c.Query("name")
	email := c.Query("email")
	if name != "" {
		res, err = u.UserService.SearchUser("name", name)
		check = true
	} else if email != "" {
		res, err = u.UserService.SearchUser("email", email)
		check = true
	}
	if check == false {
		response = pkg.Response{Status: http.StatusBadRequest, Data: "", Message: "bad request"}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response = pkg.Response{Status: http.StatusNotFound, Data: "", Message: "no users found"}
			c.JSON(http.StatusNotFound, response)
			return
		}
		response = pkg.Response{Status: http.StatusServiceUnavailable, Data: "", Message: "invalid input"}
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}
	response = pkg.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	}
	c.JSON(http.StatusOK, response)
}
