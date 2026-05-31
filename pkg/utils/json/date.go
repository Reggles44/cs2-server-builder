package json

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

const format = `"2006-01-02T15:04:05Z"`

func (t *DateTime) UnmarshalJSON(b []byte) error {
	datetime, err := time.Parse(format, string(b))
	if err != nil {
		return err
	}

	t.Time = datetime
	return nil
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format(format))
}
