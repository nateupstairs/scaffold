package scaffold

import (
	"bytes"
	"database/sql"
	"log"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
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

func Raw(q string) {
	db.Exec(q)
	return
}

func GetRaw(q string) *Rows {
	result := new(Rows)

	result.Rows = make([]*Row, 0)
	result.Cols = make([]string, 0)

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, err := rows.ColumnTypes()
	if err != nil {
		return result
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
			case "BOOL":
				cell.Type = CellBool
				scanList = append(scanList, &cell.BoolVal)
			case "TEXT", "VARCHAR", "NVARCHAR":
				cell.Type = CellString
				scanList = append(scanList, &cell.StringVal)
			case "INT", "BIGINT", "INT8":
				cell.Type = CellInt
				scanList = append(scanList, &cell.IntVal)
			case "FLOAT", "DECIMAL":
				cell.Type = CellFloat
				scanList = append(scanList, &cell.FloatVal)
			default:
				panic("field type not accounted for" + c.DatabaseTypeName())
			}
		}

		err := rows.Scan(scanList...)
		if err != nil {
			log.Fatal(err)
		}
		result.Rows = append(result.Rows, row)
	}

	return result
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

	tmpl, err = tmpl.New("schema").Funcs(funcMap).Parse(`
		CREATE TABLE IF NOT EXISTS {{.Name}} (
			{{ range $index, $cell := .Cells -}}
				{{if $index}},{{end -}}
				{{$cell.Name}} {{$cell.SQL}}
			{{end}}
		)
	`)
	if err != nil {
		spew.Dump(err)
	}

	tmpl, err = tmpl.New("insert").Funcs(funcMap).Parse(`
		INSERT INTO {{.table.Name}} (
			{{ range $index, $field := .fields -}}
				{{if $index}},{{end -}}
				{{$field}}
			{{end}}
		)
		values (
			{{ range $index, $field := .fields -}}
				{{- if $index}},{{end -}}
				${{inc $index}}
			{{end}}
		)
	`)
	if err != nil {
		spew.Dump(err)
	}

	tmpl, err = tmpl.New("select").Funcs(funcMap).Parse(`
		SELECT
			{{ range $index, $field := .fields -}}
				{{if $index}},{{end}}{{$field}}
			{{end}}
		FROM {{.table.Name}}
		{{ if .query -}}
			{{ if .query.Filters -}}
				{{ range $index, $filter := .query.Filters -}}
				{{if $index}}AND{{else}}WHERE{{end}} {{$filter.Field}} {{$filter.Comparison}} {{$filter.Value}}
				{{- end -}}
			{{- end -}}
			{{- if .query.Orders -}}
				{{- range $index, $order := .query.Orders -}}
				{{- if $index}},{{else}}
		ORDER BY{{end}} {{$order.Field}} {{$order.Direction}}
				{{- end -}}
			{{- end }}
		LIMIT {{.query.Limit}}
		OFFSET {{.query.Offset}}
		{{- end -}}
	`)
	if err != nil {
		spew.Dump(err)
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
