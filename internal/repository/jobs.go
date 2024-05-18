
package repository

import (
    "database/sql"
)

type Jobs struct {
    Job_id string
    Job_title string
    Min_salary string
    Max_salary string
}

func CreateJobs(db *sql.DB, entity Jobs) (sql.Result, error) {
    query := "INSERT INTO jobs (job_id, job_title, min_salary, max_salary) VALUES (?, ?, ?, ?)"
    result, err := db.Exec(query, entity.Job_id, entity.Job_title, entity.Min_salary, entity.Max_salary)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetJobsByID(db *sql.DB, id interface{}) (*Jobs, error) {
    query := "SELECT * FROM jobs WHERE job_id = ?"
    row := db.QueryRow(query, id)

    var entity Jobs
    err := row.Scan(
        &entity.Job_id,
        &entity.Job_title,
        &entity.Min_salary,
        &entity.Max_salary,
    )
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func UpdateJobs(db *sql.DB, entity Jobs, id interface{}) (sql.Result, error) {
    query := "UPDATE jobs SET job_title = ?, min_salary = ?, max_salary = ? WHERE job_id = ?"
    args := []interface{}{
        entity.Job_id,
        entity.Job_title,
        entity.Min_salary,
        entity.Max_salary,
        id,
    }
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func DeleteJobs(db *sql.DB, id interface{}) (sql.Result, error) {
    query := "DELETE FROM jobs WHERE job_id = ?"
    result, err := db.Exec(query, id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetAllJobs(db *sql.DB, offset, limit int) ([]Jobs, error) {
    query := "SELECT * FROM jobs LIMIT ? OFFSET ?"
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []Jobs
    for rows.Next() {
        var entity Jobs
        err := rows.Scan(
            &entity.Job_id,
            &entity.Job_title,
            &entity.Min_salary,
            &entity.Max_salary,
        )
        if err != nil {
            return nil, err
        }
        results = append(results, entity)
    }
    return results, nil
}
