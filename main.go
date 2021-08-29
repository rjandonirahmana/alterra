package main

import (
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:12345@tcp(127.0.0.1:3306)/alterra?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db1 := NewDB(db)

	c := echo.New()
	c.POST("/users", db1.CreateUserController)
	c.GET("/users", db1.GetUsersController)
	c.GET("/user/:id", db1.GetUserController)
	c.DELETE("/user/:id", db1.DeleteUserController)
	c.PUT("/user/:id", db1.UpdateUserController)

	c.Logger.Fatal(c.Start(":1323"))

}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type dbstruct struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) *dbstruct {
	return &dbstruct{db: db}
}

func (d *dbstruct) GetUsersController(c echo.Context) error {

	users := []User{}

	err := d.db.Find(&users).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})

}

func (db *dbstruct) GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var user User

	db.db.Find(&user, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get one user with id %d", id),
		"users":    user,
	})

}

func (db *dbstruct) CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if err := db.db.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func (d *dbstruct) DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := User{}
	users := []User{}
	d.db.Where("id = ?", id).Delete(&user)
	d.db.Find(&users)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":   fmt.Sprintf("success delete user with id %d", id),
		"users left": users,
	})

}

func (d *dbstruct) UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := []User{}

	tx := d.db.Find(&user, id)

	if tx.Error != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if tx.RowsAffected > 0 {
		newUser := User{}

		c.Bind(&newUser)

		err2 := d.db.Model(&user).Updates(newUser).Error

		if err2 != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "user update failed",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages":     fmt.Sprintf("success update user with id %d", id),
			"updated user": user,
		})

	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "user update failed",
	})
}
