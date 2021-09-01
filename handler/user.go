package handler

import (
	"fmt"
	"mvcApi/database"
	model "mvcApi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo database.Repository
}

func NewHandler(repo database.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})

}

func (h *Handler) CreateUser(c echo.Context) error {
	var user model.Users
	c.Bind(&user)

	user, err := h.repo.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := model.Users{}
	c.Bind(&user)
	user1, err := h.repo.UpdateUser(id, &user)

	if err != nil {
		fmt.Println("disini", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "oke",
		"updated user": user1,
	})

}

func (h *Handler) DeleteController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := model.Users{}

	if err = h.repo.DeleteUser(id, user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to delete user in handler",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success delete user with id %d", id),
	})

}

func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.repo.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get user in handler",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get one user with id %d", id),
		"users":    user,
	})

}

// func (h *Handler) Login(c echo.Context) error {
// 	var user model.Users
// 	c.Bind(&user)

// 	err :=  h.repo.LoginUser(&user)

// }
