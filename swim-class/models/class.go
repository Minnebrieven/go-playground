package models

import "time"

type Class struct {
	ID               uint
	Name             string
	TotalAppointment int
	Appointments     []Appointment
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
