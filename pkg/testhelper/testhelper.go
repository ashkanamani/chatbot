package testhelper

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log/slog"
	"os"
)

type RetryFunc func(resource *dockertest.Resource) error

func IsIntegrationTest() bool {
	return os.Getenv("INTEGRATION_TEST") == "true"
}

func StartDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		slog.Error("could not construct docker pool")
		os.Exit(1)
	}
	err = pool.Client.Ping()
	if err != nil {
		slog.Error("could not ping docker pool")
		os.Exit(1)
	}
	return pool
}

func StartDockerInstance(pool *dockertest.Pool, image, tag string, retryFunc RetryFunc, env ...string) *dockertest.Resource {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: image,
		Tag:        tag,
		Env:        env,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		slog.Error("could not start docker instance", "image", image, "tag", tag, "err", err.Error())
		os.Exit(1)
	}
	if err := resource.Expire(30); err != nil {
		slog.Error("could not set resource expiration", "image", image, "tag", tag, "err", err.Error())
		os.Exit(1)
	}
	if err := pool.Retry(func() error {
		return retryFunc(resource)
	}); err != nil {
		slog.Error("could not connect to the resource", "image", image, "tag", tag, "err", err.Error())
		os.Exit(1)
	}
	slog.Info("connected to docker instance", "image", image, "tag", tag)
	return resource
}
