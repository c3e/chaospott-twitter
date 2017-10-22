package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	"github.com/PuloV/ics-golang"
	"github.com/fronbasal/chaospott-twitter/helpers"
	"github.com/fronbasal/chaospott-twitter/structs"
)

var tweetQueue []structs.CalTweet
var calendar *ics.Calendar

func main() {
	creds := helpers.GetTwitterCredentials()
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	parser := ics.New()
	ics.FilePath = "tmp"
	ics.DeleteTempFiles = false
	ics.RepeatRuleApply = true
	url := helpers.GetConfig().Calendar
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to get calendar. Wrong URL?")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response. Broken server?")
	}
	parser.Load(string(b[:]))
	parser.Wait()
	cal, err := parser.GetCalendars()
	if err != nil {
		log.Fatal("Failed to get events: " + err.Error())
	}
	calendar = cal[0]
	events, err := calendar.GetEventsByDate(time.Now().Add(time.Hour * 24))
	if err != nil {
		log.Fatal("Could not get calendar: " + err.Error())
	}
	for _, e := range events {
		var t structs.CalTweet
		t.Text = e.GetDescription() + " am " + e.GetStart().Format("02. um 15:04") + " @ " + strings.Replace(e.GetLocation(), "\\", "", -1)
		if len(t.Text) > 140 {
			fmt.Println("Failed to add event " + e.GetDescription() + " to the queue. Text is too long.")
		}
		t.Timestamp = e.GetStart()
		tweetQueue = append(tweetQueue, t)
	}
	for _, tweet := range tweetQueue {
		fmt.Println("Tweeting: " + tweet.Text)
		t, _, err := client.Statuses.Update(tweet.Text, nil)
		if err != nil {
			log.Fatal("Failed to tweet: " + err.Error())
		}
		fmt.Println("Tweeted! https://twitter.com/" + t.User.ScreenName + "/status/" + t.IDStr)
	}
}
