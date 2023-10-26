package clinic

import "CRUD-FINAL/internal/domain"

type SQLDentistInteface interface {
	Create(dentist *domain.Dentist) (int, error)
	Read(id int) (*domain.Dentist, error)
	ReadAll()([]domain.Dentist, error)
	Update(dentist domain.Dentist) (int, error)
	Delete(id int) (int, error)
}

type SQLPatientInteface interface {
	Create(patient *domain.Patient) (int, error)
	Read(id int) (*domain.Patient, error)
	ReadAll()([]domain.Patient, error)
	Update(patient domain.Patient) (int, error)
	Delete(id int) (int, error)
}

type SQLAppointmentInteface interface {
	Create(a *domain.Appointment) (int, error)
	Read(id int) (*domain.Appointment, error)
	ReadAll()([]domain.Appointment, error)
	Update(a domain.Appointment) (int, error)
	Delete(id int) (int, error)
}