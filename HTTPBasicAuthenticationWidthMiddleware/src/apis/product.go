package apis

import (
	"fmt"
	"net/http"
)
func FindAll(w http.ResponseWriter,r*http.Request){
	fmt.Fprint(w,"Find All")
}

func Search(w http.ResponseWriter,r*http.Request){
	fmt.Fprint(w,"Search Product")
}

