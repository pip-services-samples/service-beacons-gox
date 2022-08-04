package test_persistence

import (
	"context"
	"os"
	"testing"

	persist "github.com/pip-services-samples/service-beacons-gox/persistence"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type BeaconsPostgresPersistenceTest struct {
	persistence *persist.BeaconsPostgresPersistence
	fixture     *BeaconsPersistenceFixture
}

func newBeaconsPostgresPersistenceTest() *BeaconsPostgresPersistenceTest {
	postgresUri := os.Getenv("POSTGRES_URI")
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		postgresHost = "localhost"
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		postgresPort = "5432"
	}

	postgresDatabase := os.Getenv("POSTGRES_DB")
	if postgresDatabase == "" {
		postgresDatabase = "test"
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		postgresUser = "postgres"
	}
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		postgresPassword = "postgres#"
	}

	if postgresUri == "" && postgresHost == "" {
		panic("Connection params not set")
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.uri", postgresUri,
		"connection.host", postgresHost,
		"connection.port", postgresPort,
		"connection.database", postgresDatabase,
		"credential.username", postgresUser,
		"credential.password", postgresPassword,
		"schema", "test_schema",
	)

	persistence := persist.NewBeaconsPostgresPersistence()
	persistence.Configure(context.Background(), dbConfig)

	fixture := NewBeaconsPersistenceFixture(persistence)

	return &BeaconsPostgresPersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func (c *BeaconsPostgresPersistenceTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *BeaconsPostgresPersistenceTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsPostgresPersistence(t *testing.T) {
	c := newBeaconsPostgresPersistenceTest()
	if c == nil {
		return
	}

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("Get With Filters", c.fixture.TestGetWithFilters)
	c.teardown(t)
}
