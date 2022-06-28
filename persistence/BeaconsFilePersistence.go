package persistence

import (
	"context"
	data1 "github.com/pip-services-samples/service-beacons-gox/data/version1"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cpersist "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

type BeaconsFilePersistence struct {
	BeaconsMemoryPersistence
	persister *cpersist.JsonFilePersister[data1.BeaconV1]
}

func NewBeaconsFilePersistence(path string) *BeaconsFilePersistence {
	c := BeaconsFilePersistence{
		BeaconsMemoryPersistence: *NewBeaconsMemoryPersistence(),
	}
	c.persister = cpersist.NewJsonFilePersister[data1.BeaconV1](path)
	c.IdentifiableMemoryPersistence.Loader = c.persister
	c.IdentifiableMemoryPersistence.Saver = c.persister
	return &c
}

func (c *BeaconsFilePersistence) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.BeaconsMemoryPersistence.Configure(ctx, config)
	c.persister.Configure(ctx, config)
}
