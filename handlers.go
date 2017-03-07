package main

import (
	"net/http"

	"./uaa"
	"github.com/gorilla/mux"
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

func listGroupsHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		groups, err := uaac.ListGroups()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "groups/list", groups)
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
func userHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		user, err := uaac.User(vars["id"])
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "users/user", user)
	}
}

func groupHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		group, err := uaac.Group(vars["id"])
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}

		r.HTML(w, http.StatusOK, "groups/group", group)
	}
}

func listIdentityProvidersHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		identityProviders, err := uaac.ListIdentityProviders()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}
		r.HTML(w, http.StatusOK, "identity-providers/list", identityProviders)
	}
}

func IdentityProviderHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		identityProvider, err := uaac.IdentityProvider(vars["id"])
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}
		r.HTML(w, http.StatusOK, "identity-providers/identity-provider", identityProvider)
	}
}

func listSamlServiceProvidersHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		samlServiceProviders, err := uaac.ListSamlServiceProviders()
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}
		r.HTML(w, http.StatusOK, "saml/service-providers/list", samlServiceProviders)
	}
}

func SamlServiceProviderHandler(r *render.Render, uaac uaa.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		samlServiceProvider, err := uaac.SamlServiceProvider(vars["id"])
		if err != nil {
			r.Text(w, http.StatusInternalServerError, err.Error())
		}
		r.HTML(w, http.StatusOK, "saml/service-providers/service-provider", samlServiceProvider)
	}
}
