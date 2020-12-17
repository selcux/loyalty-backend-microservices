package container

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerProvider struct {
	client *client.Client
}

func (d *DockerProvider) RemoveContainer(ctx context.Context, containerID string) error {
	cli := d.client

	return cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{RemoveVolumes: true})
}

func (d *DockerProvider) StopContainer(ctx context.Context, containerID string) error {
	cli := d.client

	return cli.ContainerStop(ctx, containerID, nil)
}

func (d *DockerProvider) StartContainer(ctx context.Context, containerID string) error {
	cli := d.client

	return cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
}

func (d *DockerProvider) Client() *client.Client {
	return d.client
}

func (d *DockerProvider) PullImage(ctx context.Context, image string) error {
	cli := d.client
	out, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(os.Stdout, out)

	return err
}

func (d *DockerProvider) CreateContainer(ctx context.Context, args CreateContainerArgs) (string, error) {
	cli := d.client
	config := &container.Config{
		Image: args.Image,
	}

	if len(args.Cmd) > 0 {
		config.Cmd = args.Cmd
	}

	cnt, err := cli.ContainerCreate(ctx, config, nil, nil, nil, args.Name)
	if err != nil {
		return "", err
	}

	return cnt.ID, nil
}

func (d *DockerProvider) Close() error {
	return d.client.Close()
}

func NewDockerProvider() (*DockerProvider, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &DockerProvider{cli}, nil
}
