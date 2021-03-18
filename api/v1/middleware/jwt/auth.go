package middleware

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/helper/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	AllowedRoles []uint64
	Response401  map[string]interface{}
	Response403  map[string]interface{}
	RefreshFunc func(tokenString string) string
}

var DefaultConfig Config = Config{
	AllowedRoles: []uint64{},
	Response401: map[string]interface{}{
		"code": 401,
		"msg":  "Unauthorized, please authenticate",
	},
	Response403: map[string]interface{}{
		"code": 403,
		"msg":  "Forbidden, you can't access this resource",
	},
	RefreshFunc: defaultRefresh
}

// Middleware init
func New(conf Config) fiber.Handler {
	config := makeConfig(conf)
	return func(c *fiber.Ctx) error {

		// Check header content existence
		if token := c.Request().Header.Peek("Authorization"); content == nil {

			c.SendStatus(401)
			return c.JSON(config.Response401)

		} else {
			// Check token existence
			header := strings.Split(string(content), " ")
			if len(header) < 2 {

				c.SendStatus(401)
				return c.JSON(config.Response401)

			} else {
				// Check token validity
				tokenString := header[1]
				claim, err := auth.VerifyToken(tokenString)

				if e, ok := err.(*jwt.ValidationError); ok {
					
					// Check authorization
					if e == nil {
						for _, roleID := config.AllowedRoles {
							if roleID == claims.Role {
								return c.Next()
							}
						}

						c.SendStatus(403)
						return c.JSON(conf.Response403)
					}

					// Check token expiration
					if e.Errors == jwt.ValidationErrorExpired {
						// Refresh token if possible
						if newToken := conf.RefreshFunc(tokenString); newToken != "" {
							c.SendStatus(202)
							return c.JSON(jwt.Map{
								"token": newToken,
								"msg": "Token expired, here's something new",
							})
						}
						c.SendStatus(401)
						return c.JSON(config.Response401)

					} else {
						c.SendStatus(401)
						return c.JSON(config.Response401)
					}
				}
			}
		}
	}
}

// Default functions
func makeConfig(c ...Config) Config {
	if len(c) < 1 {
		return DefaultConfig
	}

	conf := c[0]

	if conf.AllowedRoles == nil {
		conf.AllowedRoles = DefaultConfig.AllowedRoles
	}

	if conf.Response401 == nil {
		conf.Response401 = DefaultConfig.Response401
	}

	if conf.Response403 == nil {
		conf.Response403 = DefaultConfig.Response403
	}

	if conf.RefreshFunc == nil {
		conf.RefreshFunc = DefaultConfig.RefreshFunc
	}

	return conf
}

// Default token refresh function
func defaultRefresh(tokenString string) string {
	s := service.New()
	return s.Auth.RefreshToken(tokenString)
}