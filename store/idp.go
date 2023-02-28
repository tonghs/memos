package store

type IdentityProviderType string

const (
	IdentityProviderOAuth2 IdentityProviderType = "OAUTH2"
)

type IdentityProviderConfig struct {
	OAuth2Config *IdentityProviderOAuth2Config
}

type IdentityProviderOAuth2Config struct {
	ClientID     string        `json:"clientId"`
	ClientSecret string        `json:"clientSecret"`
	AuthURL      string        `json:"authUrl"`
	TokenURL     string        `json:"tokenUrl"`
	UserInfoURL  string        `json:"userInfoUrl"`
	Scopes       []string      `json:"scopes"`
	FieldMapping *FieldMapping `json:"fieldMapping"`
}

type FieldMapping struct {
	Identifier  string `json:"identifier"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type IdentityProviderMessage struct {
	ID               int
	Name             string
	Type             IdentityProviderType
	IdentifierFilter string
	Config           *IdentityProviderConfig
}

type FindIdentityProviderMessage struct {
	ID *int
}

type UpdateIdentityProviderMessage struct {
	ID               int
	Type             IdentityProviderType
	Name             *string
	IdentifierFilter *string
	Config           *IdentityProviderConfig
}

type DeleteIdentityProviderMessage struct {
	ID int
}
