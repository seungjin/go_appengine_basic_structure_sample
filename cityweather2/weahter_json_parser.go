package cityweather2

import "encoding/json"

type Response struct {
	Name    string    `json:"name"`
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

type Main struct {
	Temp     float32
	Pressure float32
	Humidity float32
	Temp_min float32
	Temp_max float32
}

type Weather struct {
	Main        string
	Description string
}

func Weather_json_parser(weather_json []byte) (string, string, float32, float32, float32) {

	weather := &Response{}
	err := json.Unmarshal(weather_json, &weather)
	if err != nil {
		panic(err)
	}

	return weather.Name, weather.Weather[0].Description, weather.Main.Temp, weather.Main.Temp_min, weather.Main.Temp_max

}
