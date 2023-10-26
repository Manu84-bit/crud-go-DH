package clinic

import "CRUD-FINAL/internal/domain"

type SQLDentistInteface interface {
	Create(dentist *domain.Dentist) (int, error)
	Read(id int) (*domain.Dentist, error)
	ReadAll()([]domain.Dentist, error)
	Update(dentist domain.Dentist) (int, error)
	Delete(id int) (int, error)
}

type SQLPatientInterface interface {
	CreateP(patient *domain.Patient) (int, error)
	ReadP(id int) (*domain.Patient, error)
	ReadAllP()([]domain.Patient, error)
	// UpdateP(patient domain.Patient) (int, error)
	// DeleteP(id int) (int, error)
}

type SQLAppointmentInteface interface {
	CreateA(a *domain.Appointment) (int, error)
	ReadA(id int) (*domain.Appointment, error)
	ReadAllA()([]domain.Appointment, error)
	UpdateA(a domain.Appointment) (int, error)
	DeleteA(id int) (int, error)
}