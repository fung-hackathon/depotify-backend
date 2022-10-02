package geo

import (
	"math"
	"testing"
)

func TestGetAltitude(t *testing.T) {

	ignorable := Altitude(1e-5)

	tests := []struct {
		name   string
		arg    Coordinate
		want   Altitude
		hasErr bool
	}{
		{
			name:   "はこだて未来大の標高(m)",
			arg:    Coordinate{Longitude: 140.766944, Latitude: 41.841806},
			want:   Altitude(133.3),
			hasErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			altitude, err := GetAltitude(tt.arg)

			if (err != nil) != tt.hasErr {
				t.Errorf("GetAltitude() error = %v, hasErr %v", err, tt.hasErr)
			}
			if math.Abs(float64(altitude-tt.want)) > float64(ignorable) {
				t.Errorf("GetAltitude() = %f, want %f", altitude, float64(tt.want))
			}

		})
	}
}

func TestGetDistance(t *testing.T) {

	ignorable := Distance(1e-5)

	type args struct {
		ca, cb Coordinate
	}

	tests := []struct {
		name   string
		arg    args
		want   Distance
		hasErr bool
	}{
		{
			name: "はこだて未来大と函館市役所の距離(Km)",
			arg: args{
				ca: Coordinate{Longitude: 140.766944, Latitude: 41.841806},
				cb: Coordinate{Longitude: 140.72892, Latitude: 41.76867},
			},
			want:   Distance(8.716124),
			hasErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distance, err := GetDistance(tt.arg.ca, tt.arg.cb)

			if (err != nil) != tt.hasErr {
				t.Errorf("GetDistance() error = %v, hasErr %v", err, tt.hasErr)
			}
			if math.Abs(float64(distance-tt.want)) > float64(ignorable) {
				t.Errorf("GetDistance() = %f, want %f", distance, float64(tt.want))
			}
		})
	}
}
