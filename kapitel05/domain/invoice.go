package domain

import (
	"io/ioutil"
	"time"
)

type Position struct {
	Hours float32
	Price float32
}

type Invoice struct {
	ID         int                         `json:"id"`
	Month      int                         `json:"month"`
	Year       int                         `json:"year"`
	Status     string                      `json:"status"`
	CustomerID int                         `json:"customerId"`
	Positions  map[int]map[string]Position `json:"positions,omitempty"`
	Bookings   []Booking                   `json:"-"`
	Updated    time.Time                   `json:"updated,omitempty"`
}

func (invoice *Invoice) AddBooking() {

}

func (invoice *Invoice) AddPosition(projectID int, activity string, hours float32,
	rate float32) {
	if invoice.Positions == nil {
		invoice.Positions = make(map[int]map[string]Position)
	}

	if invoice.Positions[projectID] == nil {
		invoice.Positions[projectID] = make(map[string]Position)
	}

	if p, ok := invoice.Positions[projectID][activity]; ok {
		p.Hours = p.Hours + hours
		p.Price = p.Price + hours*rate
		invoice.Positions[projectID][activity] = p
	} else {
		position := Position{Hours: hours, Price: hours * rate}
		invoice.Positions[projectID][activity] = position
	}
}

func (invoice *Invoice) ToPDF() []byte {
	dat, _ := ioutil.ReadFile("/tmp/invoice.pdf")
	return dat
}

func (invoice Invoice) IsReadyForAggregation() bool {
	return invoice.Status == "ready for aggregation"
}

type Operation string

func (invoice Invoice) GetOperations() []Operation {
	switch invoice.Status {
	case "open":
		return []Operation{"book", "charge", "cancel", "bookings"}
	case "payment expected":
		return []Operation{"payment", "bookings"}
	case "payed":
		return []Operation{"archive"}
	case "archived":
		return []Operation{"revoke"}
	case "revoked":
		return []Operation{"archive"}
	default:
		return []Operation{}
	}
}
