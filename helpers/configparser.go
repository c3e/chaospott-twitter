package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fronbasal/chaospott-twitter/structs"
)

// GetTwitterCredentials returns a struct with the twitter API keys
func GetTwitterCredentials() structs.TwitterCredentials {
	var v structs.TwitterCredentials
	json.Unmarshal(read("credentials.json"), &v)
	return v
}

// GetConfig returns a struct with the configuration file
func GetConfig() structs.Config {
	var v structs.Config
	json.Unmarshal(read("config.json"), &v)
	return v
}

func read(filename string) []byte {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Failed to read config")
	}
	return b
}
