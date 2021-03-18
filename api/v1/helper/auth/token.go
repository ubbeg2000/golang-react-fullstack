package auth

import (
	"time"

	"github.com/ubbeg2000/golang-react-fullstack/api/v1/helper/types"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/models"
)

// MakeToken : make a token
func MakeToken(user models.User) (string, error) {
	claim := types.TokenClaim{
		UID:  user.ID,
		Role: user.Role.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(viper.GetInt(`application.jwt.expiration`))*3600000,
			Issuer:    viper.GetString(`application.name`),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(viper.GetString(`application.jwt.secret`)))
}

// VerifyToken : verify a token
func VerifyToken(tokenString string) (types.TokenClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &types.TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString(`application.jwt.secret`)), nil
	})

	claims, _ := token.Claims.(types.TokenClaim)

	return claims, err
}
