package appointments

import (
	"CRUD-FINAL/internal/domain"
	"CRUD-FINAL/pkg/clinic"
	"CRUD-FINAL/pkg/web"
	"fmt"
)

type IRepository interface {
	GetByIdentifier(id int)(*domain.AppointmentDTO, error)
	GetAllAppointments()([]domain.AppointmentDTO, error)
	SaveAppointment(d *domain.Appointment)(int, error)
	DeleteById(id int)(int, error)
	UpdateAppointment(a domain.Appointment)(int, error)
}

type RepositoryImpl struct {
	AppointRepo clinic.SQLAppointmentInterface
}

func (r *RepositoryImpl) GetByIdentifier(id int) (*domain.AppointmentDTO, error){
	appointment, err :=	r.AppointRepo.ReadA(id)
	if err != nil {
		return nil, web.NewNotFoundException(fmt.Sprintf("Appointment id %d not found", id))
	}

	return appointment, nil
}

func (r *RepositoryImpl) GetAllAppointments()([]domain.AppointmentDTO, error){
	appointments, err := r.AppointRepo.ReadAllA()
		if err != nil {
		return nil, web.NewNotFoundException("appointments not found")
	}

	return appointments, nil
}

func (r *RepositoryImpl) SaveAppointment(a *domain.Appointment) (int, error){
	newDentist, err := r.AppointRepo.CreateA(a)
	if err != nil{
		return 0, web.NewBadRequestApiError("Wrong data")
	}

	return newDentist, nil
}

func (r *RepositoryImpl) DeleteById(id int) (int, error) {
  idA, err := r.AppointRepo.DeleteA(id)
  if err != nil {
	return 0, web.NewNotFoundException(fmt.Sprintf("appointment id %d not exists", idA))
  }
  return idA, nil
}

func (r *RepositoryImpl) UpdateAppointment(a domain.Appointment) (int, error){
idA, err := r.AppointRepo.UpdateA(a)
if err != nil {
	return 0, web.NewNotFoundException(fmt.Sprintf("appointment id %d not exists", idA))
  }

  return idA, nil
}
