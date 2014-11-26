package channels_test

import (
	"os"

	. "github.com/thetonymaster/go-ios-notification-server/channels"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Channel", func() {

	var (
		publisher *Publisher
		err       error
	)
	BeforeEach(func() {
		rmqtx := os.Getenv("RMQ_TX")
		publisher, err = NewPublisher("test_channel", "fanout", rmqtx)
	})

	Describe("Create a new Publisher", func() {
		Context("With basic parameters", func() {
			It("Should be not nil", func() {
				Expect(publisher).ToNot(BeNil())
				Expect(publisher.URL).To(Equal(os.Getenv("RMQ_TX")))
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

})
