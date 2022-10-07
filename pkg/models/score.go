package models

import (
	"funhackathon2022-backend/pkg/geo"
	"funhackathon2022-backend/pkg/models/dto"
	"math"
)

func CalculateScore(ca, cb dto.Coordinate) (int64, error) {

	var err error

	distance, err := geo.GetDistanceM(ca, cb)
	if err != nil {
		return 0, err
	}

	altitude_a, err := geo.GetAltitudeM(ca)
	if err != nil {
		return 0, err
	}

	altitude_b, err := geo.GetAltitudeM(cb)
	if err != nil {
		return 0, err
	}

	return int64(float64(distance)/10. + math.Abs(float64(altitude_a)-float64(altitude_b))), nil
}
