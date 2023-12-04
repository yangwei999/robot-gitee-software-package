package useradapterimpl

type Config struct {
	AppId         string `json:"app_id"                 required:"true"`
	AppSecret     string `json:"app_secret"             required:"true"`
	TokenEndpoint string `json:"token_endpoint"         required:"true"`
	UserEndpoint  string `json:"user_endpoint"          required:"true"`
}
