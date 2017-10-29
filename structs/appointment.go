package structs

import "time"

type WeeklyEvent struct {
	Texts      []string     `json:"texts"`
	AnnounceAt time.Weekday `json:"announce_at"` //TODO: Make sure at sunday still works
}
