package scaffold

import (
	"bytes"
	"database/sql"
	"errors"
	"log"
	"text/template"
)

var db *sql.DB
var tmpl *template.Template
var tables map[string]*Table

// NewTable generates a table
func NewTable(name string, cells []*Cell) *Table {
	t := new(Table)

	t.Name = name
	t.Cells = cells
	t.CellMap = make(map[string]int)

	for index, value := range cells {
		t.CellMap[value.Name] = index
	}

	tables[name] = t

	return t
}

// GetTable gets table def
func GetTable(name string) *Table {
	val, found := tables[name]
	if found {
		return val
	}

	return nil
}

// Raw runs a raw query
func Raw(q string) {
	db.Exec(q)
	return
}

// GetRaw runs a raw query that expects results
func GetRaw(q string) (*Rows, error) {
	result := new(Rows)

	result.Rows = make([]*Row, 0)
	result.Cols = make([]string, 0)

	rows, err := db.Query(q)
	if err != nil {
		return result, errors.New("Failure to execute query")
	}
	defer rows.Close()

	cols, err := rows.ColumnTypes()
	if err != nil {
		return result, errors.New("Failure to extract column types")
	}

	for _, v := range cols {
		result.Cols = append(result.Cols, v.Name())
	}

	for rows.Next() {
		row := new(Row)
		row.Cells = make([]*Cell, 0)

		scanList := make([]interface{}, 0)

		for _, c := range cols {
			cell := new(Cell)
			row.Cells = append(row.Cells, cell)

			cell.Name = c.Name()

			switch c.DatabaseTypeName() {
			case "BOOL", "BIT":
				cell.Type = CellBool
				scanList = append(scanList, &cell.BoolVal)
			case "TEXT", "VARCHAR", "NVARCHAR":
				cell.Type = CellString
				scanList = append(scanList, &cell.StringVal)
			case "INT", "INT4", "INT8", "BIGINT":
				cell.Type = CellInt
				scanList = append(scanList, &cell.IntVal)
			case "FLOAT", "FLOAT4", "FLOAT8", "DECIMAL", "MONEY", "NUMERIC":
				cell.Type = CellFloat
				scanList = append(scanList, &cell.FloatVal)
			case "DATETIME":
				cell.Type = CellTime
				scanList = append(scanList, &cell.TimeVal)
			default:
				panic("field type not accounted for: " + c.DatabaseTypeName())
			}
		}

		err := rows.Scan(scanList...)
		if err != nil {
			return result, errors.New("Failure to scan row")
		}
		result.Rows = append(result.Rows, row)
	}

	return result, nil
}

func init() {
	tables = make(map[string]*Table)

	var err error

	tmpl = new(template.Template)

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}

	tmpl, err = tmpl.New("query").Funcs(funcMap).Parse(queryTemplate)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = tmpl.New("schema").Funcs(funcMap).Parse(schemaTemplate)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = tmpl.New("insert").Funcs(funcMap).Parse(insertTemplate)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = tmpl.New("select").Funcs(funcMap).Parse(selectTemplate)
	if err != nil {
		log.Fatal(err)
	}
}

// Bootstrap connects the db
func Bootstrap(_db *sql.DB, tables []*Table) {
	db = _db

	for _, table := range tables {
		var b bytes.Buffer

		tmpl.ExecuteTemplate(&b, "schema", table)

		_, err := db.Exec(b.String())
		if err != nil {
			log.Fatal(err)
		}
	}
}
