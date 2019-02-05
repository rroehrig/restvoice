// Code generated by mockery v1.0.0
package mocks

import domain "github.com/rwirdemann/restvoice/kapitel05/domain"
import mock "github.com/stretchr/testify/mock"

// UpdateInvoicePort is an autogenerated mock type for the UpdateInvoicePort type
type UpdateInvoicePort struct {
	mock.Mock
}

// ActivityById provides a mock function with given fields: id
func (_m *UpdateInvoicePort) ActivityById(id int) domain.Activity {
	ret := _m.Called(id)

	var r0 domain.Activity
	if rf, ok := ret.Get(0).(func(int) domain.Activity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Activity)
	}

	return r0
}

// GetBookingsByInvoiceId provides a mock function with given fields: invoiceId
func (_m *UpdateInvoicePort) GetBookingsByInvoiceId(invoiceId int) []domain.Booking {
	ret := _m.Called(invoiceId)

	var r0 []domain.Booking
	if rf, ok := ret.Get(0).(func(int) []domain.Booking); ok {
		r0 = rf(invoiceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Booking)
		}
	}

	return r0
}

// RateByProjectIdAndActivityId provides a mock function with given fields: projectId, activityId
func (_m *UpdateInvoicePort) RateByProjectIdAndActivityId(projectId int, activityId int) domain.Rate {
	ret := _m.Called(projectId, activityId)

	var r0 domain.Rate
	if rf, ok := ret.Get(0).(func(int, int) domain.Rate); ok {
		r0 = rf(projectId, activityId)
	} else {
		r0 = ret.Get(0).(domain.Rate)
	}

	return r0
}

// UpdateInvoice provides a mock function with given fields: invoice
func (_m *UpdateInvoicePort) UpdateInvoice(invoice domain.Invoice) error {
	ret := _m.Called(invoice)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Invoice) error); ok {
		r0 = rf(invoice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
