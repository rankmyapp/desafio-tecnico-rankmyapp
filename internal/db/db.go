package db

import (
    "fmt"
    "time"

    "github.com/sirupsen/logrus"
    "desafio-tecnico/internal/config"
    "desafio-tecnico/internal/domain"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectPostgresWithRetry(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

    var db *gorm.DB
    var err error
    for i := 0; i < 20; i++ {
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            sqlDB, _ := db.DB()
            sqlDB.SetMaxOpenConns(10)
            sqlDB.SetMaxIdleConns(5)
            sqlDB.SetConnMaxLifetime(time.Hour)
            return db, nil
        }
        logrus.Warnf("db connect attempt %d failed: %v", i+1, err)
        time.Sleep(2 * time.Second)
    }
    return db, err
}

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(&domain.User{})
}