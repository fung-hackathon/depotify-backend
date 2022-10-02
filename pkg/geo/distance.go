package geo

import (
	"encoding/json"
	"funhackathon2022-backend/pkg/config"
	"funhackathon2022-backend/pkg/models"
	"net/http"
)

type Distance float64

func GetDistance(ca, cb models.Coordinate) (Distance, error) {
	base := "https://map.yahooapis.jp/dist/V1/distance"
	appid := config.YOLP_APPID
	output := "json"

	response, err :=
		http.Get(base + "?appid=" + appid +
			"&coordinates=" + ca.ToString() + "%20" + cb.ToString() + "&output=" + output)

	if err != nil {
		return Distance(0.), ErrFailedToGetAccessToYOLP
	}

	var obj interface{}

	if err := json.NewDecoder(response.Body).Decode(&obj); err != nil {
		return Distance(0.), ErrUnableToDecodeResponse
	}

	distance :=
		Distance(obj.(map[string]interface{})["Feature"].([]interface{})[0].(map[string]interface{})["Geometry"].(map[string]interface{})["Distance"].(float64))

	return distance, nil
}
