package myapp

import (
	"fmt"
	"time"
)
import "net/http"

import "mymodule"
import "cityweather"
import "cityweather2"

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test1", testFunc1)
	http.HandleFunc("/mymodule_dummyfunc1", call_module1_dummyfunc1)
	http.HandleFunc("/cityweather", call_cityweather)
	http.HandleFunc("/cityweather2", call_cityweather2)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func call_module1_dummyfunc1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mymodule.DummyFunc1())
}

func call_cityweather(w http.ResponseWriter, r *http.Request) {
	now1 := time.Now().UnixNano()
	fmt.Fprint(w, "<p>City Weather</p>")

	city_list := []string{"seoul", "newyork", "beijing", "london"}

	for _, city := range city_list {
		weather_description, temp, temp_min, temp_max := cityweather.CityWeather(r, city)
		fmt.Fprintf(w, "<li> %s: %s (Temp: %.2f, Min:%.2f/Max:%.2f)", city, weather_description, temp, temp_min, temp_max)
	}
	now2 := time.Now().UnixNano()
	fmt.Fprintf(w, "<br/><br/>> %d", now2-now1)

}

func call_cityweather2(w http.ResponseWriter, r *http.Request) {
	now1 := time.Now().UnixNano()
	fmt.Fprint(w, "<p>City Weather</p>")

	city_list := []string{"seoul", "newyork", "beijing", "london"}

	message1 := make(chan string, 4)
	message2 := make(chan []float32, 4)

	for _, city := range city_list {
		//message1 := make(chan string)
		//message2 := make(chan []float32)
		go cityweather2.CityWeather(r, message1, message2, city)
		//weather_description := <-message1
		//temp_info := <-message2
		//fmt.Fprintf(w, <-message1)
		//fmt.Fprintf(w, "<li> %s: %s (Temp: %.2f, Min:%.2f/Max:%.2f)", city, weather_description, temp_info[0], temp_info[1], temp_info[2])
	}

	for _, city := range city_list {
		select {
		case weather_description := <-message1:
			temp_info := <-message2
			fmt.Fprintf(w, "<li> %s: %s (Temp: %.2f, Min:%.2f/Max:%.2f)", city, weather_description, temp_info[0], temp_info[1], temp_info[2])
		case <-time.After(time.Second * 2):
			fmt.Fprintf(w, "Timeout!")
		}
	}

	//close(message1)
	//close(message2)

	now2 := time.Now().UnixNano()
	fmt.Fprintf(w, "<br/><br/>> %d", now2-now1)

}
