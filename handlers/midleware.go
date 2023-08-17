package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func (a *AuthHandler) UserIdentity(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "authorization header is empty",
		})
		logrus.Errorf("authorization header is empty, user: %s", c.Param("username"))
		return errors.New("authorization header is empty")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "authorization header is invalid",
		})
		logrus.Errorf("authorization header is invalid, user: %s", c.Param("username"))
		return errors.New("authorization header is invalid")
	}

	userId, userRole, err := ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		logrus.Errorf("token can not be parsed")
		return errors.New(err.Error())
	}

	c.Set("userId", userId)
	c.Set("userRole", userRole)

	return nil
}

func ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *TokenClaims")
	}

	return claims.UserId, claims.Role, nil

}
