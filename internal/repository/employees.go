
package repository

import (
    "database/sql"
)

type Employees struct {
    Employee_id string
    First_name string
    Last_name string
    Email string
    Phone_number string
    Hire_date string
    Job_id string
    Salary string
    Commission_pct string
    Manager_id string
    Department_id string
}

func CreateEmployees(db *sql.DB, entity Employees) (sql.Result, error) {
    query := "INSERT INTO employees (employee_id, first_name, last_name, email, phone_number, hire_date, job_id, salary, commission_pct, manager_id, department_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
    result, err := db.Exec(query, entity.Employee_id, entity.First_name, entity.Last_name, entity.Email, entity.Phone_number, entity.Hire_date, entity.Job_id, entity.Salary, entity.Commission_pct, entity.Manager_id, entity.Department_id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetEmployeesByID(db *sql.DB, id interface{}) (*Employees, error) {
    query := "SELECT * FROM employees WHERE employee_id = ?"
    row := db.QueryRow(query, id)

    var entity Employees
    err := row.Scan(
        &entity.Employee_id,
        &entity.First_name,
        &entity.Last_name,
        &entity.Email,
        &entity.Phone_number,
        &entity.Hire_date,
        &entity.Job_id,
        &entity.Salary,
        &entity.Commission_pct,
        &entity.Manager_id,
        &entity.Department_id,
    )
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func UpdateEmployees(db *sql.DB, entity Employees, id interface{}) (sql.Result, error) {
    query := "UPDATE employees SET first_name = ?, last_name = ?, email = ?, phone_number = ?, hire_date = ?, job_id = ?, salary = ?, commission_pct = ?, manager_id = ?, department_id = ? WHERE employee_id = ?"
    args := []interface{}{
        entity.Employee_id,
        entity.First_name,
        entity.Last_name,
        entity.Email,
        entity.Phone_number,
        entity.Hire_date,
        entity.Job_id,
        entity.Salary,
        entity.Commission_pct,
        entity.Manager_id,
        entity.Department_id,
        id,
    }
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func DeleteEmployees(db *sql.DB, id interface{}) (sql.Result, error) {
    query := "DELETE FROM employees WHERE employee_id = ?"
    result, err := db.Exec(query, id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetAllEmployees(db *sql.DB, offset, limit int) ([]Employees, error) {
    query := "SELECT * FROM employees LIMIT ? OFFSET ?"
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []Employees
    for rows.Next() {
        var entity Employees
        err := rows.Scan(
            &entity.Employee_id,
            &entity.First_name,
            &entity.Last_name,
            &entity.Email,
            &entity.Phone_number,
            &entity.Hire_date,
            &entity.Job_id,
            &entity.Salary,
            &entity.Commission_pct,
            &entity.Manager_id,
            &entity.Department_id,
        )
        if err != nil {
            return nil, err
        }
        results = append(results, entity)
    }
    return results, nil
}
