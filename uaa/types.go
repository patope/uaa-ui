package uaa

import (
	"encoding/json"
)

type ServerInfo struct {
	App            map[string]string
	Links          map[string]string
	ZoneName       string `json:"zone_name"`
	EntityID       string
	CommitID       string `json:"commit_id"`
	IDPDefinitions map[string]string
	Prompts        map[string][]string
	Timestamp      string
}

func (s ServerInfo) Version() string {
	return s.App["version"]
}

type AccessToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string
	JTI       string
}

type OauthClients struct {
	Clients      []OauthClient `json:"resources"`
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Schemas      []string
}

type OauthClient struct {
	ID   string `json:"client_id"`
	Name string
	//AutoApprove          bool - TODO - this field can be a bool or an array??
	Action                 string
	Scope                  []string
	ResourceIDs            []string `json:"resource_ids"`
	Authorities            []string
	AuthorizedGrantTypes   []string `json:"authorized_grant_types"`
	LastModified           int
	RedirectURI            []string `json:"redirect_uri"`
	SignupRedirectURL      string   `json:"signup_redirect_url"`
	ChangeEmailRedirectURL string   `json:"change_email_redirect_url"`
}

type IdentityZone struct {
	ID           string
	Subdomain    string
	Name         string
	Version      int
	Description  string
	Created      int
	LastModified int `json:"last_modified"`
}

type Users struct {
	Users        []User `json:"resources"`
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Schemas      []string
}

type Groups struct {
	Groups []struct {
		GUID        string `json:"id"`
		DisplayName string `json:"displayName"`
		Members     []struct {
			Value       string `json:"value"`
			Type        string `json:"type"`
			Origin      string `json:"origin"`
			DisplayName string `json:"display"`
		}
	} `json:"resources"`
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Schemas      []string
}

type User struct {
	GUID       string `json:"id"`
	ExternalID string `json:"externalId"`
	Username   string `json:"userName"`
	Name       Name   `json:"name"`
	Groups     []struct {
		Value   string
		Display string
		Type    string
	}
	Emails               []UserEmail
	Active               bool   `json:"active"`
	Verified             bool   `json:"verified"`
	Origin               string `json:"origin"`
	ZoneID               string
	PasswordLastModified string
	Schemas              []string
}

type Name struct {
	GivenName  string
	FamilyName string
}

type UserEmail struct {
	Value   string
	Primary bool
}

type Approval struct {
	UserID        string
	ClientID      string
	Scope         string
	Status        string
	LastUpdatedAt string
	ExpiresAt     string
}

type Group struct {
	GUID        string `json:"id"`
	DisplayName string `json:"displayName"`
	ZoneID      string `json:"zoneid"`
	Members     []struct {
		Origin string `json:"origin"`
		Type   string `json:"type"`
		Value  string `json:"value"`
	} `json:"members"`
}

type IdentityProvider struct {
	GUID           string          `json:"id"`
	OriginKey      string          `json:"originKey"`
	Name           string          `json:"name"`
	Type           string          `json:"type"`
	Config         json.RawMessage `json:"config"`
	Version        int             `json:"version"`
	Created        int64           `json:"created"`
	Active         bool            `json:"active"`
	IdentityZoneID string          `json:"identityZoneId"`
	LastModified   int64           `json:"last_modified"`
}

type SamlServiceProvider struct {
	GUID     string `json:"id"`
	EntityID string `json:"entityId"`
	Name     string `json:"name"`
	Config   struct {
		MetaDataLocation         string `json:"metaDataLocation"`
		NameID                   string `json:"nameID"`
		SingleSignOnServiceIndex int32  `json:"singleSignOnServiceIndex"`
		MetadataTrustCheck       bool   `json:"metadataTrustCheck"`
		SkipSslValidation        bool   `json:"skipSslValidation"`
	} `json:"config"`
	Version        int32 `json:"version"`
	Created        int64 `json:"created"`
	LastModified   int64 `json:"lastModified"`
	Active         int64 `json:"active"`
	IdentityZoneID int64 `json:"identityZoneId"`
}
