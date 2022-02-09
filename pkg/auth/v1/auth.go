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

type SignInfo struct {
	// Secret 签发密钥
	Secret string
	// Timeout 有效期 hour
	Timeout int
	// Issuer 签发者 "iamctl"
	Issuer string
	// Audience 接收者 "iam.authz.marmotedu.com",
	Audience string
}

func Sign(instanceID string, info SignInfo) string {

	claims := MyClaims{
		instanceID,
		jwt.StandardClaims{
			Audience:  info.Audience,
			ExpiresAt: time.Now().Add(time.Duration(info.Timeout) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    info.Issuer,
			NotBefore: time.Now().Add(0).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(info.Secret))

	return tokenString
}

func ParseToken(tokenString string) (bool, *MyClaims) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("!s3@fd&29vnsb213gbvdasd32smz#!bjqa"), nil
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
