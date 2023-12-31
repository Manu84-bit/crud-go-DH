package clinic

import (
	"CRUD-FINAL/internal/domain"
	"database/sql"
)



type SqlPatientImp struct {
	DB *sql.DB
}

func (s *SqlPatientImp) ReadP(id int)(*domain.Patient, error) {
	var dentist domain.Patient
	query:= "SELECT * FROM patients WHERE id = ?"
	row := s.DB.QueryRow(query, id)
	err:= row.Scan(
		&dentist.Id, 
		&dentist.Name,
		&dentist.LastName,
		&dentist.Address,
		&dentist.DNI,
		&dentist.DischargeDate,
		)
	if err != nil {
		return nil, err
	}

	return &dentist, err

}
	func (s *SqlPatientImp) ReadAllP()([]domain.Patient, error) {
	query := "SELECT * FROM patients"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var patients []domain.Patient

	for rows.Next() {
		p := domain.Patient{}
		_ = rows.Scan(&p.Id, &p.Name, &p.LastName, &p.Address, &p.DNI, &p.DischargeDate)

		patients = append(patients, p)
	}

	return patients, nil
}

func(s *SqlPatientImp) CreateP(patient *domain.Patient) (int, error) {
	query:= "INSERT INTO patients(name, last_name, address, dni, discharge_date) VALUES (?,?,?,?,?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(patient.Name, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SqlPatientImp) DeleteP(id int)(int, error){
	query:= "DELETE FROM patients WHERE id= ?";
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SqlPatientImp) UpdateP(patient domain.Patient)(int, error){
	query := "UPDATE patients SET name = ?, last_name = ?, address = ?, dni = ?, discharge_date = ? WHERE id = ?";
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	_, err2 := stmt.Exec(patient.Name, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate, patient.Id)
	if err2 != nil {
		return 0, err
	}
	return patient.Id, nil
}