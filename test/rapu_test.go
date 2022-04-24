package test

import (
	"testing"

	"github.com/ianyulistios/raplin"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		_, errs := conn.GetChannel()
		assert.Nil(t, errs)
	}
}

func TestCreateSingleExchange(t *testing.T) {
	var (
		exchangeSetting     raplin.RabbitDataDeclare
		exchangeSettingBulk []raplin.RabbitDataDeclare
	)
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		exchangeSetting.ExchangeName = "test-rabbitmq-single"
		exchangeSetting.DurableQueue = false
		exchangeSetting.ArgsExchange = nil
		exchangeSetting.AutoDeleteExchange = true
		exchangeSetting.ExchangeType = "direct"
		exchangeSetting.NoWaitExchange = false
		exchangeSetting.InternalExchange = false

		exchangeSettingBulk = append(exchangeSettingBulk, exchangeSetting)
		errr := conn.ExchangeDeclaration(exchangeSettingBulk)
		assert.Nil(t, errr)
	}
}

func TestGetChannel(t *testing.T) {
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		_, errr := conn.GetChannel()
		assert.Nil(t, errr)
	}
}

func TestGetConnection(t *testing.T) {
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		connection := conn.GetConnection()
		assert.NotNil(t, connection)
	}
}

func TestCreateMultiExchange(t *testing.T) {
	var (
		exchangeSettingBulk []raplin.RabbitDataDeclare
	)
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		exchangeSettingOne := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil}
		exchangeSettingTwo := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq2", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil}
		exchangeSettingBulk = append(exchangeSettingBulk, exchangeSettingOne)
		exchangeSettingBulk = append(exchangeSettingBulk, exchangeSettingTwo)
		errr := conn.ExchangeDeclaration(exchangeSettingBulk)
		assert.Nil(t, errr)
	}
}

func TestCreateQueue(t *testing.T) {
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		queueSetting := raplin.RabbitDataDeclare{QueueName: "TESTING QUEUE", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil}
		errr := conn.DeclareQueue(queueSetting)
		assert.Nil(t, errr)
	}
}

func TestCreateQueueWithArgument(t *testing.T) {
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		argsSetting := make(amqp.Table)
		argsSetting["x-message-ttl"] = 10000
		queueSetting := raplin.RabbitDataDeclare{QueueName: "TESTING QUEUE WITH ARGUMENT", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: argsSetting}
		errr := conn.DeclareQueue(queueSetting)
		assert.Nil(t, errr)
	}
}

func TestBindQueue(t *testing.T) {
	var (
		exchangeSettingBulk []raplin.RabbitDataDeclare
	)
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		rabbitSetting := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq-bind", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil, QueueName: "TESTING QUEUE BIND", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil, NoWaitBind: false, KeyBind: ""}
		exchangeSettingBulk = append(exchangeSettingBulk, rabbitSetting)
		errr := conn.ExchangeDeclaration(exchangeSettingBulk)
		assert.Nil(t, errr)

		errr = conn.DeclareQueue(rabbitSetting)
		assert.Nil(t, errr)

		errr = conn.BindQueue(rabbitSetting)
		assert.Nil(t, errr)
	}
}

func TestCreateQueueComplex(t *testing.T) {
	var (
		exchangeSettingBulk []raplin.RabbitDataDeclare
	)
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		rabbitSetting := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq-complex", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil, QueueName: "TESTING QUEUE COMPLEX", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil, NoWaitBind: false, KeyBind: ""}
		exchangeSettingBulk = append(exchangeSettingBulk, rabbitSetting)
		errrs := conn.ExchangeDeclaration(exchangeSettingBulk)
		assert.Nil(t, errrs)

		errrs = conn.DeclareMultiQueueAndBind(exchangeSettingBulk)
		assert.Nil(t, errrs)
	}
}

func TestCreateMultiQueueComplex(t *testing.T) {
	var (
		exchangeSettingBulk []raplin.RabbitDataDeclare
	)
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		rabbitSettingOne := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq-multi-complex-one", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil, QueueName: "TESTING QUEUE MULTI COMPLEX ONE", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil, NoWaitBind: false, KeyBind: ""}
		rabbitSettingTwo := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq-multi-complex-two", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil, QueueName: "TESTING QUEUE MULTI COMPLEX TWO", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil, NoWaitBind: false, KeyBind: ""}
		exchangeSettingBulk = append(exchangeSettingBulk, rabbitSettingOne)
		exchangeSettingBulk = append(exchangeSettingBulk, rabbitSettingTwo)
		errrs := conn.ExchangeDeclaration(exchangeSettingBulk)
		assert.Nil(t, errrs)

		errrs = conn.DeclareMultiQueueAndBind(exchangeSettingBulk)
		assert.Nil(t, errrs)
	}
}

func TestPublishMessage(t *testing.T) {
	var (
		exchangeSettingBulk []raplin.RabbitDataDeclare
	)
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		rabbitSettingOne := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq-message-publish", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil, QueueName: "TESTING QUEUE MESSAGE PUBLISH", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil, NoWaitBind: false, KeyBind: ""}
		exchangeSettingBulk = append(exchangeSettingBulk, rabbitSettingOne)
		errrs := conn.ExchangeDeclaration(exchangeSettingBulk)
		assert.Nil(t, errrs)

		errrs = conn.DeclareMultiQueueAndBind(exchangeSettingBulk)
		assert.Nil(t, errrs)
		dummyData := `{"message": "OK"}`

		errrs = conn.PublishMessage(rabbitSettingOne.ExchangeName, "", "text/plain", dummyData, false, false, 2)
		assert.Nil(t, errrs)
	}
}

func TestPublishAndReadMessage(t *testing.T) {
	var (
		exchangeSettingBulk []raplin.RabbitDataDeclare
		messages            <-chan amqp.Delivery
	)
	initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
	conn, err := initRaplin.RabbitConnection()
	if err != nil {
		assert.NotNil(t, err)
	} else {
		rabbitSettingOne := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq-message-publish-read", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil, QueueName: "TESTING QUEUE MESSAGE PUBLISH AND READ", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil, NoWaitBind: false, KeyBind: ""}
		exchangeSettingBulk = append(exchangeSettingBulk, rabbitSettingOne)
		errrs := conn.ExchangeDeclaration(exchangeSettingBulk)
		assert.Nil(t, errrs)

		errrs = conn.DeclareMultiQueueAndBind(exchangeSettingBulk)
		assert.Nil(t, errrs)
		dummyData := `{"message": "OK"}`

		errrs = conn.PublishMessage(rabbitSettingOne.ExchangeName, "", "text/plain", dummyData, false, false, 2)
		assert.Nil(t, errrs)

		messages, errrs = conn.ReadMessage(rabbitSettingOne.QueueName, "", false, false, false, false, nil)
		assert.Nil(t, errrs)

		for msg := range messages {
			assert.NotEqual(t, string(msg.Body), "")
			msg.Ack(true)
			break
		}
	}
}
