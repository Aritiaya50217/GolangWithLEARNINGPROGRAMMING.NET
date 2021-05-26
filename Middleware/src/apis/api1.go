package apis

import (
	"net/http"
	"fmt"
)
func Demo1API(w http.ResponseWriter,r*http.Request) {
	fmt.Fprintf(w,"Demo1 API")
}
