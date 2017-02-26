package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"./uaa"
	"github.com/unrolled/render"
)

var (
	r    = render.New()
	uaac = new(uaa.FakeUaac)
)

func TestGetServerInfo(t *testing.T) {
	serverInfo := getServerInfo()
	handler := serverInfoHandler(r, uaac)
	expectedBody := getAndExecuteHTMLTemplate("templates/info.tmpl", serverInfo)

	uaac.GetServerInfoReturns(serverInfo, nil)
	responseBody, err := executeRequest(handler, "GET", nil, http.StatusOK)
	if err != nil {
		t.Errorf("%v", err)
	}

	if bytes.Compare(expectedBody, responseBody) != 0 {
		t.Errorf("The expected HTML was not generated in the call to serverInfoHandler: \n%s\n\n%s", string(expectedBody), string(responseBody))
	}
}

func TestGetOauthClients(t *testing.T) {
	testOauthClients := new(uaa.OauthClients)
	handler := listOauthClientsHandler(r, uaac)
	expectedBody := getAndExecuteHTMLTemplate("templates/clients/list.tmpl", testOauthClients)

	uaac.ListOauthClientsReturns(*testOauthClients, nil)
	responseBody, err := executeRequest(handler, "GET", nil, http.StatusOK)
	if err != nil {
		t.Errorf("%v", err)
	}

	if bytes.Compare(expectedBody, responseBody) != 0 {
		t.Errorf("The expected HTML was not generated in the call to listOauthClientsHandler: \n\n%s\n\n%s", string(expectedBody), string(responseBody))
	}
}

func TestGetIdentityZones(t *testing.T) {
	identityZones := getIdentityZones()
	handler := listIdentityZonesHandler(r, uaac)
	expectedBody := getAndExecuteHTMLTemplate("templates/zones/list.tmpl", identityZones)

	uaac.ListIdentityZonesReturns(identityZones, nil)
	responseBody, err := executeRequest(handler, "GET", nil, http.StatusOK)
	if err != nil {
		t.Errorf("%v", err)
	}

	if bytes.Compare(expectedBody, responseBody) != 0 {
		t.Errorf("The expected HTML was not generated in the call to listIdentityZonesHandler: \n\n%s\n\n%s", string(expectedBody), string(responseBody))
	}
}

func executeRequest(handler http.Handler, requestMethod string, requestBody io.Reader, expectedHTTPResponseCode int) ([]byte, error) {
	var body []byte

	server := httptest.NewServer(handler)
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL, requestBody)
	if err != nil {
		return body, fmt.Errorf("Error in creating %s request for %v: %v", requestMethod, handler, err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return body, fmt.Errorf("Error in %s to %v: %v", requestMethod, handler, err)
	}

	defer res.Body.Close()
	if res.StatusCode != expectedHTTPResponseCode {
		return body, fmt.Errorf("Expected status code of %d but got %d", expectedHTTPResponseCode, res.StatusCode)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return body, fmt.Errorf("Failed to read response body: %v", err)
	}

	return body, nil
}

func getAndExecuteHTMLTemplate(templateFileName string, data interface{}) []byte {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		panic("Failed to Parse template files " + err.Error())
	}

	buffer := bytes.NewBuffer(make([]byte, 0))
	err = t.Execute(buffer, data)
	if err != nil {
		panic("Failed to execute template " + templateFileName + "; error: " + err.Error())
	}

	return buffer.Bytes()
}

func getServerInfo() uaa.ServerInfo {
	responseBody := []byte(`{
      "app" : {
        "version" : "3.5.0-SNAPSHOT"
      },
      "links" : {},
      "zone_name" : "uaa",
      "entityID" : "cloudfoundry-saml-login",
      "commit_id" : "6681d65",
      "idpDefinitions" : {
        "SAML" : "http://localhost:8080/uaa/saml/discovery?returnIDParam=idp&entityID=cloudfoundry-saml-login&idp=SAML&isPassive=true"
      },
      "prompts" : {
        "username" : [ "text", "Email" ],
        "password" : [ "password", "Password" ]
      },
      "timestamp" : "2016-05-18T18:20:54+0000"
    }`)

	var info uaa.ServerInfo

	err := json.Unmarshal(responseBody, &info)
	if err != nil {
		panic("Failed to unmarshall json to ServerInfo: " + err.Error())
	}

	return info
}

func getOauthClients() uaa.OauthClients {
	clients := make([]uaa.OauthClient, 1)
	clients[0] = uaa.OauthClient{
		ID:                     "test id",
		Name:                   "test name",
		Action:                 "test action",
		Scope:                  []string{"scope1"},
		ResourceIDs:            []string{"resourceid1"},
		Authorities:            []string{"authority1"},
		AuthorizedGrantTypes:   []string{"authorizedGrantType1"},
		LastModified:           123,
		RedirectURI:            []string{"redirectUri1"},
		SignupRedirectURL:      "signupRedirectUrl1",
		ChangeEmailRedirectURL: "ChangeEmailRedirectURL1",
	}

	return uaa.OauthClients{
		Clients:      clients,
		StartIndex:   1,
		ItemsPerPage: 10,
		TotalResults: 1,
		Schemas:      []string{"schemaUrl1"},
	}
}

func getIdentityZones() []uaa.IdentityZone {
	identityZones := make([]uaa.IdentityZone, 1)
	identityZones[0] = uaa.IdentityZone{
		ID:           "TestID",
		Subdomain:    "Test.SubDomain",
		Name:         "TestName",
		Version:      1,
		Created:      1234,
		LastModified: 5678,
	}

	return identityZones
}
