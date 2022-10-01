package geo

import "strconv"

type JsonObject interface{}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func (coord *Coordinate) toString() string {
	decimals := 8
	return strconv.FormatFloat(coord.Longitude, 'f', decimals, 64) + "," + strconv.FormatFloat(coord.Latitude, 'f', decimals, 64)
}
