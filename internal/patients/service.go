package patients

import "CRUD-FINAL/internal/domain"

type IService interface {

	GetAllPatients()([]domain.Patient, error)
	// GetByIdentifier(id int)(*domain.Patient, error)
	// SavePatientt(d *domain.Patient)(int, error)
	// DeletePatient(id int)(int, error)
	// UpdatePatient(d domain.Patient)(int, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetAllPatients()([]domain.Patient, error){
	patients, err := s.Repository.GetAll()
		if err != nil{
		return nil, err
	}

	return patients, nil
}
// func (s *Service) GetByIdentifier(id int)(*domain.Dentist, error){
// 	d, err := s.Repository.GetById(id)
// 	if err != nil{
// 		return nil, err
// 	}

// 	return d, nil
// }

// func (s *Service) DeleteDentist(id int)(int, error){
// 	id,err := s.Repository.DeleteById(id);
// 		if err != nil{
// 		return 0, err
// 	}
// 	return id, nil
// }

// func (s *Service) SaveDentist(d *domain.Dentist)(int, error){
//   id, err := s.Repository.Save(d)
//   if err != nil{
// 		return 0, err
// 	}
// 	return id, nil
// }

// func (s *Service) UpdateDentist(d domain.Dentist)(int, error){
// 	id, err := s.Repository.Update(d)
// 	 if err != nil{
// 		return 0, err
// 	}
// 	return id, nil
// }
