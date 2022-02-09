package v1

import (
	"fmt"
	"testing"
)

var tokenString string

func TestSign(t *testing.T) {
	instanceID := "user-1234"
	info := SignInfo{
		Secret:   "!s3@fd&29vnsb213gbvdasd32smz#!bjqa",
		Timeout:  2,
		Issuer:   "auth_gateway",
		Audience: "auth.gateway.com",
	}
	tokenString = Sign(instanceID, info)
}

func TestParse(t *testing.T) {
	fmt.Println(tokenString)
	ParseToken(tokenString)
}
