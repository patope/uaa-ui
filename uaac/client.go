package uaac

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetServerInfo(at AccessToken) (ServerInfo, error) {
	var info ServerInfo

	client := &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("UAA_URL")+"/info", nil)
	req.Header.Set("Authorization", "Bearer "+at.Token)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return info, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return info, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &info)
	if err != nil {
		return info, fmt.Errorf("Unable to unmarshall json response to type ServerInfo; error: %v response: %s", err.Error(), string(body))
	}

	return info, nil
}

func GetClientToken() (AccessToken, error) {
	var at AccessToken

	params := url.Values{}
	params.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	params.Set("client_id", os.Getenv("CLIENT_ID"))
	params.Set("grant_type", "client_credentials")
	params.Set("response_type", "token")

	client := &http.Client{}

	req, err := http.NewRequest("POST", os.Getenv("UAA_URL")+"/oauth/token", strings.NewReader(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return at, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return at, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return at, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &at)
	if err != nil {
		return at, err
	}

	return at, nil
}

func ListOauthClients(at AccessToken) (OauthClients, error) {
	var clients OauthClients

	client := &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("UAA_URL")+"/oauth/clients", nil)
	req.Header.Set("Authorization", "Bearer "+at.Token)
	if err != nil {
		return clients, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return clients, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clients, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &clients)
	if err != nil {
		return clients, fmt.Errorf("Unable to unmarshall json response to type OauthClients; error: %v response: %s", err.Error(), string(body))
	}

	return clients, nil
}

func ListZones(at AccessToken) ([]Zone, error) {
	var z []Zone

	client := &http.Client{}
	//TODO - get the base for the url from an environment variable
	req, err := http.NewRequest("GET", os.Getenv("UAA_URL")+"/identity-zones", nil)
	req.Header.Set("Authorization", "Bearer "+at.Token)
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return z, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return z, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return z, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &z)
	if err != nil {
		return z, fmt.Errorf("Unable to unmarshall json response to a slice of Zones; error: %v response: %s", err.Error(), string(body))
	}

	return z, nil
}
