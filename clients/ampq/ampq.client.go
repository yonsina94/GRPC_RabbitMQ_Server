package ampq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct{
	Channel *amqp.Channel
}

func New(rabbitUrl string) (conn *Client, err error) {
	Conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		return nil, err
	}

	ch, err := Conn.Channel()

	conn = &Client{
		Channel: ch,
	}

	return
}

func (cli *Client) Publish(routingKey string, data json.RawMessage) (err error){
	return cli.Channel.Publish("events",routingKey,false,true, amqp.Publishing{
		ContentType: "application/json",
		Body: data,
		DeliveryMode: amqp.Persistent,
	})
}