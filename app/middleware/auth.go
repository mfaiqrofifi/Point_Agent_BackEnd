package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/labstack/echo/v4/middleware"
)

type JwtMyClaims struct {
	UserId int
	Role   string
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) GenerateToken(UserId int, Role string) (string, error) {
	claims := JwtMyClaims{
		UserId,
		Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}
func (jwtConf *ConfigJWT) Knew() echo.MiddlewareFunc {
	jwt := middleware.JWT(jwtConf.SecretJWT)
	return jwt
}

func GetUser(c echo.Context) (string, int) {
	tmp := c.Get("user")
	if tmp != nil {
		user := tmp.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		Role := claims["Role"].(string)
		Id := claims["UserId"].(float64)
		return Role, int(Id)
	} else {
		return "", 0
	}
}
