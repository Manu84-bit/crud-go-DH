package clinic

import (
	"CRUD-FINAL/internal/domain"
	"database/sql"
)



type SqlAppointImp struct {
	DB *sql.DB
}

func (s *SqlDentistImp) ReadA(id int)(*domain.Patient, error) {
	var dentist domain.Patient
	query:= "SELECT * FROM dentists WHERE id = ?"
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
	func (s *SqlDentistImp) ReadAllA()([]domain.Dentist, error) {
	query := "SELECT * FROM dentists"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var dentists []domain.Dentist

	for rows.Next() {
		d := domain.Dentist{}
		_ = rows.Scan(&d.Id, &d.Name, &d.LastName, &d.LicenseCode)
		dentists = append(dentists, d)
	}

	return dentists, nil
}

func(s *SqlDentistImp) CreateA(dentist *domain.Dentist) (int, error) {
	query:= "INSERT INTO dentists(name, last_name, license) VALUES (?,?,?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(dentist.Name, dentist.LastName, dentist.LicenseCode)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SqlDentistImp) DeleteA(id int)(int, error){
	query:= "DELETE FROM dentists WHERE id= ?";
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

func (s *SqlDentistImp) UpdateA(dentist domain.Dentist)(int, error){
	query := "UPDATE dentists SET name = ?, last_name = ?, license = ? WHERE id = ?";
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	_, err2 := stmt.Exec(dentist.Name, dentist.LastName, dentist.LicenseCode, dentist.Id)
	if err2 != nil {
		return 0, err
	}
	return dentist.Id, nil
}