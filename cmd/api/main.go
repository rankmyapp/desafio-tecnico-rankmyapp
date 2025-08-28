package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"desafio-tecnico/internal/config"
	"desafio-tecnico/internal/db"
	router "desafio-tecnico/internal/http"
	"desafio-tecnico/internal/repository"
	"desafio-tecnico/internal/service"
	"desafio-tecnico/pkg/queue"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
    _ = godotenv.Load() // carrega .env se existir

    cfg := config.Load()

    logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
    logrus.SetLevel(logrus.InfoLevel)

    // DB
    gormDB, err := db.ConnectPostgresWithRetry(cfg)
    if err != nil {
        log.Fatalf("database connection failed: %v", err)
    }
    if err := db.AutoMigrate(gormDB); err != nil {
        log.Fatalf("auto migrate failed: %v", err)
    }

    // RabbitMQ Publisher
    publisher, err := queue.NewPublisher(cfg)
    if err != nil {
        log.Fatalf("rabbitmq publisher failed: %v", err)
    }
    defer publisher.Close()

    // Layers
    userRepo := repository.NewUserRepository(gormDB)
    jwtSvc := service.NewJWTService(cfg)
    userSvc := service.NewUserService(userRepo, jwtSvc)

    // Router
    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())

    app := router.Register(r, cfg, userSvc, publisher)

    srv := &http.Server{
        Addr:           fmt.Sprintf(":%d", cfg.AppPort),
        Handler:        app,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   15 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    logrus.Infof("API listening on :%d", cfg.AppPort)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        logrus.Fatalf("server error: %v", err)
    }
}