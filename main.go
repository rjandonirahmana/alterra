package main

import (
	"mvcApi/config"
	"mvcApi/database"
	"mvcApi/handler"
	"mvcApi/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.NewDb()

	middleWare := middleware.NewMidlleWare()

	repoBook := database.NewRepoBook(db)
	repo := database.NewRepo(db, middleWare)

	handlerBook := handler.NewHandlerBook(repoBook)
	handler := handler.NewHandler(repo)

	c := echo.New()
	c.GET("/users", handler.GetUsers)
	c.POST("/user", handler.CreateUser)
	c.DELETE("/user/:id", handler.DeleteController)
	c.GET("/user/:id", handler.GetUser)
	c.PUT("/user/:id", handler.UpdateUser)

	c.GET("/books", handlerBook.GetBooks)
	c.POST("/book", handlerBook.CreateBook)
	c.DELETE("/book/:id", handlerBook.DeleteBook)
	c.GET("/book/:id", handlerBook.GetBookByID)
	c.PUT("/book/:id", handlerBook.UpdateBook)

	c.Logger.Fatal(c.Start(":1323"))

}

func authMidleware(authService middleware.ServiceIn, user database.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "not allowed",
			})
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "not allowed",
			})
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "not allowed",
			})
		}

		userID := int(claim["userID"].(float64))

		user, err := user.GetUser(userID)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "not allowed",
			})
		}

		c.Set("currentUser", user)
		return nil

	}
}
