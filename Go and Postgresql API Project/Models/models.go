package models

import (
	"database/sql"
	"log"
	"time"
)

//Make this struct according to your postgres db table
type PostgresStruct struct {
	Employee_id   int
	First_name    string
	Last_name     string
	Email         string
	Phone_number  int64
	Hire_date     string
	Job_id        int
	Salary        float64
	Department_id int
}

var emp PostgresStruct

//Function for formate the date in YYYY-MM-DD format
func formatDate(dateStr string) string {

	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Fatal(err)
	}
	return t.Format("2006-01-02")
}

func SeeAllData(DB *sql.DB) ([]PostgresStruct, error) {
	query := `SELECT employee_id, first_name, last_name, email, phone_number, hire_date, job_id, salary, department_id FROM employees`

	//DB.Query is used for all the rows in the table
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var employees []PostgresStruct

	for rows.Next() {
		var emp PostgresStruct

		//scan through each rows 
		err := rows.Scan(&emp.Employee_id, &emp.First_name, &emp.Last_name, &emp.Email, &emp.Phone_number, &emp.Hire_date, &emp.Job_id, &emp.Salary, &emp.Department_id)
		if err != nil {
			log.Fatal(err)
		}

		//append all the scanned emp in empty slice employees 
		employees = append(employees, emp)
	}

	if err != nil {
		log.Fatal(err)
	}
	return employees, nil
}

func CreateEmployee(DB *sql.DB, emp PostgresStruct) {

	//sql query for inserting the data 
	query := `INSERT INTO employees(first_name, last_name, email, phone_number, hire_date, job_id, salary, department_id)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8);`

	//check if the hire_date is not empty
	if emp.Hire_date == "" {
		log.Println("Warning hire_date is not empty")
	}

	formattedDate := formatDate(emp.Hire_date)

	//execute the query in emp
	_, err := DB.Exec(query, emp.First_name, emp.Last_name, emp.Email, emp.Phone_number, formattedDate, emp.Job_id, emp.Salary, emp.Department_id)

	if err != nil {
		log.Fatal(err)
	}
}

func ReadEmployees(DB *sql.DB, id int) PostgresStruct {

	//sql query for reading of fetching one row
	query := `SELECT first_name, last_name, email, phone_number, hire_date, job_id, salary, department_id FROM employees WHERE employee_id=$1;`

	//DB.QueryRow is used to scan only one row at a time
	err := DB.QueryRow(query, id).Scan(&emp.First_name, &emp.Last_name, &emp.Email, &emp.Phone_number, &emp.Hire_date, &emp.Job_id, &emp.Salary, &emp.Department_id)

	if err != nil {
		log.Fatal(err)
	}

	return emp
}

func UpdateEmployee(DB *sql.DB, emp PostgresStruct) {

	// sql query to update the employees table
	query := `UPDATE employees 
	SET first_name=$1, last_name=$2, email=$3, phone_number=$4, hire_date=$5, job_id=$6, salary=$7, department_id=$8 
	WHERE employee_id = $9;`

	//check if the date is empty
	if emp.Hire_date == "" {
		log.Println("Warning hire_date is not empty")
	}

	formattedDate := formatDate(emp.Hire_date)

	//execute the query in emp
	_, err := DB.Exec(query, emp.First_name, emp.Last_name, emp.Email, emp.Phone_number, formattedDate, emp.Job_id, emp.Salary, emp.Department_id, emp.Employee_id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteEmployee(DB *sql.DB, id int) {

	//sql query for deleting the data 
	query := `DELETE FROM employees WHERE employee_id=$1`

	_, err := DB.Exec(query, id)

	if err != nil {
		log.Fatal(err)
	}
}
