package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCountry(w, r)
		} else {
			if r.Method == http.MethodPost {
				addCountry(w, r)
			} else { //This shows when the method used was different to GET
				w.WriteHeader(http.StatusMethodNotAllowed) //This show us the corresponding method header
				fmt.Fprintf(w, "Method not allowed")
				return
			}
		}
	})
}
