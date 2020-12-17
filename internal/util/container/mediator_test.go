package container

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	testImage         = "alpine"
	testContainerName = "mediator-test-1"
)

func TestDockerMediator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Container Suite")
}

var _ = Describe("Container", func() {
	var containerID string
	ctx := context.Background()
	var provider *DockerProvider
	var mediator *Mediator

	BeforeEach(func() {
		var err error
		provider, err = NewDockerProvider()
		Expect(err).ShouldNot(HaveOccurred())

		err = provider.PullImage(ctx, testImage)
		Expect(err).ShouldNot(HaveOccurred())

		containerID, err = provider.CreateContainer(ctx, CreateContainerArgs{
			Image: testImage,
			Name:  testContainerName,
			Cmd:   []string{"tail", "-f", "/dev/null"},
		})
		Expect(err).ShouldNot(HaveOccurred())

		err = provider.StartContainer(ctx, containerID)
		Expect(err).ShouldNot(HaveOccurred())

		mediator = NewMediator(provider.Client(), containerID)

	})

	AfterEach(func() {
		err := provider.StopContainer(ctx, containerID)
		Expect(err).ShouldNot(HaveOccurred())

		err = provider.RemoveContainer(ctx, containerID)
		Expect(err).ShouldNot(HaveOccurred())
	})

	testText := "Hello mello!"

	Describe("container executes command", func() {
		It(fmt.Sprintf("should contain %s", testText), func() {
			err := mediator.Exec(ctx, "echo", testText, ">>", "/proc/1/fd/1")
			Expect(err).ShouldNot(HaveOccurred())

			logs, err := mediator.Logs(ctx)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(logs).Should(ContainSubstring(testText))

		})
	})
})
