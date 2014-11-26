package channels_test

import (
	"os"

	. "github.com/thetonymaster/go-ios-notification-server/channels"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Channel", func() {

	var (
		publisher  *Publisher
		subscriber *Subscriber
		err        error
	)
	BeforeEach(func() {
		rmqtx := os.Getenv("RABBITMQ_BIGWIG_TX_URL")
		rmqrx := os.Getenv("RABBITMQ_BIGWIG_RX_URL")
		publisher, err = NewPublisher("test_channel", "fanout", rmqtx)
		subscriber, err = NewSubscriber("test_channel", "fanout", rmqrx, "", "")
	})

	Describe("Create a new Publisher", func() {
		Context("With basic parameters", func() {
			It("Should be not nil", func() {
				Expect(publisher).ToNot(BeNil())
				Expect(publisher.URL).To(Equal(os.Getenv("RABBITMQ_BIGWIG_TX_URL")))
				Expect(publisher.ChannelName).To(Equal("test_channel"))
			})
			It("Should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Send a message", func() {
		Context("In the default channel", func() {
			type Message struct {
				Msg string
			}
			msg := Message{
				Msg: "Hello",
			}

			It("Should not return an error", func() {
				err = publisher.Publish(msg, "", false, false)

				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Create a new Subscriber", func() {
		Context("With basic parameters", func() {
			It("Should be not nil", func() {
				Expect(subscriber).ToNot(BeNil())
				Expect(subscriber.URL).To(Equal(os.Getenv("RABBITMQ_BIGWIG_RX_URL")))
				Expect(subscriber.ChannelName).To(Equal("test_channel"))
			})
			It("Should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Send a message", func() {
		Context("With a hello json", func() {
			type Message struct {
				Msg string
			}
			msg := Message{
				Msg: "Hello",
			}
			var receivedmessage []byte
			It("Should not return an error sending it", func() {

				go func() {
					for d := range subscriber.DeliveryChannel {
						receivedmessage = d.Body
					}
				}()
				err = publisher.Publish(msg, "", false, false)
				Expect(err).To(BeNil())

			})
			It("Should receive a message", func() {

				Expect(string(receivedmessage)).To(Equal(`{"Msg":"Hello"}`))
			})
		})
	})

})
