package model

import (
	"time"
)

type FixLog struct {
	ID    int       `json:"ID"`
	CarID int       `json:"carID"`
	Fee   int       `json:"fee"`
	Time  time.Time `josn:"time"`
}
