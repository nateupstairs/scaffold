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
		row.Cells = make(map[string]*Cell, 0)

		scanList := make([]interface{}, 0)

		for _, c := range cols {
			cell := new(Cell)
			row.Cells[c.Name()] = cell

			cell.Name = c.Name()

			switch c.DatabaseTypeName() {
			case "BOOL", "BIT":
				data := NewSQLBool()
				cell.Data = data
				cell.Type = CellBool
				scanList = append(scanList, cell.CellTarget())
			case "BOOL[]", "BIT[]":
				xx := NewSQLBoolArray()

				cell.Data = xx
				scanList = append(scanList, cell.CellTarget())
			case "TEXT", "VARCHAR", "NVARCHAR":
				data := NewSQLString()
				cell.Data = data
				cell.Type = CellString
				scanList = append(scanList, cell.CellTarget())
			case "TEXT[]", "VARCHAR[]", "NVARCHAR[]":
				xx := NewSQLStringArray()

				cell.Data = xx
				scanList = append(scanList, cell.CellTarget())
			case "INT", "INT4", "INT8", "BIGINT":
				data := NewSQLInt()
				cell.Data = data
				cell.Type = CellInt
				scanList = append(scanList, cell.CellTarget())
			case "INT[]", "INT4[]", "INT8[]", "BIGINT[]":
				xx := NewSQLIntArray()

				cell.Data = xx
				scanList = append(scanList, cell.CellTarget())
			case "FLOAT", "FLOAT4", "FLOAT8", "DECIMAL", "MONEY", "NUMERIC":
				data := NewSQLFloat()
				cell.Data = data
				cell.Type = CellFloat
				scanList = append(scanList, cell.CellTarget())
			case "FLOAT[]", "FLOAT4[]", "FLOAT8[]", "DECIMAL[]", "MONEY[]", "NUMERIC[]":
				xx := NewSQLFloatArray()

				cell.Data = xx
				scanList = append(scanList, cell.CellTarget())
			case "DATE":
				data := NewSQLDate()
				cell.Data = data
				cell.Type = CellDate
				scanList = append(scanList, cell.CellTarget())
			case "DATE[]":
				xx := NewSQLDateArray()

				cell.Data = xx
				scanList = append(scanList, cell.CellTarget())
			case "DATETIME":
				data := NewSQLDatetime()
				cell.Data = data
				cell.Type = CellDatetime
				scanList = append(scanList, cell.CellTarget())
			case "DATETIME[]":
				xx := NewSQLDatetimeArray()

				cell.Data = xx
				scanList = append(scanList, cell.CellTarget())
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

	tmpl, err = tmpl.New("filter").Funcs(funcMap).Parse(filterTemplate)
	if err != nil {
		log.Fatal(err)
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
