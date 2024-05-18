
package repository

import (
    "database/sql"
)

type Countries struct {
    Country_id string
    Country_name string
    Region_id int
}

func CreateCountries(db *sql.DB, entity Countries) (sql.Result, error) {
    query := "INSERT INTO countries (country_id, country_name, region_id) VALUES (?, ?, ?)"
    result, err := db.Exec(query, entity.Country_id, entity.Country_name, entity.Region_id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetCountriesByID(db *sql.DB, id string) (*Countries, error) {
    query := "SELECT * FROM countries WHERE country_id = ?"
    row := db.QueryRow(query, id)

    var entity Countries
    err := row.Scan(
        &entity.Country_id,
        &entity.Country_name,
        &entity.Region_id,
    )
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func UpdateCountries(db *sql.DB, entity Countries, id string) (sql.Result, error) {
    query := "UPDATE countries SET country_name = ?, region_id = ? WHERE country_id = ?"
    args := []interface{}{
        entity.Country_id,
        entity.Country_name,
        entity.Region_id,
        id,
    }
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func DeleteCountries(db *sql.DB, id string) (sql.Result, error) {
    query := "DELETE FROM countries WHERE country_id = ?"
    result, err := db.Exec(query, id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetAllCountries(db *sql.DB, offset, limit int) ([]Countries, error) {
    query := "SELECT * FROM countries LIMIT ? OFFSET ?"
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []Countries
    for rows.Next() {
        var entity Countries
        err := rows.Scan(
            &entity.Country_id,
            &entity.Country_name,
            &entity.Region_id,
        )
        if err != nil {
            return nil, err
        }
        results = append(results, entity)
    }
    return results, nil
}
