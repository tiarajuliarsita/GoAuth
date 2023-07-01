package middleware

import (
	"lear-jwt/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized1"})
				return
			}else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
				return
			}
		}
		
		claims := &config.JWTClaims{}
		
		// parse
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.Jwt_key, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)

			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "unauthorized2",
				})
				return
			case jwt.ValidationErrorExpired:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "unauthorized, token exp",
				})
				return
				
			default:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "unauthorized, token exp",
				})
				return

			}

		}

		if !token.Valid{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized, token exp",
			})
			return
		}
		c.Next()
	}
}
