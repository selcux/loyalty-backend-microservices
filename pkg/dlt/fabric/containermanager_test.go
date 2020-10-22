package fabric

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/onsi/gomega"
	"io"
	"strings"
	"testing"
)

func TestContainerManager_RunContainer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	cManager := NewContainerManager()
	containerID, err := cManager.RunContainer("test1")
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, containerID, options)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	defer out.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, out)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	logs := buf.String()
	g.Expect(logs).Should(gomega.ContainSubstring("hello world"))

	err = cli.ContainerStop(ctx, containerID, nil)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())
	err = cli.Close()
	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	ctx = context.Background()
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	g.Expect(err).ShouldNot(gomega.HaveOccurred())
	defer cli.Close()

	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{RemoveVolumes: true})
	g.Expect(err).ShouldNot(gomega.HaveOccurred())
}
