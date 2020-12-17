package container

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"log"
	"strings"
	"time"
)

type CreateContainerArgs struct {
	Image, Name string
	Cmd         []string
}

type Provider interface {
	PullImage(ctx context.Context, image string) error
	CreateContainer(ctx context.Context, args CreateContainerArgs) (string, error)
	StartContainer(ctx context.Context, containerID string) error
	StopContainer(ctx context.Context, containerID string) error
	RemoveContainer(ctx context.Context, containerID string) error
	Close() error
	Client() *client.Client
}

type Commander interface {
	Exec(ctx context.Context, cmd string, args ...string) error
	Logs(ctx context.Context) (string, error)
}

type Mediator struct {
	client      *client.Client
	containerID string
}

func (m *Mediator) Logs(ctx context.Context) (string, error) {
	cli := m.client

	logReader, err := cli.ContainerLogs(ctx, m.containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return "", err
	}
	defer logReader.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, logReader)

	return buf.String(), err
}

func (m *Mediator) Exec(ctx context.Context, cmd string, args ...string) error {
	cli := m.client
	actualCmd := append([]string{cmd}, args...)
	fullCmdArgs := strings.Join(actualCmd, " ")
	log.Println("sh", "-c", fullCmdArgs)

	exec, err := cli.ContainerExecCreate(ctx, m.containerID, types.ExecConfig{
		Cmd: []string{"sh", "-c", fullCmdArgs},
	})
	if err != nil {
		return err
	}

	_, err = cli.ContainerExecAttach(ctx, exec.ID, types.ExecStartCheck{})

	err = cli.ContainerExecStart(ctx, exec.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}

	var exitCode int
	for {
		inspect, err := cli.ContainerExecInspect(ctx, exec.ID)
		if err != nil {
			return err
		}

		if !inspect.Running {
			exitCode = inspect.ExitCode
			break
		}

		time.Sleep(100 * time.Millisecond)
	}

	if exitCode != 0 {
		return fmt.Errorf("exit code: %d", exitCode)
	}

	return nil
}

func NewMediator(client *client.Client, containerID string) *Mediator {
	return &Mediator{client: client, containerID: containerID}
}
