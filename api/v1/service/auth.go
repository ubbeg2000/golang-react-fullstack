package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/helper/auth"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/models"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/repo"
)

type AuthService struct {
	repo repo.Repo
}

// Authenticate : authenticate a user
func (s *AuthService) Authenticate(username string, password string) (models.User, string) {
	if user := userService.FindByUsernameOrEmail(username); user.ID != 0 {
		if valid := auth.Verify(password, user.Password); valid {
			token, _ := auth.MakeToken(user)
			return user, token
		}
		return models.User{}, ""
	} else {
		return models.User{}, ""
	}
}

func (s *AuthService) RefreshToken(tokenString string) string {
	claims, err := auth.VerifyToken(tokenString)
	if e, ok := err.(*jwt.ValidationError); ok {
		if e.Errors == jwt.ValidationErrorExpired {
			if time.Now().Unix()-claims.ExpiresAt <= 2*360000 {
				user := userService.FindByID(int(claims.UID))
				token, _ := auth.MakeToken(user)
				return token
			}
			return ""
		}
		return tokenString
	}
	return tokenString
}
