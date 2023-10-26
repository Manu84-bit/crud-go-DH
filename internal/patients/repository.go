package patients

import (
	"CRUD-FINAL/internal/domain"
	"CRUD-FINAL/pkg/clinic"
	"CRUD-FINAL/pkg/web"
	"fmt"
)

type IRepository interface {
	Save(patient *domain.Patient) (int, error)
	GetById(id int) (*domain.Patient, error)
	GetAll()([]domain.Patient, error)
	// Update(d domain.Patient) (int, error)
	// DeleteById(id int) (int, error)
}

type RepositoryImpl struct {
	PatientRepo clinic.SQLPatientInterface
}

func (r *RepositoryImpl) Save(patient *domain.Patient) (int, error){
	newPatient, err := r.PatientRepo.CreateP(patient)
	if err != nil{
		return 0, web.NewBadRequestApiError("Wrong data")
	}

	return newPatient, nil
}

func (r *RepositoryImpl) GetAll()([]domain.Patient, error){
	patients, err := r.PatientRepo.ReadAllP()
		if err != nil {
		return nil, web.NewNotFoundException("dentists not found")
	}

	return patients, nil
}

func (r *RepositoryImpl) GetById(id int) (*domain.Patient, error){
	patient, err :=	r.PatientRepo.ReadP(id)
	if err != nil {
		return nil, web.NewNotFoundException(fmt.Sprintf("patient id %d not found", id))
	}

	return patient, nil
}

// func (r *RepositoryImpl) DeleteById(id int) (int, error) {
//   idDentist, err := r.DentistRepo.Delete(id)
//   if err != nil {
// 	return 0, web.NewNotFoundException(fmt.Sprintf("dentist id %d not exists", idDentist))
//   }
//   return idDentist, nil
// }

// func (r *RepositoryImpl) Update(d domain.Dentist) (int, error){
// idDentist, err := r.DentistRepo.Update(d)
// if err != nil {
// 	return 0, web.NewNotFoundException(fmt.Sprintf("dentist id %d not exists", idDentist))
//   }

//   return idDentist, nil
// }
