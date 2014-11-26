package channels

import "github.com/streadway/amqp"

//Subscriber contains the necessary information to subscribe to a channel
type Subscriber struct {
	ChannelName     string
	URL             string
	DeliveryChannel <-chan amqp.Delivery
}

//NewSubscriber returns a new publisher struct and it's possible errors
func NewSubscriber(channel, chtype, url, queuename, consumer string) (*Subscriber, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := getChannel(channel, chtype, true, false, false, false, conn)
	if err != nil {
		return nil, err
	}

	q, err := getQueue(queuename, true, false, false, false, ch)
	if err != nil {
		return nil, err
	}

	err = bindQueue(ch, q, channel, "", false)
	if err != nil {
		return nil, err
	}

	msgs, err := getDeliveryChannel(ch, q, consumer, true, false, false, false)

	subscriber := &Subscriber{
		ChannelName:     channel,
		URL:             url,
		DeliveryChannel: msgs,
	}

	return subscriber, nil
}
