package structs

import (
	"time"
)

// CalTweet for the tweet queue
type CalTweet struct {
	Text      string
	Timestamp time.Time
}
