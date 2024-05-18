
package repository

import (
    "database/sql"
)

type Departments struct {
    Department_id string
    Department_name string
    Manager_id string
    Location_id string
}

func CreateDepartments(db *sql.DB, entity Departments) (sql.Result, error) {
    query := "INSERT INTO departments (department_id, department_name, manager_id, location_id) VALUES (?, ?, ?, ?)"
    result, err := db.Exec(query, entity.Department_id, entity.Department_name, entity.Manager_id, entity.Location_id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetDepartmentsByID(db *sql.DB, id interface{}) (*Departments, error) {
    query := "SELECT * FROM departments WHERE department_id = ?"
    row := db.QueryRow(query, id)

    var entity Departments
    err := row.Scan(
        &entity.Department_id,
        &entity.Department_name,
        &entity.Manager_id,
        &entity.Location_id,
    )
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func UpdateDepartments(db *sql.DB, entity Departments, id interface{}) (sql.Result, error) {
    query := "UPDATE departments SET department_name = ?, manager_id = ?, location_id = ? WHERE department_id = ?"
    args := []interface{}{
        entity.Department_id,
        entity.Department_name,
        entity.Manager_id,
        entity.Location_id,
        id,
    }
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func DeleteDepartments(db *sql.DB, id interface{}) (sql.Result, error) {
    query := "DELETE FROM departments WHERE department_id = ?"
    result, err := db.Exec(query, id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetAllDepartments(db *sql.DB, offset, limit int) ([]Departments, error) {
    query := "SELECT * FROM departments LIMIT ? OFFSET ?"
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []Departments
    for rows.Next() {
        var entity Departments
        err := rows.Scan(
            &entity.Department_id,
            &entity.Department_name,
            &entity.Manager_id,
            &entity.Location_id,
        )
        if err != nil {
            return nil, err
        }
        results = append(results, entity)
    }
    return results, nil
}
