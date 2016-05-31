package uaac

type AccessToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
	JTI       string `json:"jti"`
}

type Zone struct {
	ID           string `json:"id"`
	Subdomain    string `json:"subdomain"`
	Name         string `json:"name"`
	Version      int    `json:"version"`
	Description  string `json:"description"`
	Created      int    `json:"created"`
	LastModified int    `json:"last_modified"`
}
