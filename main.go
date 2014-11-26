package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/thetonymaster/go-ios-notification-server/channels"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/send-notification", sendNotification)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func sendNotification(res http.ResponseWriter, req *http.Request) {
	rmqtx := os.Getenv("RABBITMQ_BIGWIG_TX_URL")
	publisher, err := channels.NewPublisher("test_channel", "fanout", rmqtx)
	if err != nil {
		panic(err)
	}

	type Message struct {
		Msg string
	}

	msg := &Message{
		Msg: "Hello",
	}

	err = publisher.Publish(msg, "", false, false)
	if err != nil {
		panic(err)
	}

	publisher.Close <- true
}
