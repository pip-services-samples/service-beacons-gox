package test_persistence

import (
	"context"
	"testing"

	persist "github.com/pip-services-samples/service-beacons-gox/persistence"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type BeaconsFilePersistenceTest struct {
	persistence *persist.BeaconsFilePersistence
	fixture     *BeaconsPersistenceFixture
}

func newBeaconsFilePersistenceTest() *BeaconsFilePersistenceTest {
	persistence := persist.NewBeaconsFilePersistence("")
	persistence.Configure(context.Background(), cconf.NewConfigParamsFromTuples(
		"path", "../../temp/beacons.test.json",
	))

	fixture := NewBeaconsPersistenceFixture(persistence)

	return &BeaconsFilePersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func (c *BeaconsFilePersistenceTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *BeaconsFilePersistenceTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsFilePersistence(t *testing.T) {
	c := newBeaconsFilePersistenceTest()
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
