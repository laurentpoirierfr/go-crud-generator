package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

type Column struct {
	Name         string
	Type         string
	IsPrimaryKey bool
}

type Table struct {
	Name    string
	Columns []Column
}

func parseSchemaFile(filepath string) ([]Table, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tables []Table
	var currentTable *Table
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(strings.ToUpper(line), "CREATE TABLE") {
			tableName := extractTableName(line)
			currentTable = &Table{Name: tableName}
		} else if strings.Contains(strings.ToUpper(line), "PRIMARY KEY") {
			if currentTable != nil {
				pkColumn := extractPrimaryKey(line)
				for i, col := range currentTable.Columns {
					if col.Name == pkColumn {
						currentTable.Columns[i].IsPrimaryKey = true
					}
				}
			}
		} else if currentTable != nil && (strings.HasPrefix(line, ")") || strings.HasPrefix(line, ";")) {
			tables = append(tables, *currentTable)
			currentTable = nil
		} else if currentTable != nil {
			col := extractColumn(line)
			if col != nil {
				currentTable.Columns = append(currentTable.Columns, *col)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}

func extractTableName(line string) string {
	re := regexp.MustCompile(`CREATE TABLE (\w+)`)
	match := re.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractPrimaryKey(line string) string {
	re := regexp.MustCompile(`PRIMARY KEY \((\w+)\)`)
	match := re.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractColumn(line string) *Column {
	re := regexp.MustCompile(`(\w+)\s+(\w+)[\s,]*`)
	match := re.FindStringSubmatch(line)
	if len(match) > 2 {
		return &Column{Name: match[1], Type: match[2]}
	}
	return nil
}

const crudTemplate = `
package {{.PackageName}}

import (
    "database/sql"
)

type {{.TableName}} struct {
{{- range .Columns }}
    {{ toTitleCase .Name }} {{ .Type }}
{{- end }}
}

func Create{{.TableName}}(db *sql.DB, entity {{.TableName}}) (sql.Result, error) {
    query := "INSERT INTO {{.TableNameLower}} ({{.ColumnNames}}) VALUES ({{.Placeholders}})"
    result, err := db.Exec(query, {{.ColumnValues}})
    if err != nil {
        return nil, err
    }
    return result, nil
}

func Get{{.TableName}}ByID(db *sql.DB, id {{.PrimaryKeyType}}) (*{{.TableName}}, error) {
    query := "SELECT * FROM {{.TableNameLower}} WHERE {{.PrimaryKey}} = ?"
    row := db.QueryRow(query, id)

    var entity {{.TableName}}
    err := row.Scan(
{{- range .Columns }}
        &entity.{{ toTitleCase .Name }},
{{- end }}
    )
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func Update{{.TableName}}(db *sql.DB, entity {{.TableName}}, id {{.PrimaryKeyType}}) (sql.Result, error) {
    query := "UPDATE {{.TableNameLower}} SET {{.UpdatePlaceholders}} WHERE {{.PrimaryKey}} = ?"
    args := []interface{}{
{{- range .Columns }}
        entity.{{ toTitleCase .Name }},
{{- end }}
        id,
    }
    result, err := db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func Delete{{.TableName}}(db *sql.DB, id {{.PrimaryKeyType}}) (sql.Result, error) {
    query := "DELETE FROM {{.TableNameLower}} WHERE {{.PrimaryKey}} = ?"
    result, err := db.Exec(query, id)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetAll{{.TableName}}(db *sql.DB, offset, limit int) ([]{{.TableName}}, error) {
    query := "SELECT * FROM {{.TableNameLower}} LIMIT ? OFFSET ?"
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []{{.TableName}}
    for rows.Next() {
        var entity {{.TableName}}
        err := rows.Scan(
{{- range .Columns }}
            &entity.{{ toTitleCase .Name }},
{{- end }}
        )
        if err != nil {
            return nil, err
        }
        results = append(results, entity)
    }
    return results, nil
}
`

type CrudTemplateData struct {
	PackageName        string
	TableName          string
	TableNameLower     string
	PrimaryKey         string
	PrimaryKeyType     string
	Columns            []Column
	ColumnNames        string
	Placeholders       string
	ColumnValues       string
	UpdatePlaceholders string
	UpdateValues       string
}

func goType(sqlType string) string {
	switch strings.ToUpper(sqlType) {
	case "NUMBER":
		return "int"
	case "VARCHAR2", "CHAR":
		return "string"
	case "DATE":
		return "time.Time"
	default:
		return "string"
	}
}

func toTitleCase(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return str
}

func generateCrudCode(tables []Table, packageName, targetDir string) error {
	funcMap := template.FuncMap{
		"goType":      goType,
		"toTitleCase": toTitleCase,
	}

	tmpl, err := template.New("crud").Funcs(funcMap).Parse(crudTemplate)
	if err != nil {
		return err
	}

	// Create target directory if it doesn't exist
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		err := os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, table := range tables {
		for i, col := range table.Columns {
			table.Columns[i].Type = goType(col.Type)
		}
		primaryKey := ""
		primaryKeyType := ""
		for _, col := range table.Columns {
			if col.IsPrimaryKey {
				primaryKey = col.Name
				primaryKeyType = col.Type
				break
			}
		}

		var colNames []string
		var placeholders []string
		var colValues []string
		var updatePlaceholders []string
		var updateValues []string

		for _, col := range table.Columns {
			colNames = append(colNames, col.Name)
			placeholders = append(placeholders, "?")
			colValues = append(colValues, "entity."+toTitleCase(col.Name))
			if !col.IsPrimaryKey {
				updatePlaceholders = append(updatePlaceholders, col.Name+" = ?")
				updateValues = append(updateValues, "entity."+toTitleCase(col.Name))
			}
		}

		if len(primaryKey) == 0 {
			primaryKeyType = "interface{}"
		}

		data := CrudTemplateData{
			PackageName:        packageName,
			TableName:          toTitleCase(table.Name),
			TableNameLower:     strings.ToLower(table.Name),
			PrimaryKey:         primaryKey,
			PrimaryKeyType:     primaryKeyType,
			Columns:            table.Columns,
			ColumnNames:        strings.Join(colNames, ", "),
			Placeholders:       strings.Join(placeholders, ", "),
			ColumnValues:       strings.Join(colValues, ", "),
			UpdatePlaceholders: strings.Join(updatePlaceholders, ", "),
			UpdateValues:       strings.Join(updateValues, ", "),
		}

		file, err := os.Create(fmt.Sprintf("%s/%s.go", targetDir, strings.ToLower(table.Name)))
		if err != nil {
			return err
		}
		defer file.Close()

		err = tmpl.Execute(file, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <schema.sql> <package_name> <target_dir>")
		return
	}

	schemaFile := os.Args[1]
	packageName := os.Args[2]
	targetDir := os.Args[3]

	tables, err := parseSchemaFile(schemaFile)
	if err != nil {
		fmt.Println("Error parsing schema file:", err)
		return
	}

	err = generateCrudCode(tables, packageName, targetDir)
	if err != nil {
		fmt.Println("Error generating CRUD code:", err)
		return
	}

	fmt.Println("CRUD code generated successfully.")
}
