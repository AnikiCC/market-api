package middlewares

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Refresh  bool   `json:"refresh"`
	jwt.RegisteredClaims
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := GetValidatedClaims(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		c.Set("claims", claims)
		return next(c)
	}
}

func GenerateJWT(userId int, username string, refresh bool) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	if refresh {
		expirationTime = time.Now().Add(7 * 24 * time.Hour)
	}

	claims := &Claims{
		UserId:   userId,
		Username: username,
		Refresh:  refresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetValidatedClaims(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, http.ErrAbortHandler
	}

	return claims, nil
}
