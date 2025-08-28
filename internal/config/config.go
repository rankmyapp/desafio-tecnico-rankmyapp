package config

import (
    "os"
    "strconv"
    "time"
)

type Config struct {
    AppEnv         string
    AppPort        int
    DBHost         string
    DBPort         int
    DBUser         string
    DBPassword     string
    DBName         string
    DBSSLMode      string
    JWTSecret      string
    JWTExpiresIn   time.Duration
    RabbitURL      string
    EmailQueueName string
    SMTPHost       string
    SMTPPort       int
    SMTPUser       string
    SMTPPass       string
    SMTPFrom       string
}

func Load() *Config {
    return &Config{
        AppEnv:         getEnv("APP_ENV", "dev"),
        AppPort:        getEnvAsInt("APP_PORT", 8080),
        DBHost:         getEnv("DB_HOST", "localhost"),
        DBPort:         getEnvAsInt("DB_PORT", 5432),
        DBUser:         getEnv("DB_USER", "app"),
        DBPassword:     getEnv("DB_PASSWORD", "app"),
        DBName:         getEnv("DB_NAME", "appdb"),
        DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
        JWTSecret:      getEnv("JWT_SECRET", "supersecret"),
        JWTExpiresIn:   getEnvAsDuration("JWT_EXPIRES_IN", time.Hour),
        RabbitURL:      getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
        EmailQueueName: getEnv("EMAIL_QUEUE_NAME", "emails"),
        SMTPHost:       getEnv("SMTP_HOST", "localhost"),
        SMTPPort:       getEnvAsInt("SMTP_PORT", 1025),
        SMTPUser:       getEnv("SMTP_USER", ""),
        SMTPPass:       getEnv("SMTP_PASS", ""),
        SMTPFrom:       getEnv("SMTP_FROM", "no-reply@example.com"),
    }
}

func getEnv(key string, def string) string {
    if v, ok := os.LookupEnv(key); ok {
        return v
    }
    return def
}

func getEnvAsInt(key string, def int) int {
    if v, ok := os.LookupEnv(key); ok {
        if i, err := strconv.Atoi(v); err == nil {
            return i
        }
    }
    return def
}

func getEnvAsDuration(key string, def time.Duration) time.Duration {
    if v, ok := os.LookupEnv(key); ok {
        if d, err := time.ParseDuration(v); err == nil {
            return d
        }
    }
    return def
}