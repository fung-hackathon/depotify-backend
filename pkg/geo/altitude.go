package geo

import (
	"encoding/json"
	"funhackathon2022-backend/pkg/config"
	"net/http"
)

type Altitude float64

func GetAltitude(coord Coordinate) (Altitude, error) {
	base := "https://map.yahooapis.jp/alt/V1/getAltitude"
	appid := config.YOLP_APPID
	output := "json"

	response, err :=
		http.Get(base + "?appid=" + appid +
			"&coordinates=" + coord.toString() + "&output=" + output)

	if err != nil {
		return Altitude(0.), ErrFailedToGetAccessToYOLP
	}

	var obj JsonObject

	if err := json.NewDecoder(response.Body).Decode(&obj); err != nil {
		return Altitude(0.), ErrUnableToDecodeResponse
	}

	altitude :=
		Altitude(obj.(map[string]interface{})["Feature"].([]interface{})[0].(map[string]interface{})["Property"].(map[string]interface{})["Altitude"].(float64))

	return altitude, nil
}
