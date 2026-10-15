package event

import "time"

type PostCreated struct {
	PostID     int
	Caption    string
	OccurredAt time.Time
}
