package block

import "time"

type timestamp struct {
	created time.Time
}

func NewTimestamp() *timestamp {
	return &timestamp{
		created: time.Now().UTC(),
	}
}

func (t *timestamp) Created() time.Time {
	return t.created
}
