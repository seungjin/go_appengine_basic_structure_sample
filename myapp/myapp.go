package myapp

import "fmt"
import "net/http"

import "mymodule"

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test1", testFunc1)
	http.HandleFunc("/mymodule_dummyfunc1", call_module1_dummyfunc1)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func call_module1_dummyfunc1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mymodule.DummyFunc1())
}
