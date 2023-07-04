package main
//Import necessary packages
import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	_ "mattn/go-sqlite3"
)

func main() {
	start := time.Now()
	bulkInsert()
	//Calculating total execution time
	elapsed := time.Since(start)
	fmt.Printf("The time of execution of Go program is : %s", elapsed)
}

// To get SQLite connection
func getSQLiteDB() (*sql.DB, error) {
	dbConn, err := sql.Open("sqlite3", "./sqlite_employees.db")
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// Function helps to insert data in bulk
func bulkInsert() error {
	dbConn, err := getSQLiteDB()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	records, err := loadCSV()
	if err != nil {
		return err
	}
	transaction, err := dbConn.Begin()
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer([]byte("INSERT INTO employee_details (emp_id, emp_name, emp_dob, emp_role, emp_dept) VALUES "))
	for i := 0; i < 249; i++ {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString("(?,?,?,?,?)")
	}
	bulkInsertStmt, err := transaction.Prepare(buf.String())
	if err != nil {
		return err
	}
	var values []interface{}
	for _, record := range records {

		emp_id, _ := strconv.Atoi(record[0])
		values = append(values, emp_id)
		values = append(values, record[1])
		values = append(values, record[2])
		values = append(values, record[3])
		values = append(values, record[4])
		if len(values) == 996 {
			if _, err = bulkInsertStmt.Exec(values...); err != nil {
				return err
			}
			values = values[0:0]
		}
	}
	if len(values) > 0 {
		buf := bytes.NewBuffer([]byte("INSERT INTO employee_details (emp_id, emp_name, emp_dob, emp_role, emp_dept) VALUES "))
		for i := 0; i < len(values)/5; i++ {
			if i > 0 {
				buf.WriteString(",")
			}
			buf.WriteString("(?,?,?,?, ?)")
		}
		if _, err = transaction.Exec(buf.String(), values...); err != nil {
			return err
		}
	}
	return transaction.Commit()
}

func loadCSV() ([][]string, error) {
	f, err := os.Open("employees.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	return r.ReadAll()
}
