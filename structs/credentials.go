package structs

// TwitterCredentials struct for twitter creds
type TwitterCredentials struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// Config struct for general configuration
type Config struct {
	Calendar string `json:"calendar"`
}
