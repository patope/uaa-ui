package main

import (
	"fmt"
	"os"

	"./uaa"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	r := render.New(render.Options{
		Layout: "layout",
	})
	uaac := getUaac()
	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/", serverInfoHandler(r, uaac)).Methods("GET")
	router.HandleFunc("/zones", listIdentityZonesHandler(r, uaac)).Methods("GET")
	router.HandleFunc("/clients", listOauthClientsHandler(r, uaac)).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler(r, uaac)).Methods("GET")
	router.HandleFunc("/users", listUsersHandler(r, uaac)).Methods("GET")
	router.HandleFunc("/groups", listGroupsHandler(r, uaac)).Methods("GET")
	router.HandleFunc("/groups/{id}", groupHandler(r, uaac)).Methods("GET")

	n.UseHandler(router)

	addy := ":" + port

	n.Run(addy)
	fmt.Printf("Server running at %v\n", addy)
}

func getUaac() uaa.Client {
	serverURL := os.Getenv("UAA_URL")
	clientID := os.Getenv("UAA_CLIENT_ID")
	clientSecret := os.Getenv("UAA_CLIENT_SECRET")
	uaac, err := uaa.NewClient(serverURL, clientID, clientSecret)
	if err != nil {
		panic("Failed to initialize uaa client; check your UAA_URL and your UAA_CLIENT_ID")
	}

	return uaac
}
