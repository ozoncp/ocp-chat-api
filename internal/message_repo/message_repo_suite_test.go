package message_repo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMessageRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MessageRepo Suite")
}
