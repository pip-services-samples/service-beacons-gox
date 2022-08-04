package persistence

import (
	"context"
	data1 "github.com/pip-services-samples/service-beacons-gox/data/version1"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cpg "github.com/pip-services3-gox/pip-services3-postgres-gox/persistence"
	"strings"
)

type BeaconsPostgresPersistence struct {
	cpg.IdentifiableJsonPostgresPersistence[data1.BeaconV1, string]
}

func NewBeaconsPostgresPersistence() *BeaconsPostgresPersistence {
	c := &BeaconsPostgresPersistence{}
	c.IdentifiableJsonPostgresPersistence = *cpg.InheritIdentifiableJsonPostgresPersistence[data1.BeaconV1, string](c, "beacons")
	return c
}

func (c *BeaconsPostgresPersistence) DefineSchema() {
	c.ClearSchema()
	c.IdentifiableJsonPostgresPersistence.DefineSchema()
	c.EnsureTable("", "")
	c.EnsureIndex(c.TableName+"_key", map[string]string{"(data->'type')": "1"}, nil)
	c.EnsureIndex(c.TableName+"_key", map[string]string{"(data->'udi')": "1"}, nil)
}

func (c *BeaconsPostgresPersistence) GetPageByFilter(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (cdata.DataPage[data1.BeaconV1], error) {
	filterObj := ""
	if id, ok := filter.GetAsNullableString("id"); ok && id != "" {
		filterObj += "id='" + id + "'"
	}
	if siteId, ok := filter.GetAsNullableString("site_id"); ok && siteId != "" {
		filterObj += "data->>'site_id'='" + siteId + "'"
	}
	if typeId, ok := filter.GetAsNullableString("type"); ok && typeId != "" {
		filterObj += "data->>'type'='" + typeId + "'"
	}
	if udi, ok := filter.GetAsNullableString("udi"); ok && udi != "" {
		filterObj += "data->>'udi'='" + udi + "'"
	}
	if label, ok := filter.GetAsNullableString("label"); ok && label != "" {
		filterObj += "data->>'label'='" + label + "'"
	}
	if udis, ok := filter.GetAsObject("udis"); ok {
		udisStr := ""
		switch _udis := udis.(type) {
		case []string:
			if len(_udis) > 0 {
				udisStr = "['" + strings.Join(_udis, "','") + "']"
			}
			break
		case string:
			if _udisArr := strings.Split(_udis, ","); len(_udisArr) > 0 {
				udisStr = "['" + strings.Join(_udisArr, "','") + "']"
			}
			break
		}
		if len(udisStr) > 0 {
			filterObj += "data->>'udi'= ANY (ARRAY " + udisStr + ")"
		}
	}

	return c.IdentifiableJsonPostgresPersistence.GetPageByFilter(ctx, correlationId,
		filterObj, paging,
		"", "",
	)
}

func (c *BeaconsPostgresPersistence) GetOneByUdi(ctx context.Context, correlationId string, udi string) (data1.BeaconV1, error) {

	paging := *cdata.NewPagingParams(0, 1, false)
	page, err := c.IdentifiableJsonPostgresPersistence.GetPageByFilter(ctx, correlationId,
		"data->>'udi'='"+udi+"'", paging,
		"", "",
	)
	if err != nil {
		return data1.BeaconV1{}, err
	}
	if page.HasData() {
		return page.Data[0], nil
	}
	return data1.BeaconV1{}, nil
}
