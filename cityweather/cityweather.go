package cityweather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

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

func CityWeather(r *http.Request, cityname string) (string, float32, float32, float32) {

	weather_info := openweather_request(r, cityname)

	weather := &Response{}

	err := json.Unmarshal(weather_info, &weather)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 1)
	return weather.Weather[0].Description, weather.Main.Temp, weather.Main.Temp_min, weather.Main.Temp_max
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
