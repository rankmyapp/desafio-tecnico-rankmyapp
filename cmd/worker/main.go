package main

import (
    "encoding/json"
    "log"
    "time"

    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"
    "desafio-tecnico/internal/config"
    "desafio-tecnico/pkg/mailer"
    "desafio-tecnico/pkg/queue"
)

type EmailPayload struct {
    To      string `json:"to"`
    Subject string `json:"subject"`
    Body    string `json:"body"`
}

func main() {
    _ = godotenv.Load()

    cfg := config.Load()
    logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
    logrus.SetLevel(logrus.InfoLevel)

    consumer, err := queue.NewConsumer(cfg)
    if err != nil {
        log.Fatalf("consumer init error: %v", err)
    }
    defer consumer.Close()

    m := mailer.New(cfg)

    logrus.Infof("Worker listening queue=%s ...", cfg.EmailQueueName)
    msgs, err := consumer.Consume(cfg.EmailQueueName)
    if err != nil {
        log.Fatalf("consume error: %v", err)
    }

    for d := range msgs {
        var payload EmailPayload
        if err := json.Unmarshal(d.Body, &payload); err != nil {
            logrus.Errorf("invalid message: %v", err)
            d.Nack(false, false)
            continue
        }

        if err := m.Send(payload.To, payload.Subject, payload.Body); err != nil {
            logrus.Errorf("send email error: %v", err)
            d.Nack(false, true) // requeue
            continue
        }

        d.Ack(false)
        time.Sleep(50 * time.Millisecond) // leve respiro
    }
}