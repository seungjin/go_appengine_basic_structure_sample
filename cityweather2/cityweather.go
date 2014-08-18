package cityweather2

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"appengine"
	"appengine/urlfetch"
)

func CityWeather(r *http.Request, message chan []byte, cityname string) {

	weather_info := openweather_request(r, cityname)

	/* NEVER NEVER NEVER DO THIS.. float to string is way way slow!!!
	weather := &Response{}
	err := json.Unmarshal(weather_info, &weather)
	if err != nil {
		panic(err)
	}

	message <- []string{
		weather.Name,
		weather.Weather[0].Description,
		strconv.FormatFloat(float64(weather.Main.Temp), 'f', 2, 32),
		strconv.FormatFloat(float64(weather.Main.Temp_min), 'f', 2, 32),
		strconv.FormatFloat(float64(weather.Main.Temp_max), 'f', 2, 32)}
	*/
	time.Sleep(time.Second * 1)
	message <- weather_info
}

//http://api.openweathermap.org/data/2.5/weather?q=

func openweather_request(r *http.Request, cityname string) []byte {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get("http://api.openweathermap.org/data/2.5/weather?q=" + strings.Replace(cityname, " ", "%20", -1) + "&units=metric")
	//c.Infof("Just called %v", cityname)
	if err != nil {
		c.Errorf("Error from here: %v", err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("Error from here: %v", err)
	}
	return contents
}
