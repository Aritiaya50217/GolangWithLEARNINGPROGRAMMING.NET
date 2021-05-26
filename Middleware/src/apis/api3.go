package apis

import (
	"net/http"
	"fmt"
)
func Demo3API(w http.ResponseWriter,r*http.Request) {
	fmt.Fprintf(w,"Demo3 API")
}
