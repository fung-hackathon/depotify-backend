package dto

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
		Latitude  string `json:"lat"`
		Longitude string `json:"lng"`
	}
	/*
		Coordinates struct {
			UserId      string       `json:"userid"`
			Coordinates []Coordinate `json:"coordinates"`
		}
	*/

	Emotion struct {
		UserId  string   `json:"userid"`
		Emotion []string `json:"emotion"`
	}
)

func (coord *Coordinate) ToString() string {
	return coord.Longitude + "," + coord.Latitude
}
