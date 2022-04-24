package src

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type Connection struct {
	*amqp.Connection
	delay int
}

func Dial(URL string, delay int) (*Connection, error) {

	conn, err := amqp.Dial(URL)
	if err != nil {
		return nil, err
	}

	connection := &Connection{
		Connection: conn,
		delay:      delay,
	}

	go func() {
		for {
			reason, ok := <-connection.Connection.NotifyClose(make(chan *amqp.Error))
			// exit this goroutine if closed by developer
			if !ok {
				fmt.Println("connection closed")
				break
			}
			fmt.Println("connection closed, reason: %v", reason)

			// reconnect if not closed by developer
			for {
				// wait 1s for reconnect
				time.Sleep(time.Duration(delay) * time.Second)

				conn, err := amqp.Dial(URL)
				if err == nil {
					connection.Connection = conn
					fmt.Println("reconnect success")
					break
				}

				fmt.Println("reconnect failed, err: %v", err)
			}
		}
	}()

	return connection, nil
}
