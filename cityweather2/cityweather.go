package cityweather2

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"appengine"
	"appengine/urlfetch"
)

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

func CityWeather(r *http.Request, message1 chan string, message2 chan []float32, cityname string) {

	weather_info := openweather_request(r, cityname)

	weather := &Response{}

	err := json.Unmarshal(weather_info, &weather)
	if err != nil {
		panic(err)
	}

	//return weather.Weather[0].Description, weather.Main.Temp, weather.Main.Temp_min, weather.Main.Temp_max
	/*
		messages <- []string{weather.Weather[0].Description,
			strconv.FormatFloat(float64(weather.Main.Temp), 'f', 2, 32),
			strconv.FormatFloat(float64(weather.Main.Temp_min), 'f', 2, 32),
			strconv.FormatFloat(float64(weather.Main.Temp_max), 'f', 2, 32)}
	*/
	message1 <- weather.Weather[0].Description
	message2 <- []float32{weather.Main.Temp, weather.Main.Temp_min, weather.Main.Temp_max}
}

//http://api.openweathermap.org/data/2.5/weather?q=

func openweather_request(r *http.Request, cityname string) []byte {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get("http://api.openweathermap.org/data/2.5/weather?q=" + cityname + "&units=metric")
	if err != nil {
		os.Exit(1)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}
	return contents
}
