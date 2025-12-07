package main

import (
	"code/middliwere"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Middleware barcha protected route'larda ishlaydi
    r.Use(middliwere.AuthMiddleware())

    r.GET("/protected", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "success": true,
            "message": "Siz authorized bo‘ldingiz!",
        })
    })

    r.GET("/public", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "success": true,
            "message": "Bu route public",
        })
    })

    r.Run(":8080")
}
