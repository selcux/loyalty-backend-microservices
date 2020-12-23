// +build container

package fabric

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/onsi/gomega"
	"io"
	"log"
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

	err = executeEcho(ctx, containerID, "hello world")
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

	err = removeContainer(ctx, containerID)
	g.Expect(err).ShouldNot(gomega.HaveOccurred())
}

func removeContainer(ctx context.Context, containerID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	defer cli.Close()

	return cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{RemoveVolumes: true})
}

func executeEcho(ctx context.Context, containerID, message string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("client.NewClientWithOpts")
		return err
	}
	defer cli.Close()

	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		log.Println("cli.ContainerStart")
		return err
	}

	statusCh, errCh := cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Println("<-errCh")
			return err
		}
	case <-statusCh:
	}

	c := types.ExecConfig{AttachStdout: true, AttachStderr: true,
		Cmd: []string{"echo", message}}

	execID, err := cli.ContainerExecCreate(ctx, containerID, c)
	if err != nil {
		log.Println("cli.ContainerExecCreate")
		return err
	}
	log.Println("execID", execID)

	_, err = cli.ContainerExecAttach(ctx, execID.ID, types.ExecStartCheck{})
	if err != nil {
		log.Println("cli.ContainerExecAttach")
		return err
	}

	err = cli.ContainerExecStart(ctx, execID.ID, types.ExecStartCheck{})
	return err
}
