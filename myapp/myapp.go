package myapp

import "fmt"
import "net/http"

import "mymodule"
import "cityweather"

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test1", testFunc1)
	http.HandleFunc("/mymodule_dummyfunc1", call_module1_dummyfunc1)
	http.HandleFunc("/cityweather", call_cityweather)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func call_module1_dummyfunc1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mymodule.DummyFunc1())
}

func call_cityweather(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<p>City Weather</p>")

	city_list := []string{"seoul", "newyork", "beijing", "london"}

	for _, city := range city_list {
		weather_description, temp, temp_min, temp_max := cityweather.CityWeather(r, city)
		fmt.Fprintf(w, "<li> %s: %s (Temp: %.2f, Min:%.2f/Max:%.2f)", city, weather_description, temp, temp_min, temp_max)
	}

}
