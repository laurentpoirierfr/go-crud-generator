
package repository

import (
    "database/sql"
)

type Regions struct {
    Region_id string
    Region_name string
}

func CreateRegions(db *sql.DB, entity Regions) (sql.Result, error) {
    query := "INSERT INTO regions (region_id, region_name) VALUES (?, ?)"
    result, err := db.Exec(query, entity.Region_id, entity.Region_name)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetRegionsByID(db *sql.DB, id interface{}) (*Regions, error) {
    query := "SELECT * FROM regions WHERE region_id = ?"
    row := db.QueryRow(query, id)

    var entity Regions
    err := row.Scan(
        &entity.Region_id,
        &entity.Region_name,
    )
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func UpdateRegions(db *sql.DB, entity Regions, id interface{}) (sql.Result, error) {
    query := "UPDATE regions SET region_name = ? WHERE region_id = ?"
    args := []interface{}{
        entity.Region_id,
        entity.Region_name,
        id,
    }
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func DeleteRegions(db *sql.DB, id interface{}) (sql.Result, error) {
    query := "DELETE FROM regions WHERE region_id = ?"
    result, err := db.Exec(query, id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetAllRegions(db *sql.DB, offset, limit int) ([]Regions, error) {
    query := "SELECT * FROM regions LIMIT ? OFFSET ?"
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []Regions
    for rows.Next() {
        var entity Regions
        err := rows.Scan(
            &entity.Region_id,
            &entity.Region_name,
        )
        if err != nil {
            return nil, err
        }
        results = append(results, entity)
    }
    return results, nil
}
