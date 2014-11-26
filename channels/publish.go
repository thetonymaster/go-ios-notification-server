package channels

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

//Publisher contains the necessary information to send messages a channel
type Publisher struct {
	ChannelName string
	URL         string
	channel     *amqp.Channel
	connection  *amqp.Connection
}

//NewPublisher returns a new Publisher struct
func NewPublisher(channel, chtype, url string) (*Publisher, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := getChannel(channel, chtype, true, false, false, false, conn)
	if err != nil {
		return nil, err
	}

	publisher := &Publisher{
		ChannelName: channel,
		URL:         url,
		channel:     ch,
		connection:  conn,
	}
	return publisher, nil

}

//Publish sends a message to the queue
func (publisher *Publisher) Publish(v interface{}, routingkey string, mandatory, immediate bool) error {

	body, err := json.Marshal(v)

	err = publisher.channel.Publish(
		publisher.ChannelName,
		routingkey,
		mandatory,
		immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}
