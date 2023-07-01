package config

import "github.com/golang-jwt/jwt/v4"

var Jwt_key = []byte("haudgutghdvhagdyuahydgwy")

type JWTClaims struct {
	Username string
	jwt.RegisteredClaims
}
