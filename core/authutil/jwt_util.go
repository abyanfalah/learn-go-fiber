package authutil

import (
	"learn-fiber/core/config"
	"learn-fiber/core/config/db"
	"learn-fiber/model"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtConfig() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetEnv("JWT_SECRET"))},
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	return exception.Unauthorized(err.Error())
		// },
	})
}

func createJwt(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"ID":    user.ID,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func BlackListToken(token string, duration time.Duration) error {
	return db.Redis().Set(db.RedisCtx(), token, true, duration).Err()
}

func IsBlacklistedToken(token string) (bool, error) {
	res, err := db.Redis().Exists(db.RedisCtx(), token).Result()
	return res == 1, err
}

func GetJwt(c *fiber.Ctx) string {
	token := c.Cookies(config.GetEnv("COOKIE_NAME"))
	if token == "" {
		return ""
	}

	return token
}
