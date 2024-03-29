package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/ahargunyllib/task-5-pbi-btpns-Nugraha_Billy_Viandy/helpers"
    "net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(context *gin.Context) {
        err := helpers.ValidateJWT(context)
        if err != nil {
            context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
            context.Abort()
            return
        }
        context.Next()
    }
}

