package usecase

import (
	"github.com/rwirdemann/restvoice/kapitel05/domain"
)

type UpdateInvoicePort interface {
	UpdateInvoice(invoice domain.Invoice) error
	RateByProjectIdAndActivityId(projectId int, activityId int) domain.Rate
	ActivityById(id int) domain.Activity
	GetBookingsByInvoiceId(invoiceId int) []domain.Booking
}

type UpdateInvoice struct {
	repository UpdateInvoicePort
}

func NewUpdateInvoice(repository UpdateInvoicePort) UpdateInvoice {
	return UpdateInvoice{repository: repository}
}

func (u UpdateInvoice) Run(invoice domain.Invoice) error {
	if invoice.IsReadyForAggregation() {
		bookings := u.repository.GetBookingsByInvoiceId(invoice.Id)
		for _, b := range bookings {
			activity := u.repository.ActivityById(b.ActivityId)
			rate := u.repository.RateByProjectIdAndActivityId(b.ProjectId, b.ActivityId)
			invoice.AddPosition(b.ProjectId, activity.Name, b.Hours, rate.Price)
		}
		invoice.Status = "payment expected"
	}

	return u.repository.UpdateInvoice(invoice)
}
