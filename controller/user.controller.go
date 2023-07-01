package controller

import (
	"lear-jwt/config"
	"lear-jwt/initializers"
	"lear-jwt/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	user := new(models.User)
	var userReq struct {
		Name     string
		Email    string
		Password string
	}

	err := c.Bind(&userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "data required",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to hash password",
		})
		return
	}
	user.Name = userReq.Name
	user.Email = userReq.Email
	user.Password = string(hash)

	err = initializers.DB.Table("users").Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "name alreaady exist",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": "succes created user",
	})

}

func Login(c *gin.Context) {
	user := new(models.User)
	var userReq struct {
		Password string
		Name     string
	}

	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "required",
		})
		return
	}
	err = initializers.DB.Table("users").Where("name = ?", userReq.Name).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid username or password",
		})
		return
	}
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid username or password",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid username or password",
		})
		return
	}

	// generate a jwt token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"user": user.ID,
	// 	"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	// })

	// tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "failed to create token",
	// 	})
	// 	log.Println(err)
	// 	return
	// }

	// making token
	expTime := time.Now().Add(time.Hour * 24)
	claims := &config.JWTClaims{
		Username: userReq.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "lear-jwt",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// algoritma for signing
	tokenAlgthm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	tokenString, err := tokenAlgthm.SignedString(config.Jwt_key)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		log.Println(err)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 24, "", "", false, true)

	c.JSON(200, gin.H{
		"message": "succes login",
	})
}

func LogOut(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(200, gin.H{
		"message": "succes logout",
	})
}
