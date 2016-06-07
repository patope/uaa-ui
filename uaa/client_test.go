package uaa

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "access_token" : "4a53a3331b2445cfaca43c9af00439e8",
      "token_type" : "bearer",
      "expires_in" : 43199,
      "scope" : "clients.read emails.write scim.userids password.write idps.write notifications.write oauth.login scim.write critical_notifications.write",
      "jti" : "4a53a3331b2445cfaca43c9af00439e8"
    }`)
	}))
	defer ts.Close()

	if _, err := NewClient(ts.URL, "fake-client", "big-secret"); err != nil {
		t.Errorf("Failed to initialize client: %s", err.Error())
	}
}

func TestGetAccessToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "access_token" : "4a53a3331b2445cfaca43c9af00439e8",
      "token_type" : "bearer",
      "expires_in" : 43199,
      "scope" : "clients.read emails.write scim.userids password.write idps.write notifications.write oauth.login scim.write critical_notifications.write",
      "jti" : "4a53a3331b2445cfaca43c9af00439e8"
    }`)
	}))
	defer ts.Close()

	uaac := &uaaClient{
		serverURL:    ts.URL,
		clientID:     "clientID",
		clientSecret: "clientSecret",
	}

	accessToken, err := uaac.getAccessToken()
	if err != nil {
		t.Errorf("Failed to get access token: %s", err.Error())
	}

	if len(accessToken.Token) == 0 {
		t.Error("AccessToken.Token was blank")
	}
}

func TestGetServerInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "app" : {
        "version" : "3.5.0-SNAPSHOT"
      },
      "links" : {
        "uaa" : "http://localhost:8080/uaa",
        "passwd" : "/forgot_password",
        "login" : "http://localhost:8080/uaa",
        "register" : "/create_account"
      },
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
	}))
	defer ts.Close()

	uaac := &uaaClient{
		serverURL:    ts.URL,
		clientID:     "clientID",
		clientSecret: "clientSecret",
	}

	serverInfo, err := uaac.GetServerInfo()
	if err != nil {
		t.Errorf("Failed to get ServerInfo: %v", err)
	}

	if len(serverInfo.Version()) == 0 {
		t.Error("ServerInfo.Version() was blank")
	}
}

