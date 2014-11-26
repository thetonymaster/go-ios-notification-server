package channels

import "github.com/streadway/amqp"

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

func getQueue(name string, durable, autodelete, exclusive, nowait bool, ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		name,       // name
		durable,    // durable
		autodelete, // delete when usused
		exclusive,  // exclusive
		nowait,     // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, err
	}
	return &q, err
}

func bindQueue(ch *amqp.Channel, q *amqp.Queue, channelname, routingkey string, nowait bool) error {
	err := ch.QueueBind(
		q.Name,      // queue name
		routingkey,  // routing key
		channelname, // exchange
		nowait,
		nil)
	if err != nil {
		return err
	}
	return nil
}

func getDeliveryChannel(ch *amqp.Channel, q *amqp.Queue, consumer string, autoack, exclusive, nolocal, nowait bool) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		q.Name,    // queue
		"",        // consumer
		autoack,   // auto-ack
		exclusive, // exclusive
		nolocal,   // no-local
		nowait,    // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
