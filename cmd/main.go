package main

import (
	"CRUD-FINAL/cmd/server/handlerdentist"
	"CRUD-FINAL/cmd/server/handlerpatient"
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


	dsn1 := "root:Servidor2$@tcp(localhost:3306)/?parseTime=true"
	// dsn2 := "user:user@tcp(mysql-db:33060)/go_db"

    // Read SQL script from file
    sqlScript, err := os.ReadFile("./init/init.sql")
    if err != nil {
        log.Fatal(err)
    }

    // Split the script into individual SQL statements
    sqlStatements := strings.Split(string(sqlScript), ";")

    // Open a database connection
    db, err := sql.Open("mysql", dsn1)
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
	apiDentist.GET("/:id", handlerD.GetById)
	apiDentist.GET("", handlerD.GetAll)
	apiDentist.DELETE("/delete/:id", handlerD.DeleteById)
	apiDentist.POST("/new", handlerD.SaveDentist)
	apiDentist.PUT("/update/:id", handlerD.UpdateDentist)

    clinicP := clinic.SqlPatientImp{DB: db}
	repoP := patients.RepositoryImpl{PatientRepo: &clinicP}
	serviceP := patients.Service{Repository: &repoP}
	handlerP := handlerpatient.PatientHandler{PatientService: &serviceP}

	apiPatient := r.Group("/patients")
		// Define routes within the group
		apiPatient.GET("", handlerP.GetAll)
		// apiPatient.POST("/endpoint2", endpoint2Handler)



	r.Run(":3000")
}