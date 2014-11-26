package channels_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChannels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Channels Suite")
}
