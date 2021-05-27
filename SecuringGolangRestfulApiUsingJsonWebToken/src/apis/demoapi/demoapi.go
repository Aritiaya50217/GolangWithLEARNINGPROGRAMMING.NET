package demoapi

import (
	"net/http"
	"fmt"
)

func Demo1(w http.ResponseWriter,r*http.Request) {
	fmt.Fprint(w,"Demo 1 Api")
}

func Demo2(w http.ResponseWriter,r*http.Request) {
	fmt.Fprint(w,"Demo 2 Api")
}