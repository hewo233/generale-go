package common

import (
	"generale-go/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("Remember_to_change_KEY")

type Claim struct {
	UserUid string
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(1 * 24 * time.Hour) //token's expiration Time

	claims := &Claim{
		UserUid: user.Uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errTS := token.SignedString(jwtKey)
	if errTS != nil {
		return "", errTS
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claim, err
}
