package main

import (
	"net/http"

	"github.com/dave-malone/uaa-ui/uaa"
	"github.com/unrolled/render"
)

func serverInfoHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		info, err := uaac.GetServerInfo()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "info", info)
	}
}

func listOauthClientsHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		clients, err := uaac.ListOauthClients()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "clients/list", clients)
	}
}

func listIdentityZonesHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		zones, err := uaac.ListIdentityZones()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "zones/list", zones)
	}
}

func listUsersHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		users, err := uaac.ListUsers()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "users/list", users)
	}
}
