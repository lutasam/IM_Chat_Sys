package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/model"
	"time"
)

type JWTStruct struct {
	UserID         uint64 `json:"user_id"`
	Account        string `json:"account"`
	StandardClaims jwt.StandardClaims
}

func (J JWTStruct) Valid() error {
	return nil
}

// GenerateJWTInUser generates a JWT by username and password
func GenerateJWTInUser(user *model.User) (string, error) {
	timeNow := time.Now().Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = JWTStruct{
		UserID:  user.ID,
		Account: user.Account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeNow + common.EXPIRETIME,
			Issuer:    common.ISSUER,
			NotBefore: timeNow,
		},
	}
	tokenString, err := token.SignedString([]byte(common.JWTSECRETSALT))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWTToken parse tokenString to JWTStruct
func ParseJWTToken(tokenString string) (*JWTStruct, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTStruct{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(common.JWTSECRETSALT), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*JWTStruct), nil
}
