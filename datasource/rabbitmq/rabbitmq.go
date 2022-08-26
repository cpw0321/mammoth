package rabbitmq

import (
	"context"
	"fmt"

	configs "github.com/cpw0321/mammoth/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

var MQConn *amqp.Connection

type MQMessage struct {
	BifUserBid string `json:"bif_user_bid"`
	AppId      string `json:"app_id"`
	Type       string `json:"type"`
	Body       string `json:"body"`
}

// InitRabbitMQ ...
func InitRabbitMQ() {
	// 新建一个连接
	conn, err := amqp.Dial(configs.Conf.Rabbitmq.Addr)
	if err != nil {
		panic(fmt.Errorf("init rabbitmq err: %v \n", err))
	}
	MQConn = conn
	return
}

func QueueDeclare(conn *amqp.Connection, queueName string) (*amqp.Channel, *amqp.Queue, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	// channel.ExchangeDeclare()
	queue, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, nil, err
	}
	return channel, &queue, nil
}

func Publish(channel *amqp.Channel, queueName string, message []byte) error {
	return channel.PublishWithContext(
		context.Background(),
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}

func Consume(channel *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	return channel.Consume(queueName, "", false, false, false, false, nil)
}
