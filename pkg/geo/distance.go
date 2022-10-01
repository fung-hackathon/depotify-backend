package geo

import (
	"encoding/json"
	"net/http"
	"os"
)

type Distance float64

func GetDistance(ca, cb Coordinate) (Distance, error) {
	base := "https://map.yahooapis.jp/dist/V1/distance"
	appid := os.Getenv("YOLP_APPID")
	output := "json"

	response, err :=
		http.Get(base + "?appid=" + appid +
			"&coordinates=" + ca.toString() + "%20" + cb.toString() + "&output=" + output)

	if err != nil {
		return Distance(0.), ErrFailedToGetAccessToYOLP
	}

	var obj JsonObject

	if err := json.NewDecoder(response.Body).Decode(&obj); err != nil {
		return Distance(0.), ErrUnableToDecodeResponse
	}

	distance :=
		Distance(obj.(map[string]interface{})["Feature"].([]interface{})[0].(map[string]interface{})["Geometry"].(map[string]interface{})["Distance"].(float64))

	return distance, nil
}
