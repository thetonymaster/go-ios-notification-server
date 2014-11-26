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
			})
			It("Should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

})
