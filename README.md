# RAPLIN - Simple RabbitMQ Plugin Based on AMQP.

RAPLIN (RABBITQ PLUGIN) - Simple RabbitMQ plugin based on AMQP.

The project has been setting to be able to handle subscribe or publish message with Gorotine.

The library functions:
- InitRaplin -> to init the connection
- RabbitConnection -> to connect to rabbitmq server
- GetConnection -> to get the rabbitmq connection
- GetChannel -> to get the channel
- DeclareExchange -> to decalare exchange multiple or single
- DeclareQueue -> to declare new queue
- BindQueue -> to bind the queue with the exchange
- DeclareMultiQueueAndBind -> to declare queue and bind with the exchange(the relation between exchange and queue only 1:1)
- ReadMessage -> to read the mesage in queue based on Queue Name, Exchange, or Routing Key
- PublishMessage -> to publish the message to queue with custom setting

# Installation

```
$ go get github.com/ianyulistios/raplin
```
# Why RAPLIN?

With Raplin, you can configure rabbit without complex settings:

- **Make Instance Init**
```
initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
```

- **Connect To RabbitMQ Server**
```
initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
conn, err := initRaplin.RabbitConnection()
```

- **Declare Exchange**
```
var (
	exchangeSetting     raplin.RabbitDataDeclare
	exchangeSettingBulk []raplin.RabbitDataDeclare
    )

    initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
    conn, err := initRaplin.RabbitConnection()

    if err != nil {
        fmt.Println(err.Error())
    }else {
        exchangeSetting := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil}

        exchangeSettingBulk = append(exchangeSettingBulk, exchangeSetting)
        _ := conn.ExchangeDeclaration(exchangeSettingBulk)
    }
```
- **Declare Queue**
```
initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
conn, err := initRaplin.RabbitConnection()

if err != nil {
   fmt.Println(err.Error())
} else {
    queueSetting := raplin.RabbitDataDeclare{QueueName: "TESTING QUEUE", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil}
    _ := conn.DeclareQueue(queueSetting)
}
```
**Bind Queue**
```
var (
	exchangeSettingBulk []raplin.RabbitDataDeclare
    )
    initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
    conn, err := initRaplin.RabbitConnection()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        rabbitSetting := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq-bind", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil, QueueName: "TESTING QUEUE BIND", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil, NoWaitBind: false, KeyBind: ""}
        exchangeSettingBulk = append(exchangeSettingBulk, rabbitSetting)

        _ := conn.ExchangeDeclaration(exchangeSettingBulk)
        _ = conn.DeclareQueue(rabbitSetting)
        _ = conn.BindQueue(rabbitSetting)
    }
```

**Complex 1:1 Queue, Exchange, Bind**
```
    var (
        exchangeSettingBulk []raplin.RabbitDataDeclare
    )
    initRaplin := raplin.InitRaplin("127.0.0.1", "5672", "guest", "guest", 3)
    conn, err := initRaplin.RabbitConnection()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        rabbitSetting := raplin.RabbitDataDeclare{ExchangeName: "test-rabbitmq-complex", ExchangeType: "direct", DurableExchange: false, AutoDeleteExchange: true, InternalExchange: false, NoWaitExchange: false, ArgsExchange: nil, QueueName: "TESTING QUEUE COMPLEX", DurableQueue: false, AutoDeleteQueue: true, ExclusiveQueue: false, NoWaitQueue: false, ArgsQueue: nil, NoWaitBind: false, KeyBind: ""}
        exchangeSettingBulk = append(exchangeSettingBulk, rabbitSetting)
        
        conn.ExchangeDeclaration(exchangeSettingBulk)
        conn.DeclareMultiQueueAndBind(exchangeSettingBulk)
    }
```
For example please look into test case

# LICENSE
RAPLIN is MIT License.
