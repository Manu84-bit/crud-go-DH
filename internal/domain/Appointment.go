package domain

type Appointment struct {
	Id      int     `json:"id"`
	Patient Patient `json:"patient" binding:"required"`
	Dentist Dentist `json:"dentist" binding:"required"`
	Date    string  `json:"date" binding:"required"`
}

type AppointmentDTO struct {
	Id          int    `json:"id"`
	PatientName string `json:"patient_name" binding:"required"`
	DentistName string `json:"dentist_name" binding:"required"`
	Date        string `json:"date" binding:"required"`
}
