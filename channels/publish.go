package channels

import "github.com/streadway/amqp"

//Publisher contains the necessary information to create a channel
type Publisher struct {
	ChannelName string
	URL         string
	Channel     *amqp.Channel
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
		Channel:     ch,
		connection:  conn,
	}
	return publisher, nil

}

func getChannel(channel, chtype string, durable, autodeleted, internal, nowait bool, conn *amqp.Connection) (*amqp.Channel, error) {

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		channel,     // name
		chtype,      // type
		durable,     // durable
		autodeleted, // auto-deleted
		internal,    // internal
		nowait,      // no-wait
		nil,         // arguments
	)
	if err != nil {
		return nil, err
	}

	return ch, nil

}
