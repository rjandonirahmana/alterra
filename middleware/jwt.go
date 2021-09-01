package middleware

import (
	"errors"
	model "mvcApi/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Service struct{}

func NewMidlleWare() *Service {
	return &Service{}
}

type ServiceIn interface {
	GenerateToken(userID int) (string, error)
	ExtractTokenUserId(c echo.Context) int
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

func (s *Service) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["Authorization"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(model.SECRET_KEY))
}

func (s *Service) ExtractTokenUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)

	if user.Valid {
		claim := user.Claims.(jwt.MapClaims)
		userID := int(claim["userID"].(float64))
		return userID
	}

	return 0
}

func (s *Service) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("INVALID")
		}

		return []byte(model.SECRET_KEY), nil

	})
	if err != nil {
		return token, err
	}

	return token, nil
}
