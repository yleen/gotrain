package config

type Jwt struct {
	AccessTokenSecret      string `json:"access_token_secret" yaml:"access_token_secret"`
	RefreshTokenSecret     string `json:"refresh_token_secret" yaml:"refresh_token_secret"`
	AccessTokenExpireTime  string `json:"access_token_expire_time" yaml:"access_token_expire_time"`
	RefreshTokenExpireTime string `json:"refresh_token_expire_time" yaml:"refresh_token_expire_time"`
	Issuer                 string `json:"issuer" yaml:"issuer"`
}
