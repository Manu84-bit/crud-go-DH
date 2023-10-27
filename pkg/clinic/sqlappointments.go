package clinic

import (
	"CRUD-FINAL/internal/domain"
	"database/sql"
)



type SqlAppointImp struct {
	DB *sql.DB
	
}

func (s *SqlAppointImp) ReadA(id int)(*domain.AppointmentDTO, error) {
	var ap domain.AppointmentDTO
	// query:= "SELECT * FROM dto WHERE id = ?"
	query1:= `SELECT appointments.id, concat (patients.name,' ' ,patients.last_name) as patient, concat (dentists.name, ' ', dentists.last_name) as dentist, appointments.date
			FROM appointments
			INNER JOIN patients ON appointments.patient_id=patients.id
			INNER JOIN dentists ON appointments.dentist_id=dentists.id WHERE appointments.id = ?`
	row := s.DB.QueryRow(query1, id)

	err:= row.Scan(
		&ap.Id, 
		&ap.PatientName,
		&ap.DentistName,
		&ap.Date,
		)
	if err != nil {
		return nil, err
	}
	return &ap, err
}

func (s *SqlAppointImp) ReadAllA()([]domain.AppointmentDTO, error) {
	// query := "SELECT * FROM appointments"
	query1:= `SELECT appointments.id, concat (patients.name,' ' ,patients.last_name) as patient, concat (dentists.name, ' ', dentists.last_name) as dentist, appointments.date as 'date'
			FROM appointments
			INNER JOIN patients ON appointments.patient_id=patients.id
			INNER JOIN dentists ON appointments.dentist_id=dentists.id`

	rows, err := s.DB.Query(query1)
	if err != nil {
		return nil, err
	}

	var appointments []domain.AppointmentDTO

	for rows.Next() {
		a := domain.AppointmentDTO{}
		_ = rows.Scan(&a.Id, &a.PatientName, &a.DentistName, a.Date)

		appointments = append(appointments, a)
	}

	return appointments, nil
}

func(s *SqlAppointImp) CreateA(a *domain.Appointment) (int, error) {
	query:= "INSERT INTO appointments(patient_id, dentist_id, date) VALUES (?,?,?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	
	res, err := stmt.Exec(a.Patient.Id, a.Dentist.Id, a.Date)
	if err != nil {
		return 0, err
	}
	
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return int(id), nil

}

func (s *SqlAppointImp) DeleteA(id int)(int, error){
	query:= "DELETE FROM appointments WHERE id= ?";
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

func (s *SqlAppointImp) UpdateA(a domain.Appointment)(int, error){
	query := "UPDATE appointments SET patient_id = ?, dentist_id = ?, date = ? WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	_, err2 := stmt.Exec(a.Patient.Id, a.Dentist.Id, a.Date, a.Id)
	if err2 != nil {
		return 0, err
	}
	return a.Id, nil
}

