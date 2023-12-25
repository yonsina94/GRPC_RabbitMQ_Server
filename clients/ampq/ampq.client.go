package ampq

import (
	"context"
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Channel *amqp.Channel
}

func New(rabbitUrl string) (conn *Client, err error) {
	Conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		return nil, err
	}

	defer Conn.Close()

	ch, err := Conn.Channel()

	conn = &Client{
		Channel: ch,
	}

	return
}

func (c *Client) DeclareQueque(name string, data json.RawMessage, ctx context.Context) (err error) {
	_, err = c.Channel.QueueDeclare(name, false, true, true, true, amqp.Table{})
	return
}

func (c *Client) Publish(routingKey string, data json.RawMessage, ctx context.Context) (err error) {
	if c.Channel != nil {

		defer func() {
			c.Channel.Close()
			c.Channel = nil
			err = errors.New("Channel closed")
		}()

		err = c.Channel.PublishWithContext(ctx, "events", routingKey, false, true, amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})

		return
	} else {
		err = errors.New("Channel not opened !")
		return
	}
}
