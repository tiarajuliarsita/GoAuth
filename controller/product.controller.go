package controller

import "github.com/gin-gonic/gin"

func Store(c *gin.Context) {
	data := []gin.H{
		{
			"id":    1,
			"name":  "Product 1",
			"price": 10,
		},
		{
			"id":    2,
			"name":  "Product 2",
			"price": 20,
		},
	}

	c.JSON(200, data)
}