func TestGetListOauthClients(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "resources" : [ {
        "scope" : [ "clients.read", "clients.write" ],
        "client_id" : "vleI0g",
        "resource_ids" : [ "none" ],
        "authorized_grant_types" : [ "client_credentials" ],
        "redirect_uri" : [ "http*://ant.path.wildcard/**/passback/*", "http://test1.com" ],
        "autoapprove" : [ ],
        "authorities" : [ "clients.read", "clients.write" ],
        "lastModified" : 1463595914005
      } ],
      "startIndex" : 1,
      "itemsPerPage" : 1,
      "totalResults" : 1,
      "schemas" : [ "http://cloudfoundry.org/schema/scim/oauth-clients-1.0" ]
    }`)
	}))
	defer ts.Close()

	uaac := &uaaClient{
		serverURL:    ts.URL,
		clientID:     "clientID",
		clientSecret: "clientSecret",
	}

	clients, err := uaac.ListOauthClients()
	if err != nil {
		t.Errorf("Failed to get OauthClients: %v", err)
	}

	if len(clients.Clients) == 0 {
		t.Error("[]OauthClient was empty")
	}
}

func TestGetListZones(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[ {
      "id" : "ixv4bi58",
      "subdomain" : "ixv4bi58",
      "config" : {
        "tokenPolicy" : {
          "accessTokenValidity" : -1,
          "refreshTokenValidity" : -1,
          "jwtRevocable" : false,
          "activeKeyId" : null,
          "keys" : { }
        },
        "samlConfig" : {
          "assertionSigned" : true,
          "requestSigned" : true,
          "wantAssertionSigned" : false,
          "wantAuthnRequestSigned" : false,
          "assertionTimeToLiveSeconds" : 600,
          "certificate" : null,
          "privateKey" : null,
          "privateKeyPassword" : null
        },
        "links" : {
          "logout" : {
            "redirectUrl" : "/login",
            "redirectParameterName" : "redirect",
            "disableRedirectParameter" : true,
            "whitelist" : null
          },
          "selfService" : {
            "selfServiceLinksEnabled" : true,
            "signup" : "/create_account",
            "passwd" : "/forgot_password"
          }
        },
        "prompts" : [ {
          "name" : "username",
          "type" : "text",
          "text" : "Email"
        }, {
          "name" : "password",
          "type" : "password",
          "text" : "Password"
        }, {
          "name" : "passcode",
          "type" : "password",
          "text" : "One Time Code (Get on at /passcode)"
        } ],
        "idpDiscoveryEnabled" : false
      },
      "name" : "The Twiglet Zone",
      "version" : 0,
      "description" : "Like the Twilight Zone but tastier.",
      "created" : 1463595916851,
      "last_modified" : 1463595916851
    }, {
      "id" : "tszujxkn",
      "subdomain" : "tszujxkn",
      "config" : {
        "tokenPolicy" : {
          "accessTokenValidity" : -1,
          "refreshTokenValidity" : -1,
          "jwtRevocable" : false,
          "activeKeyId" : null,
          "keys" : { }
        },
        "samlConfig" : {
          "assertionSigned" : true,
          "requestSigned" : true,
          "wantAssertionSigned" : false,
          "wantAuthnRequestSigned" : false,
          "assertionTimeToLiveSeconds" : 600,
          "certificate" : null,
          "privateKey" : null,
          "privateKeyPassword" : null
        },
        "links" : {
          "logout" : {
            "redirectUrl" : "/login",
            "redirectParameterName" : "redirect",
            "disableRedirectParameter" : true,
            "whitelist" : null
          },
          "selfService" : {
            "selfServiceLinksEnabled" : true,
            "signup" : "/create_account",
            "passwd" : "/forgot_password"
          }
        },
        "prompts" : [ {
          "name" : "username",
          "type" : "text",
          "text" : "Email"
        }, {
          "name" : "password",
          "type" : "password",
          "text" : "Password"
        }, {
          "name" : "passcode",
          "type" : "password",
          "text" : "One Time Code (Get on at /passcode)"
        } ],
        "idpDiscoveryEnabled" : false
      },
      "name" : "The Twiglet Zone",
      "version" : 0,
      "description" : "Like the Twilight Zone but tastier.",
      "created" : 1463595918196,
      "last_modified" : 1463595918196
    }, {
      "id" : "twiglet-get-1",
      "subdomain" : "twiglet-get-1",
      "config" : {
        "tokenPolicy" : {
          "accessTokenValidity" : -1,
          "refreshTokenValidity" : -1,
          "jwtRevocable" : false,
          "activeKeyId" : null,
          "keys" : { }
        },
        "samlConfig" : {
          "assertionSigned" : true,
          "requestSigned" : true,
          "wantAssertionSigned" : false,
          "wantAuthnRequestSigned" : false,
          "assertionTimeToLiveSeconds" : 600,
          "certificate" : null,
          "privateKey" : null,
          "privateKeyPassword" : null
        },
        "links" : {
          "logout" : {
            "redirectUrl" : "/login",
            "redirectParameterName" : "redirect",
            "disableRedirectParameter" : true,
            "whitelist" : null
          },
          "selfService" : {
            "selfServiceLinksEnabled" : true,
            "signup" : "/create_account",
            "passwd" : "/forgot_password"
          }
        },
        "prompts" : [ {
          "name" : "username",
          "type" : "text",
          "text" : "Email"
        }, {
          "name" : "password",
          "type" : "password",
          "text" : "Password"
        }, {
          "name" : "passcode",
          "type" : "password",
          "text" : "One Time Code (Get on at /passcode)"
        } ],
        "idpDiscoveryEnabled" : false
      },
      "name" : "The Twiglet Zone",
      "version" : 0,
      "created" : 1463595919785,
      "last_modified" : 1463595919785
    }, {
      "id" : "twiglet-get-2",
      "subdomain" : "twiglet-get-2",
      "config" : {
        "tokenPolicy" : {
          "accessTokenValidity" : -1,
          "refreshTokenValidity" : -1,
          "jwtRevocable" : false,
          "activeKeyId" : null,
          "keys" : { }
        },
        "samlConfig" : {
          "assertionSigned" : true,
          "requestSigned" : true,
          "wantAssertionSigned" : false,
          "wantAuthnRequestSigned" : false,
          "assertionTimeToLiveSeconds" : 600,
          "certificate" : null,
          "privateKey" : null,
          "privateKeyPassword" : null
        },
        "links" : {
          "logout" : {
            "redirectUrl" : "/login",
            "redirectParameterName" : "redirect",
            "disableRedirectParameter" : true,
            "whitelist" : null
          },
          "selfService" : {
            "selfServiceLinksEnabled" : true,
            "signup" : "/create_account",
            "passwd" : "/forgot_password"
          }
        },
        "prompts" : [ {
          "name" : "username",
          "type" : "text",
          "text" : "Email"
        }, {
          "name" : "password",
          "type" : "password",
          "text" : "Password"
        }, {
          "name" : "passcode",
          "type" : "password",
          "text" : "One Time Code (Get on at /passcode)"
        } ],
        "idpDiscoveryEnabled" : false
      },
      "name" : "The Twiglet Zone",
      "version" : 0,
      "created" : 1463595919814,
      "last_modified" : 1463595919814
    }, {
      "id" : "uaa",
      "subdomain" : "",
      "config" : {
        "tokenPolicy" : {
          "accessTokenValidity" : 43200,
          "refreshTokenValidity" : 2592000,
          "jwtRevocable" : false,
          "activeKeyId" : null,
          "keys" : { }
        },
        "samlConfig" : {
          "assertionSigned" : true,
          "requestSigned" : true,
          "wantAssertionSigned" : false,
          "wantAuthnRequestSigned" : false,
          "assertionTimeToLiveSeconds" : 600,
          "certificate" : "\n-----BEGIN CERTIFICATE-----\nMIIDSTCCArKgAwIBAgIBADANBgkqhkiG9w0BAQQFADB8MQswCQYDVQQGEwJhdzEO\nMAwGA1UECBMFYXJ1YmExDjAMBgNVBAoTBWFydWJhMQ4wDAYDVQQHEwVhcnViYTEO\nMAwGA1UECxMFYXJ1YmExDjAMBgNVBAMTBWFydWJhMR0wGwYJKoZIhvcNAQkBFg5h\ncnViYUBhcnViYS5hcjAeFw0xNTExMjAyMjI2MjdaFw0xNjExMTkyMjI2MjdaMHwx\nCzAJBgNVBAYTAmF3MQ4wDAYDVQQIEwVhcnViYTEOMAwGA1UEChMFYXJ1YmExDjAM\nBgNVBAcTBWFydWJhMQ4wDAYDVQQLEwVhcnViYTEOMAwGA1UEAxMFYXJ1YmExHTAb\nBgkqhkiG9w0BCQEWDmFydWJhQGFydWJhLmFyMIGfMA0GCSqGSIb3DQEBAQUAA4GN\nADCBiQKBgQDHtC5gUXxBKpEqZTLkNvFwNGnNIkggNOwOQVNbpO0WVHIivig5L39W\nqS9u0hnA+O7MCA/KlrAR4bXaeVVhwfUPYBKIpaaTWFQR5cTR1UFZJL/OF9vAfpOw\nznoD66DDCnQVpbCjtDYWX+x6imxn8HCYxhMol6ZnTbSsFW6VZjFMjQIDAQABo4Ha\nMIHXMB0GA1UdDgQWBBTx0lDzjH/iOBnOSQaSEWQLx1syGDCBpwYDVR0jBIGfMIGc\ngBTx0lDzjH/iOBnOSQaSEWQLx1syGKGBgKR+MHwxCzAJBgNVBAYTAmF3MQ4wDAYD\nVQQIEwVhcnViYTEOMAwGA1UEChMFYXJ1YmExDjAMBgNVBAcTBWFydWJhMQ4wDAYD\nVQQLEwVhcnViYTEOMAwGA1UEAxMFYXJ1YmExHTAbBgkqhkiG9w0BCQEWDmFydWJh\nQGFydWJhLmFyggEAMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEEBQADgYEAYvBJ\n0HOZbbHClXmGUjGs+GS+xC1FO/am2suCSYqNB9dyMXfOWiJ1+TLJk+o/YZt8vuxC\nKdcZYgl4l/L6PxJ982SRhc83ZW2dkAZI4M0/Ud3oePe84k8jm3A7EvH5wi5hvCkK\nRpuRBwn3Ei+jCRouxTbzKPsuCVB+1sNyxMTXzf0=\n-----END CERTIFICATE-----\n                ",
          "privateKey" : "\n-----BEGIN RSA PRIVATE KEY-----\nMIICXQIBAAKBgQDHtC5gUXxBKpEqZTLkNvFwNGnNIkggNOwOQVNbpO0WVHIivig5\nL39WqS9u0hnA+O7MCA/KlrAR4bXaeVVhwfUPYBKIpaaTWFQR5cTR1UFZJL/OF9vA\nfpOwznoD66DDCnQVpbCjtDYWX+x6imxn8HCYxhMol6ZnTbSsFW6VZjFMjQIDAQAB\nAoGAVOj2Yvuigi6wJD99AO2fgF64sYCm/BKkX3dFEw0vxTPIh58kiRP554Xt5ges\n7ZCqL9QpqrChUikO4kJ+nB8Uq2AvaZHbpCEUmbip06IlgdA440o0r0CPo1mgNxGu\nlhiWRN43Lruzfh9qKPhleg2dvyFGQxy5Gk6KW/t8IS4x4r0CQQD/dceBA+Ndj3Xp\nubHfxqNz4GTOxndc/AXAowPGpge2zpgIc7f50t8OHhG6XhsfJ0wyQEEvodDhZPYX\nkKBnXNHzAkEAyCA76vAwuxqAd3MObhiebniAU3SnPf2u4fdL1EOm92dyFs1JxyyL\ngu/DsjPjx6tRtn4YAalxCzmAMXFSb1qHfwJBAM3qx3z0gGKbUEWtPHcP7BNsrnWK\nvw6By7VC8bk/ffpaP2yYspS66Le9fzbFwoDzMVVUO/dELVZyBnhqSRHoXQcCQQCe\nA2WL8S5o7Vn19rC0GVgu3ZJlUrwiZEVLQdlrticFPXaFrn3Md82ICww3jmURaKHS\nN+l4lnMda79eSp3OMmq9AkA0p79BvYsLshUJJnvbk76pCjR28PK4dV1gSDUEqQMB\nqy45ptdwJLqLJCeNoR0JUcDNIRhOCuOPND7pcMtX6hI/\n-----END RSA PRIVATE KEY-----\n                ",
          "privateKeyPassword" : "password"
        },
        "links" : {
          "logout" : {
            "redirectUrl" : "/login",
            "redirectParameterName" : "redirect",
            "disableRedirectParameter" : true,
            "whitelist" : null
          },
          "selfService" : {
            "selfServiceLinksEnabled" : true,
            "signup" : "/create_account",
            "passwd" : "/forgot_password"
          }
        },
        "prompts" : [ {
          "name" : "username",
          "type" : "text",
          "text" : "Email"
        }, {
          "name" : "password",
          "type" : "password",
          "text" : "Password"
        }, {
          "name" : "passcode",
          "type" : "password",
          "text" : "One Time Code ( Get one at http://localhost:8080/uaa/passcode )"
        } ],
        "idpDiscoveryEnabled" : false
      },
      "name" : "uaa",
      "version" : 1,
      "description" : "The system zone for backwards compatibility",
      "created" : 946684800000,
      "last_modified" : 1463595904830
    }, {
      "id" : "y2nlq2bg",
      "subdomain" : "y2nlq2bg",
      "config" : {
        "tokenPolicy" : {
          "accessTokenValidity" : -1,
          "refreshTokenValidity" : -1,
          "jwtRevocable" : false,
          "activeKeyId" : null,
          "keys" : { }
        },
        "samlConfig" : {
          "assertionSigned" : true,
          "requestSigned" : true,
          "wantAssertionSigned" : false,
          "wantAuthnRequestSigned" : false,
          "assertionTimeToLiveSeconds" : 600,
          "certificate" : null,
          "privateKey" : null,
          "privateKeyPassword" : null
        },
        "links" : {
          "logout" : {
            "redirectUrl" : "/login",
            "redirectParameterName" : "redirect",
            "disableRedirectParameter" : true,
            "whitelist" : null
          },
          "selfService" : {
            "selfServiceLinksEnabled" : true,
            "signup" : "/create_account",
            "passwd" : "/forgot_password"
          }
        },
        "prompts" : [ {
          "name" : "username",
          "type" : "text",
          "text" : "Email"
        }, {
          "name" : "password",
          "type" : "password",
          "text" : "Password"
        }, {
          "name" : "passcode",
          "type" : "password",
          "text" : "One Time Code (Get on at /passcode)"
        } ],
        "idpDiscoveryEnabled" : false
      },
      "name" : "The Twiglet Zone",
      "version" : 0,
      "description" : "Like the Twilight Zone but tastier.",
      "created" : 1463595919020,
      "last_modified" : 1463595919020
    } ]`)
	}))
	defer ts.Close()

	uaac := &uaaClient{
		serverURL:    ts.URL,
		clientID:     "clientID",
		clientSecret: "clientSecret",
	}

	zones, err := uaac.ListZones()
	if err != nil {
		t.Errorf("Failed to get Zones: %v", err)
	}

	if len(zones) == 0 {
		t.Error("[]Zone was empty")
	}
}
