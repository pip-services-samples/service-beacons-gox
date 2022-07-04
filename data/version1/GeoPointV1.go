package data1

type GeoPointV1 struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float32 `json:"coordinates" bson:"coordinates"`
}

func (g GeoPointV1) Clone() GeoPointV1 {

	coords := make([]float32, len(g.Coordinates))
	for i := range g.Coordinates {
		coords[i] = g.Coordinates[i]
	}
	return GeoPointV1{
		Coordinates: coords,
		Type:        g.Type,
	}
}
