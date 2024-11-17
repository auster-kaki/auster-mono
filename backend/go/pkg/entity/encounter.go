package entity

import "time"

type EncounterID string

type Encounter struct {
	ID          EncounterID
	Name        string
	Place       string
	UserID      UserID
	Date        time.Time
	Description string
}

type Encounters []*Encounter
