package raplin

import (
	"encoding/json"
	"fmt"

	"github.com/ianyulistios/raplin/src"

	"github.com/streadway/amqp"
)

type RaplinAgent struct {
	rabbitURL               string
	rabbitPort              string
	rabbitUsername          string
	rabbitPassword          string
	rabbitMQConnectionDelay int
	rabbitMQConnection      *src.Connection
	rabbitMQChannel         *src.Channel
}

type RabbitDataDeclare struct {
	ExchangeName       string
	ExchangeType       string
	DurableExchange    bool
	AutoDeleteExchange bool
	InternalExchange   bool
	NoWaitExchange     bool
	ArgsExchange       amqp.Table
	QueueName          string
	DurableQueue       bool
	AutoDeleteQueue    bool
	NoWaitQueue        bool
	ExclusiveQueue     bool
	ArgsQueue          amqp.Table
	NoWaitBind         bool
	ArgsBind           amqp.Table
	KeyBind            string
}

func InitRaplin(rabbitURL, rabbitPort, rabbitUsername, rabbitPassword string, rabbitMQConnectionDelay int) *RaplinAgent {
	return &RaplinAgent{
		rabbitURL:               rabbitURL,
		rabbitPort:              rabbitPort,
		rabbitUsername:          rabbitUsername,
		rabbitPassword:          rabbitPassword,
		rabbitMQConnectionDelay: rabbitMQConnectionDelay,
	}
}

func (r *RaplinAgent) RabbitConnection() (*RaplinAgent, error) {
	var (
		ch  *src.Channel
		err error
	)
	urlRabbit := "amqp://" + r.rabbitUsername + ":" + r.rabbitPassword + "@" + r.rabbitURL + ":" + r.rabbitPort + "/"
	conn, err := src.Dial(urlRabbit, r.rabbitMQConnectionDelay)
	r.rabbitMQConnection = conn
	ch, err = conn.Channel()
	r.rabbitMQChannel = ch
	return r, err
}

func (r *RaplinAgent) GetConnection() *src.Connection {
	conn := r.rabbitMQConnection
	return conn
}

func (r *RaplinAgent) GetChannel() (*src.Channel, error) {
	ch, err := r.rabbitMQConnection.Channel()
	return ch, err
}

func (r *RaplinAgent) ExchangeDeclaration(rabbitData []RabbitDataDeclare) error {
	var (
		err error
	)
	if err == nil {
		for _, data := range rabbitData {
			err = r.rabbitMQChannel.ExchangeDeclare(
				data.ExchangeName,
				data.ExchangeType,
				data.DurableExchange,
				data.AutoDeleteExchange,
				data.InternalExchange,
				data.NoWaitExchange,
				data.ArgsExchange,
			)
		}
	}
	return err
}

func (r *RaplinAgent) BindQueue(rabbitData RabbitDataDeclare) error {
	var (
		err error
	)
	err = r.rabbitMQChannel.QueueBind(
		rabbitData.QueueName,
		rabbitData.KeyBind,
		rabbitData.ExchangeName,
		rabbitData.NoWaitBind,
		rabbitData.ArgsBind,
	)
	return err
}

func (r *RaplinAgent) DeclareQueue(rabbitData RabbitDataDeclare) error {
	_, err := r.rabbitMQChannel.QueueDeclare(
		rabbitData.QueueName,
		rabbitData.DurableQueue,
		rabbitData.AutoDeleteQueue,
		rabbitData.ExclusiveQueue,
		rabbitData.NoWaitQueue,
		rabbitData.ArgsQueue,
	)
	return err
}

func (r *RaplinAgent) DeclareMultiQueueAndBind(rabbitData []RabbitDataDeclare) error {
	var (
		errs error
	)
	for _, data := range rabbitData {
		errs = r.DeclareQueue(data)
		if errs != nil {
			fmt.Println(errs)
		} else {
			errs = r.BindQueue(data)
		}
	}
	return errs
}

func (r *RaplinAgent) ReadMessage(queueName, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	message, errs := r.rabbitMQChannel.Consume(
		queueName,
		consumer,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)
	return message, errs
}

func (r *RaplinAgent) PublishMessage(exchange, routingKey, contentType string, data interface{}, mandatory, immediate bool, deliveryMode int) error {
	var (
		err error
	)
	b, errs := json.Marshal(data)
	if errs != nil {
		err = errs
	} else {
		errrs := r.rabbitMQChannel.Publish(
			exchange,   // exchange
			routingKey, // routing key
			mandatory,  // mandatory
			immediate,  // immediate
			amqp.Publishing{
				DeliveryMode: uint8(deliveryMode), // Persistent Message is 2 and normal message is 1
				ContentType:  contentType,
				Body:         b},
		)
		err = errrs
	}
	return err
}
