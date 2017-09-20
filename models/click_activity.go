package models

import (
	"encoding/json"
	"time"
)

type ClickActivity struct {
	Count int       `json:"count" db:"count"`
	Date  time.Time `json:"date" db:"date"`
}

func (c ClickActivity) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"count": c.Count,
		"date":  c.Date.Format("2006-01-02"),
	}

	return json.Marshal(m)
}

type ClickActivities []ClickActivity
