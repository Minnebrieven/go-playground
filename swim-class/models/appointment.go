package models

import "time"

type Appointment struct {
	ID              uint
	ClassID         uint
	UserID          uint
	InstructorID    uint
	AppointmentDate time.Time
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
