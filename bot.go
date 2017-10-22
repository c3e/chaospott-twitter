package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuloV/ics-golang"
	"github.com/fronbasal/chaospott-twitter/helpers"
)

func main() {
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
	cal, _ := parser.GetCalendars()
	for _, e := range cal[0].GetEvents() {
		fmt.Println(e.GetDescription() + " am " + e.GetStart().Format("02. um 15:04") + " @ " + strings.Replace(e.GetLocation(), "\\", "", -1))
	}
}
