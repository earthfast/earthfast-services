package geo

var (
	Unknown = newRegion("unknown", 90, 0)

	Africa       = newRegion("af", -18.508389, 23.326083)
	Asia         = newRegion("as", 43.681111, 87.331111)
	Australia    = newRegion("au", -23.116667, 132.133333)
	Europe       = newRegion("eu", 55.4847, 28.7761)
	NorthAmerica = newRegion("na", 48.166667, -100.166667)
	SouthAmerica = newRegion("sa", -15.6006, -56.1004)

	regionsByID = map[string]Region{
		Africa.ID:       Africa,
		Asia.ID:         Asia,
		Australia.ID:    Australia,
		Europe.ID:       Europe,
		NorthAmerica.ID: NorthAmerica,
		SouthAmerica.ID: SouthAmerica,
	}
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func (c Coordinate) IsZero() bool {
	return c.Latitude == 0.0 && c.Longitude == 0.0
}

type Region struct {
	ID     string
	Center Coordinate
}

func newRegion(id string, centerLat, centerLon float64) Region {
	return Region{
		ID: id,
		Center: Coordinate{
			Latitude:  centerLat,
			Longitude: centerLon,
		},
	}
}

func (r Region) Distance(to Coordinate) float64 {
	return Distance(r.Center, to)
}

func GetRegion(id string) Region {
	if r, ok := regionsByID[id]; ok {
		return r
	}
	return Unknown
}
