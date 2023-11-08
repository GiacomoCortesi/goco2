package goco2

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// CO2Saving models the CO2 emissions saving per day, week, month, and year
type CO2Saving struct {
	Day   uint64 `json:"day"`
	Week  uint64 `json:"week"`
	Month uint64 `json:"month"`
	Year  uint64 `json:"year"`
}

// Intervention models a house efficiency upgrade intervention
type Intervention struct {
	ID   uuid.UUID `json:"id"`
	Date time.Time `json:"date"`
}

type Interventions map[string]Intervention

// Since returns the number of energy efficiency upgrade interventions since the specified time
func (i Interventions) Since(since time.Time) uint64 {
	var intNumber uint64
	for _, intervention := range i {
		if intervention.Date.After(since) {
			intNumber++
		}
	}
	return intNumber
}

func (i Interventions) Today() uint64 {
	return i.Since(time.Now().AddDate(0, 0, -1))
}

func (i Interventions) ThisWeek() uint64 {
	return i.Since(time.Now().AddDate(0, 0, -7))
}

func (i Interventions) ThisMonth() uint64 {
	return i.Since(time.Now().AddDate(0, -1, 0))
}

func (i Interventions) ThisYear() uint64 {
	return i.Since(time.Now().AddDate(-1, 0, 0))
}

var ErrInterventionNotFound = fmt.Errorf("cannot find intervention information")
