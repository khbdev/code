package middliwere

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware Authorization header tekshiradi
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "message": "Authorization header topilmadi",
            })
            c.Abort()
            return
        }

        // Misol uchun Bearer token tekshiruv
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" || parts[1] != "your-secret-token" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "message": "Authorization header noto‘g‘ri",
            })
            c.Abort()
            return
        }

        // Header to‘g‘ri bo‘lsa, handlerga o‘tkazish
        c.Next()
    }
}
