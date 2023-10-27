package clinic

import (
	"CRUD-FINAL/internal/domain"
	"database/sql"
)



type SqlAppointImp struct {
	DB *sql.DB
	sqlpatient SqlPatientImp
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
	// table:= `INSERT INTO dto (id= ?, patient = ?, dentist = ?, 'date' = ?) VALUES (?,?,?,?)`
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	// stmt2, err:= s.DB.Prepare(table)
	// if err != nil {
	// 	return 0, err
	// }
	res, err := stmt.Exec(a.Patient.Id, a.Dentist.Id, a.Date)
	if err != nil {
		return 0, err
	}
	
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	// _, err2 := stmt2.Exec(
	// 	a.Id,
	// 	fmt.Sprintf("%s %s", a.Patient.Name, a.Patient.LastName),
	// 	fmt.Sprintf("%s %s", a.Dentist.Name, a.Dentist.LastName),
	// 	a.Date,
	// )

	
	// if err2 != nil {
	// 	return 0, err2
	// }

	return int(id), nil

}

// func (s *SqlDentistImp) DeleteA(id int)(int, error){
// 	query:= "DELETE FROM dentists WHERE id= ?";
// 	stmt, err := s.DB.Prepare(query)
// 	if err != nil {
// 		return 0, err
// 	}
// 	_, err2 := stmt.Exec(id)
// 	if err2 != nil {
// 		return 0, err
// 	}

// 	return int(id), nil
// }

// func (s *SqlDentistImp) UpdateA(dentist domain.Dentist)(int, error){
// 	query := "UPDATE dentists SET name = ?, last_name = ?, license = ? WHERE id = ?";
// 	stmt, err := s.DB.Prepare(query)
// 	if err != nil {
// 		return 0, err
// 	}

// 	_, err2 := stmt.Exec(dentist.Name, dentist.LastName, dentist.LicenseCode, dentist.Id)
// 	if err2 != nil {
// 		return 0, err
// 	}
// 	return dentist.Id, nil
// }

func structToDTO(ap domain.Appointment) domain.AppointmentDTO {
	apDTO := domain.AppointmentDTO{}
	apDTO.Id = ap.Id
	apDTO.PatientName = ap.Patient.Name
	apDTO.DentistName = ap.Dentist.Name
	apDTO.Date = ap.Date
	return apDTO
}