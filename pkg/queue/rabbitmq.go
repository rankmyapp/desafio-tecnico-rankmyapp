package queue

import (
	"time"
	"context"

	"desafio-tecnico/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Publisher interface {
	Publish(queueName string, body []byte) error
	Close()
}

type Consumer interface {
	Consume(queueName string) (<-chan amqp.Delivery, error)
	Close()
}

type rmq struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	cfg  *config.Config
}

func connectWithRetry(url string) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error
	for i := 0; i < 20; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			return conn, nil
		}
		logrus.Warnf("rabbitmq connect attempt %d failed: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	return conn, err
}

func NewPublisher(cfg *config.Config) (Publisher, error) {
	conn, err := connectWithRetry(cfg.RabbitURL)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	// declara fila para garantir existência
	_, err = ch.QueueDeclare(cfg.EmailQueueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &rmq{conn: conn, ch: ch, cfg: cfg}, nil
}

func (r *rmq) Publish(queueName string, body []byte) error {
	return r.ch.PublishWithContext(
		context.TODO(),
		"", // default exchange
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
}

func (r *rmq) Close() {
	if r.ch != nil {
		_ = r.ch.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
}

func NewConsumer(cfg *config.Config) (Consumer, error) {
	conn, err := connectWithRetry(cfg.RabbitURL)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare(cfg.EmailQueueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	// QoS
	if err := ch.Qos(10, 0, false); err != nil {
		return nil, err
	}
	return &rmq{conn: conn, ch: ch, cfg: cfg}, nil
}

func (r *rmq) Consume(queueName string) (<-chan amqp.Delivery, error) {
	return r.ch.Consume(queueName, "", false, false, false, false, nil)
}
