package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dave-malone/uaa-fe/uaac"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	r := render.New()

	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/", defaultHandler).Methods("GET")
	router.HandleFunc("/zones", listZonesHandler(r)).Methods("GET")

	n.UseHandler(router)

	addy := ":" + port

	n.Run(addy)
	fmt.Printf("Server running at %v\n", addy)
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func listZonesHandler(r *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		token, err := uaac.GetClientToken()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}
		zones, err := uaac.ListZones(token)
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "zones/list", zones)
	}
}
