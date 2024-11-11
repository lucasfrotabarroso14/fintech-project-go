package messaging

import "github.com/streadway/amqp"

var Conn *amqp.Connection

func InitRabbitMQ() error {
	var err error
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return err
	}
	return nil
}

func CloseRabbitMQ() {
	if Conn != nil {
		Conn.Close()
	}
}
