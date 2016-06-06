package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dave-malone/uaa-ui/uaa"
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
	uaac := getUaac()
	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/", serverInfoHandler(r, uaac)).Methods("GET")
	router.HandleFunc("/zones", listZonesHandler(r, uaac)).Methods("GET")
	router.HandleFunc("/clients", listOauthClientsHandler(r, uaac)).Methods("GET")

	n.UseHandler(router)

	addy := ":" + port

	n.Run(addy)
	fmt.Printf("Server running at %v\n", addy)
}

func getUaac() *uaa.Client {
	serverURL := os.Getenv("UAA_URL")
	clientID := os.Getenv("UAA_CLIENT_ID")
	clientSecret := os.Getenv("UAA_CLIENT_SECRET")
	uaac, err := uaa.NewClient(serverURL, clientID, clientSecret)
	if err != nil {
		panic("Failed to initialize uaa client; check your UAA_URL and your UAA_CLIENT_ID")
	}

	return uaac
}

func serverInfoHandler(r *render.Render, uaac *uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		info, err := uaac.GetServerInfo()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "info", info)
	}
}

func listOauthClientsHandler(r *render.Render, uaac *uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		clients, err := uaac.ListOauthClients()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "clients/list", clients)
	}
}

func listZonesHandler(r *render.Render, uaac *uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		zones, err := uaac.ListZones()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "zones/list", zones)
	}
}
