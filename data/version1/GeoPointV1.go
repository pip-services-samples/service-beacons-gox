package data1

type GeoPointV1 struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates [][]float32 `json:"coordinates" bson:"coordinates"`
}

func (g GeoPointV1) Clone() GeoPointV1 {

	coords := make([][]float32, len(g.Coordinates))
	for i := 0; i < len(g.Coordinates); i++ {
		coords[i] = make([]float32, len(g.Coordinates[i]))
		for j := 0; j < len(g.Coordinates[i]); j++ {
			coords[i][j] = g.Coordinates[i][j]
		}
	}
	return GeoPointV1{
		Coordinates: coords,
		Type:        g.Type,
	}
}
