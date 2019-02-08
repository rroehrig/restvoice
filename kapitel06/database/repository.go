package database

import (
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rwirdemann/restvoice/kapitel05/domain"
)

type FakeRepository struct {
	nextId     int
	invoices   map[int]domain.Invoice
	bookings   map[int]map[int]domain.Booking
	projects   map[int]domain.Project
	customers  map[int]domain.Customer
	activities map[string]map[int]domain.Activity
	rates      map[int]map[int]domain.Rate
}

func (r *FakeRepository) GetActivities(userId string) []domain.Activity {
	var activities []domain.Activity
	for _, a := range r.activities[userId] {
		activities = append(activities, a)
	}
	return activities
}

func NewFakeRepository() *FakeRepository {
	r := FakeRepository{}
	r.invoices = make(map[int]domain.Invoice)
	r.bookings = make(map[int]map[int]domain.Booking)
	r.projects = make(map[int]domain.Project)
	r.customers = make(map[int]domain.Customer)
	r.rates = make(map[int]map[int]domain.Rate)
	r.activities = make(map[string]map[int]domain.Activity)
	return &r
}

func (r *FakeRepository) GetBookingsByInvoiceId(id int) []domain.Booking {
	var bookings []domain.Booking
	for _, b := range r.bookings[id] {
		bookings = append(bookings, b)
	}
	return bookings
}

func (r *FakeRepository) GetInvoice(id int, join ...string) domain.Invoice {
	i := r.invoices[id]
	if len(join) > 0 {
		if strings.Contains(join[0], "bookings") {
			i.Bookings = r.GetBookingsByInvoiceId(id)
		}
	}
	return i
}

func (r *FakeRepository) GetProject(id int) domain.Project {
	return r.projects[id]
}

func (r *FakeRepository) GetCustomer(id int) domain.Customer {
	return r.customers[id]
}

func (r *FakeRepository) CreateInvoice(invoice domain.Invoice) (domain.Invoice, error) {
	if invoice.Id == 0 {
		invoice.Id = r.getNextId()
	}
	if invoice.Status == "" {
		invoice.Status = "open"
	}
	r.invoices[invoice.Id] = invoice
	return invoice, nil
}

func (r *FakeRepository) UpdateInvoice(invoice domain.Invoice) error {
	r.invoices[invoice.Id] = invoice
	return nil
}

func (r *FakeRepository) CreateBooking(booking domain.Booking) (domain.Booking, error) {
	booking.Id = r.getNextId()
	if bookings, ok := r.bookings[booking.InvoiceId]; ok {
		bookings[booking.Id] = booking
	} else {
		bookings := make(map[int]domain.Booking)
		bookings[booking.Id] = booking
		r.bookings[booking.InvoiceId] = bookings
	}
	return booking, nil
}

func (r *FakeRepository) CreateActivity(activity domain.Activity) {
	activity.Id = r.getNextId()
	activity.Updated = time.Now().UTC()
	if activities, ok := r.activities[activity.UserId]; ok {
		activities[activity.Id] = activity
	} else {
		activities := make(map[int]domain.Activity)
		activities[activity.Id] = activity
		r.activities[activity.UserId] = activities
	}
}

func (r *FakeRepository) ActivityById(id int) domain.Activity {
	return r.activities[""][id]
}

func (r *FakeRepository) RateByProjectIdAndActivityId(projectId int, activityId int) domain.Rate {
	return r.rates[projectId][activityId]
}

func (r *FakeRepository) getNextId() int {
	r.nextId = r.nextId + 1
	return r.nextId
}

func (r *FakeRepository) CreateRate(rate domain.Rate) {
	if projectRates, ok := r.rates[rate.ProjectId]; ok {
		projectRates[rate.ActivityId] = rate
	} else {
		r.rates[rate.ProjectId] = make(map[int]domain.Rate)
		r.rates[rate.ProjectId][rate.ActivityId] = rate
	}
}

func (r *FakeRepository) CreateProject(p domain.Project) {
	p.Id = r.nextProjectId()
	r.projects[p.Id] = p
}

func (r *FakeRepository) nextProjectId() int {
	nextId := 1
	for _, i := range r.projects {
		if i.Id >= nextId {
			nextId = i.Id + 1
		}
	}
	return nextId
}