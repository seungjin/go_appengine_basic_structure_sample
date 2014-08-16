package myapp

import (
	"fmt"
	"net/http"
)

func testFunc1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello testFunc1")
}
