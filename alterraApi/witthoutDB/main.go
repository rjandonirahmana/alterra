package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	c := echo.New()
	c.POST("/users", CreateUserController)
	c.GET("/users", GetUsersController)
	c.GET("/user/:id", GetUserController)
	c.DELETE("/user/:id", DeleteUserController)
	c.PUT("/user/:id", UpdateUserController)

	c.Logger.Fatal(c.Start(":1323"))

}

type User struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})

}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if id > len(users)-1 {
		return c.JSON(http.StatusBadRequest, nil)

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get one user with id %d", id),
		"users":    users[id-1],
	})

}

func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.ID = 1
	} else {
		newId := users[len(users)-1].ID + 1
		user.ID = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if id > len(users)-1 {
		return c.JSON(http.StatusBadRequest, "id tidak ditemukan")

	}

	RemoveIndex(users, id-1)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get delete user with id %d", id),
		"users":    users,
	})

}

func RemoveIndex(s []User, index int) []User {
	return append(s[:index], s[index+1:]...)
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := users[id-1]
	user.Name = fmt.Sprintf("name change with id : %d", id)
	user.Email = fmt.Sprintf("new email@%d", id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user account success updated",
		"user":    user,
	})
}
