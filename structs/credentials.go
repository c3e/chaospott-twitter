package structs

// TwitterCredentials struct for twitter creds
type TwitterCredentials struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

// Config struct for general configuration
type Config struct {
	Calendar string `json:"calendar"`
}
