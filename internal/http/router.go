package router

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "desafio-tecnico/internal/config"
    "desafio-tecnico/internal/http/handlers"
    "desafio-tecnico/internal/http/middleware"
    "desafio-tecnico/internal/service"
    "desafio-tecnico/pkg/queue"
)

func Register(r *gin.Engine, cfg *config.Config, userSvc service.UserService, pub queue.Publisher) *gin.Engine {
    // Health
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    // v1
    v1 := r.Group("/api/v1")

    authHandler := handlers.NewAuthHandler(cfg, userSvc, pub)
    userHandler := handlers.NewUserHandler(userSvc)

    // Auth
    v1.POST("/auth/signup", authHandler.Signup)
    v1.POST("/auth/login", authHandler.Login)

    // Users (protected)
    authMW := middleware.JWT(cfg, userSvc)

    users := v1.Group("/users", authMW)
    {
        users.GET("", userHandler.List)
        users.GET("/:id", userHandler.GetByID)
        users.PUT("/:id", userHandler.Update)
        users.DELETE("/:id", userHandler.Delete)
    }

    // Log startup
    logrus.Infof("routes ready: /health, /api/v1/auth/*, /api/v1/users/*")
    return r
}

// Helpers (optional) for parsing query ints
func mustAtoi(s string, def int) int {
    if s == "" {
        return def
    }
    i, err := strconv.Atoi(s)
    if err != nil {
        return def
    }
    return i
}