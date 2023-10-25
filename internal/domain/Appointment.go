package domain

import (
	"time"
)

type Appointment struct {
	Id            int     `json:"id"`
	Patient       Patient `json:"patient" binding:"required"`
	Dentist       Dentist `json:"dentist" binding:"required"`
	Date 		  time.Time `json:"date" binding:"required"`
}