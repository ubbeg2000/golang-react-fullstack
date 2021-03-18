package types

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenClaim struct {
	jwt.StandardClaims
	UID  uint64 `json:"uid"`
	Role uint64 `json:"role"`
}
