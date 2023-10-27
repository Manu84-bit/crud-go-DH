package main

import (
	"CRUD-FINAL/cmd/server/handlerappointment"
	"CRUD-FINAL/cmd/server/handlerdentist"
	"CRUD-FINAL/cmd/server/handlerpatient"
	"CRUD-FINAL/internal/appointments"
	"CRUD-FINAL/internal/dentists"
	"CRUD-FINAL/internal/patients"
	"CRUD-FINAL/pkg/clinic"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {


	// dsn1 := "root:Servidor2$@tcp(localhost:3306)/?parseTime=true"

	//Base de datos mysql definida en el docker-compose.yaml
	dsn2 := "user:user@tcp(localhost:8083)/"

    // Read SQL script from file
    sqlScript, err := os.ReadFile("./init/init.sql")
    if err != nil {
        log.Fatal(err)
    }

    // Split the script into individual SQL statements
    sqlStatements := strings.Split(string(sqlScript), ";")

    // Open a database connection
    db, err := sql.Open("mysql", dsn2)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Execute each SQL statement
    for _, statement := range sqlStatements {
        statement = strings.TrimSpace(statement)
        if statement != "" {
            _, err = db.Exec(statement)
            if err != nil {
                log.Fatalf("Error executing SQL statement: %s\nError: %v", statement, err)
            }
        }
    }

    fmt.Println("SQL script executed successfully.")


	errPing := db.Ping()
	if errPing != nil{
		panic(errPing)
	}


	//API 
	r := gin.Default()

	clinicD := clinic.SqlDentistImp{DB: db}
	repoD := dentists.RepositoryImpl{DentistRepo: &clinicD}
	serviceD := dentists.Service{Repository: &repoD}
	handlerD := handlerdentist.DentistHandler{DentistService: &serviceD}

	apiDentist := r.Group("/dentists")
	apiDentist.GET("", handlerD.GetAll)
	apiDentist.GET("/:id", handlerD.GetById)
	apiDentist.POST("/new", handlerD.SaveDentist)
	apiDentist.DELETE("/delete/:id", handlerD.DeleteById)
	apiDentist.PUT("/update/:id", handlerD.UpdateDentist)

    clinicP := clinic.SqlPatientImp{DB: db}
	repoP := patients.RepositoryImpl{PatientRepo: &clinicP}
	serviceP := patients.Service{Repository: &repoP}
	handlerP := handlerpatient.PatientHandler{PatientService: &serviceP}

	apiPatient := r.Group("/patients")
		apiPatient.GET("", handlerP.GetAll)
		apiPatient.GET("/:id", handlerP.GetById)
		apiPatient.POST("/new", handlerP.SavePatient)
		apiPatient.DELETE("/delete/:id", handlerP.DeleteById)
	    apiPatient.PUT("/update/:id", handlerP.UpdatePatient)


	clinicA := clinic.SqlAppointImp{DB: db}
	repoA := appointments.RepositoryImpl{AppointRepo: &clinicA}
	serviceA := appointments.Service{Repository: &repoA}
	handlerA := handlerappointment.AppointHandler{ApointmentService: &serviceA}

	apiAppointment := r.Group("/appointments")
		apiAppointment.GET("", handlerA.GetAll)
		apiAppointment.GET("/:id", handlerA.GetById)
		apiAppointment.POST("/new", handlerA.SaveAppointment)
	// 	apiAppointment.DELETE("/delete/:id", handlerA.DeleteById)
	// 	apiAppointment.PUT("/update/:id", handlerA.UpdateAppointment)

	r.Run(":3000")
}