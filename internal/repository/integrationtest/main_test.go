package integrationtest

import (
	"fmt"
	"github.com/ashkanamani/chatbot/internal/repository/redis"
	"github.com/ashkanamani/chatbot/pkg/testhelper"
	"github.com/ory/dockertest/v3"
	"os"
	"testing"
)

var redisPort string
var postgresPort string

func TestMain(m *testing.M) {
	if !testhelper.IsIntegrationTest() {
		return
	}
	pool := testhelper.StartDockerPool()

	// Set up the redis container for tests
	redisResource := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest",
		func(res *dockertest.Resource) error {
			port := res.GetPort("6379/tcp")
			_, err := redis.NewRedisClient(fmt.Sprintf("%s:%s", "127.0.0.1", port))
			return err
		})
	redisPort = redisResource.GetPort("6379/tcp")
	fmt.Println("Redis port:", redisPort)

	//// Set up the postgresql container for tests
	//postgresResource := testhelper.StartDockerInstance(pool, "postgres", "latest",
	//	func(res *dockertest.Resource) error {
	//		port := res.GetPort("5432/tcp")
	//		_, err := postgres.NewPostgresConnection(fmt.Sprintf("%s:%s", "127.0.0.1", port))
	//		return err
	//	})
	//postgresPort = postgresResource.GetPort("5432/tcp")
	//fmt.Println("Postgres port:", postgresPort)
	//fmt.Println("DB URL:", fmt.Sprintf("postgres://postgres:postgres@127.0.0.1:%s/postgres?sslmode=disable", postgresPort))
	//err := migration.RunMigrations(
	//	"../../../internal/migration/sql",
	//	fmt.Sprintf("postgres://postgres:postgres@127.0.0.1:%s/postgres?sslmode=disable", postgresPort),
	//)
	//if err != nil {
	//	slog.Error("running migrations in tests failed", "err", err)
	//	os.Exit(1)
	//}

	defer func() {
		_ = redisResource.Close()
		//_ = postgresResource.Close()
	}()

	// now run tests
	exitCode := m.Run()
	os.Exit(exitCode)
}
