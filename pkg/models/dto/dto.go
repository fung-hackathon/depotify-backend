package dto

import "strconv"

type (
	Health struct {
		Message string `json:"message"`
	}

	Error struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	}

	Score struct {
		UserId string `json:"userid"`
		Score  int64  `json:"score"`
	}

	UserId struct {
		UserId string `json:"userid"`
	}

	Coordinate struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
	}

	Coordinates struct {
		UserId      string       `json:"userid"`
		Coordinates []Coordinate `json:"coordinates"`
	}

	Emotion struct {
		UserId  string   `json:"userid"`
		Emotion []string `json:"emotion"`
	}
)

func (coord *Coordinate) ToString() string {
	decimals := 8
	return strconv.FormatFloat(coord.Longitude, 'f', decimals, 64) + "," + strconv.FormatFloat(coord.Latitude, 'f', decimals, 64)
}
