package model

import "time"

type MONTH_ID int

type Month struct {
	ID      MONTH_ID `json:"id"`
	Display string   `json:"display"`
}

func (m *Month) IsCurrentMonth() bool {
	if m.Display == time.Now().Format("2006-01") {
		return true
	}
	return false
}
