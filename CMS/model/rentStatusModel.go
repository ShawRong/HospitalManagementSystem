package model

import (
	"time"
)

type RentStatus struct {
	ID        int       `json:"ID"`
	UserID    int       `json:"userID"`
	CarID     int       `json:"carID"`
	Fine      int       `json:"fine"`
	Cost      int       `json:"cost"`
	Date      time.Time `json:"date"`
	ChargerID int       `json:"chargerID"`
}
