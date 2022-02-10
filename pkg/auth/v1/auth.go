package v1

import (
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type MyClaims struct {
	InstanceID string `json:"instanceID"`
	jwt.StandardClaims
}

func Sign(instanceID, audience, issuer, secret string, timeout int) string {

	claims := MyClaims{
		instanceID,
		jwt.StandardClaims{
			Audience:  audience,
			ExpiresAt: time.Now().Add(time.Duration(timeout) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
			NotBefore: time.Now().Add(0).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString
}

func ParseToken(tokenString, secret string) (bool, *MyClaims) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		log.Errorf("jwt.ParseWithClaims failed: %v", err)
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return true, claims
	} else {
		return false, claims
	}
}
