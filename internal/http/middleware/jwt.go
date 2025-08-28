package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "desafio-tecnico/internal/config"
    "desafio-tecnico/internal/service"
)

const CtxUserID = "userID"

func JWT(cfg *config.Config, userSvc service.UserService) gin.HandlerFunc {
    jwtSvc := service.NewJWTService(cfg)

    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
            return
        }
        tokenStr := strings.TrimPrefix(auth, "Bearer ")
        uid, err := jwtSvc.ParseSubject(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }
        // Optional: validar se user ainda existe
        if _, err := userSvc.GetByID(uid); err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
            return
        }
        c.Set(CtxUserID, uid)
        c.Next()
    }
}