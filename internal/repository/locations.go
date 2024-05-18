
package repository

import (
    "database/sql"
)

type Locations struct {
    Location_id string
    Street_address string
    Postal_code string
    City string
    State_province string
    Country_id string
}

func CreateLocations(db *sql.DB, entity Locations) (sql.Result, error) {
    query := "INSERT INTO locations (location_id, street_address, postal_code, city, state_province, country_id) VALUES (?, ?, ?, ?, ?, ?)"
    result, err := db.Exec(query, entity.Location_id, entity.Street_address, entity.Postal_code, entity.City, entity.State_province, entity.Country_id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetLocationsByID(db *sql.DB, id interface{}) (*Locations, error) {
    query := "SELECT * FROM locations WHERE location_id = ?"
    row := db.QueryRow(query, id)

    var entity Locations
    err := row.Scan(
        &entity.Location_id,
        &entity.Street_address,
        &entity.Postal_code,
        &entity.City,
        &entity.State_province,
        &entity.Country_id,
    )
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func UpdateLocations(db *sql.DB, entity Locations, id interface{}) (sql.Result, error) {
    query := "UPDATE locations SET street_address = ?, postal_code = ?, city = ?, state_province = ?, country_id = ? WHERE location_id = ?"
    args := []interface{}{
        entity.Location_id,
        entity.Street_address,
        entity.Postal_code,
        entity.City,
        entity.State_province,
        entity.Country_id,
        id,
    }
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func DeleteLocations(db *sql.DB, id interface{}) (sql.Result, error) {
    query := "DELETE FROM locations WHERE location_id = ?"
    result, err := db.Exec(query, id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetAllLocations(db *sql.DB, offset, limit int) ([]Locations, error) {
    query := "SELECT * FROM locations LIMIT ? OFFSET ?"
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []Locations
    for rows.Next() {
        var entity Locations
        err := rows.Scan(
            &entity.Location_id,
            &entity.Street_address,
            &entity.Postal_code,
            &entity.City,
            &entity.State_province,
            &entity.Country_id,
        )
        if err != nil {
            return nil, err
        }
        results = append(results, entity)
    }
    return results, nil
}
