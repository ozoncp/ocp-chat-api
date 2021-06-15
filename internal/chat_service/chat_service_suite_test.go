package chat_service_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChatApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ChatApi Suite")
}
