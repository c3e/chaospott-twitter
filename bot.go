package main

import (
	"github.com/dghubble/oauth1"
	"github.com/fronbasal/chaospott-twitter/helpers"
	"github.com/dghubble/go-twitter/twitter"
	"encoding/json"
	"github.com/fronbasal/chaospott-twitter/structs"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().Unix())
	var creds structs.TwitterCredentials
	json.Unmarshal(helpers.ReadJsonFile("config/credentials.json"), &creds)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	var recurring []structs.WeeklyEvent
	json.Unmarshal(helpers.ReadJsonFile("config/recurring.json"), &recurring)
	for _, event := range recurring {
		if event.AnnounceAt == time.Now().Weekday() {
			client.Statuses.Update(event.Texts[rand.Intn(len(event.Texts)-1)], nil)
		}
	}
}
