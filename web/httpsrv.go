package web

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "See where am I: %s!", r.URL.Path[1:])
}

func StartHttpSrv() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
