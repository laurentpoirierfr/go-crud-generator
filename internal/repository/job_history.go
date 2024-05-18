package repository

import (
	"database/sql"
	"time"
)

type Job_history struct {
	Employee_id   int
	Start_date    time.Time
	End_date      time.Time
	Job_id        string
	Department_id int
}

func CreateJob_history(db *sql.DB, entity Job_history) (sql.Result, error) {
	query := "INSERT INTO job_history (employee_id, start_date, end_date, job_id, department_id) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(query, entity.Employee_id, entity.Start_date, entity.End_date, entity.Job_id, entity.Department_id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetJob_historyByID(db *sql.DB, id interface{}) (*Job_history, error) {
	query := "SELECT * FROM job_history WHERE  = ?"
	row := db.QueryRow(query, id)

	var entity Job_history
	err := row.Scan(
		&entity.Employee_id,
		&entity.Start_date,
		&entity.End_date,
		&entity.Job_id,
		&entity.Department_id,
	)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func UpdateJob_history(db *sql.DB, entity Job_history, id interface{}) (sql.Result, error) {
	query := "UPDATE job_history SET employee_id = ?, start_date = ?, end_date = ?, job_id = ?, department_id = ? WHERE  = ?"
	args := []interface{}{
		entity.Employee_id,
		entity.Start_date,
		entity.End_date,
		entity.Job_id,
		entity.Department_id,
		id,
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteJob_history(db *sql.DB, id interface{}) (sql.Result, error) {
	query := "DELETE FROM job_history WHERE  = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllJob_history(db *sql.DB, offset, limit int) ([]Job_history, error) {
	query := "SELECT * FROM job_history LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Job_history
	for rows.Next() {
		var entity Job_history
		err := rows.Scan(
			&entity.Employee_id,
			&entity.Start_date,
			&entity.End_date,
			&entity.Job_id,
			&entity.Department_id,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, entity)
	}
	return results, nil
}
