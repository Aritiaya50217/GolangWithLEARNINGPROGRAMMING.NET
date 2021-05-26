package apis

import (
	"net/http"
	"fmt"
)
func Demo2API(w http.ResponseWriter,r*http.Request) {
	fmt.Fprintf(w,"Demo2 API")
}
