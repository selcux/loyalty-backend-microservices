package fabric

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"os"
)

const containerPrefix = "loyalty-cc-"

type ContainerController interface {
	PreBuildConfig() error
	RunContainer(string) (string, error)
	Deploy() error
	Install() error
}

type ContainerManager struct {
}

func (c *ContainerManager) PreBuildConfig() error {
	panic("implement me")
}

func (c *ContainerManager) RunContainer(name string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	defer cli.Close()

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		return "", err
	}
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		return "", err
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   false,
	}, nil, nil, nil, containerPrefix+name)
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", err
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", err
	}

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return resp.ID, err
}

func (c *ContainerManager) Deploy() error {
	panic("implement me")
}

func (c *ContainerManager) Install() error {
	panic("implement me")
}

func NewContainerManager() *ContainerManager {
	return &ContainerManager{}
}
