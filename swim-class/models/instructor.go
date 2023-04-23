package models

import "time"

type Instructor struct {
	ID        uint
	Name      string
	Phone     int
	Appointments []Appointment
	CreatedAt time.Time
	UpdatedAt time.Time
}
