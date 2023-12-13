package utils

import "github.com/golang-jwt/jwt"

type MyClaim struct {
	jwt.StandardClaims
	Username 	string `json:"Username"`
	Email 		string `json:"Email"`
	AccessUUID	string `json:"AccessUUID"`
}