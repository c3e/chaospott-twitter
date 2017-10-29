package helpers

import (
	"io/ioutil"
	"log"
)


// ReadJsonFile returns a file
func ReadJsonFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read config")
	}
	return b
}
