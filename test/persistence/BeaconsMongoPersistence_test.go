package test_persistence

import (
	"context"
	"os"
	"testing"

	persist "github.com/pip-services-samples/service-beacons-gox/persistence"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type BeaconsMongoPersistenceTest struct {
	persistence *persist.BeaconsMongoPersistence
	fixture     *BeaconsPersistenceFixture
}

func newBeaconsMongoPersistenceTest() *BeaconsMongoPersistenceTest {
	mongoUri := os.Getenv("MONGO_SERVICE_URI")
	mongoHost := os.Getenv("MONGO_SERVICE_HOST")

	if mongoHost == "" {
		mongoHost = "localhost"
	}
	mongoPort := os.Getenv("MONGO_SERVICE_PORT")
	if mongoPort == "" {
		mongoPort = "27017"
	}

	mongoDatabase := os.Getenv("MONGO_SERVICE_DB")
	if mongoDatabase == "" {
		mongoDatabase = "test"
	}

	// Exit if mongo connection is not set
	if mongoUri == "" && mongoHost == "" {
		return nil
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.uri", mongoUri,
		"connection.host", mongoHost,
		"connection.port", mongoPort,
		"connection.database", mongoDatabase,
	)

	persistence := persist.NewBeaconsMongoPersistence()
	persistence.Configure(context.Background(), dbConfig)

	fixture := NewBeaconsPersistenceFixture(persistence)

	return &BeaconsMongoPersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func (c *BeaconsMongoPersistenceTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *BeaconsMongoPersistenceTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsMongoPersistence(t *testing.T) {
	c := newBeaconsMongoPersistenceTest()
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
