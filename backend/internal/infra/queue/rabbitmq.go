package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
	"github.com/otaviomart1ns/backend-challenge/internal/domain/repositories"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQPublisher struct {
	channel *amqp.Channel
	queue   string
}

func NewRabbitMQPublisher(conn *amqp.Connection, queueName string) (repositories.QueuePublisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &rabbitMQPublisher{
		channel: ch,
		queue:   queueName,
	}, nil
}

func (r *rabbitMQPublisher) PublishPurchaseValidation(ctx context.Context, sale *entities.Sale) error {
	body, err := json.Marshal(sale)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"",      // exchange
		r.queue, // routing key (queue name)
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Erro ao publicar na fila: %v", err)
		return err
	}

	log.Printf("Publicado na fila %s: %s", r.queue, body)
	return nil
}
