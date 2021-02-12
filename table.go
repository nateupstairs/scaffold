package scaffold

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// Table structure
type Table struct {
	Name  string
	Cells []*Cell
}

// NewRow creates a row that conforms to the table definition
func (t *Table) NewRow() *Row {
	row := new(Row)
	row.Cells = make(map[string]*Cell, 0)

	for _, proto := range t.Cells {
		cell := new(Cell)

		cell.Name = proto.Name
		cell.Type = proto.Type
		cell.SQL = proto.SQL
		row.Cells[cell.Name] = cell
	}

	return row
}

// GetRows runs a query and returns a rows structure
func (t *Table) GetRows(q Query) (*Rows, error) {
	result := new(Rows)
	fields := make([]string, 0)

	result.Rows = make([]*Row, 0)
	result.Cols = make([]string, 0)

	for _, c := range t.Cells {
		if !c.Exclude {
			fields = append(fields, c.Name)
		}
	}

	templateVars := make(map[string]interface{}, 0)
	templateVars["table"] = t
	templateVars["fields"] = fields
	templateVars["query"] = q

	var b bytes.Buffer

	err := tmpl.ExecuteTemplate(&b, "select", templateVars)
	if err != nil {
		return result, errors.New("Failure to execute template")
	}

	rows, err := db.Query(b.String())
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
		row := t.NewRow()
		scanList := make([]interface{}, 0)

		for _, col := range t.Cells {
			c := row.Cells[col.Name]

			switch c.Type {
			case CellBool:
				data := NewSQLBool()
				c.Data = data
				scanList = append(scanList, c.CellTarget())
			case CellString:
				data := NewSQLString()
				c.Data = data
				scanList = append(scanList, c.CellTarget())
			case CellInt:
				data := NewSQLInt()
				c.Data = data
				scanList = append(scanList, c.CellTarget())
			case CellFloat:
				data := NewSQLFloat()
				c.Data = data
				scanList = append(scanList, c.CellTarget())
			case CellDate:
				data := NewSQLDate()
				c.Data = data
				scanList = append(scanList, c.CellTarget())
			case CellDatetime:
				data := NewSQLDatetime()
				c.Data = data
				scanList = append(scanList, c.CellTarget())
			case CellBoolArray:
				xx := NewSQLBoolArray()

				c.Data = xx
				scanList = append(scanList, c.CellTarget())
			case CellStringArray:
				xx := NewSQLStringArray()

				c.Data = xx
				scanList = append(scanList, c.CellTarget())
			case CellIntArray:
				xx := NewSQLIntArray()

				c.Data = xx
				scanList = append(scanList, c.CellTarget())
			case CellFloatArray:
				xx := NewSQLFloatArray()

				c.Data = xx
				scanList = append(scanList, c.CellTarget())
			case CellDateArray:
				xx := NewSQLDateArray()

				c.Data = xx
				scanList = append(scanList, c.CellTarget())
			case CellDatetimeArray:
				xx := NewSQLDatetimeArray()

				c.Data = xx
				scanList = append(scanList, c.CellTarget())
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

// Insert inserts into a table
func (t *Table) Insert(row *Row) error {
	fields := make([]string, 0)
	placeholders := make([]string, 0)

	for _, tc := range t.Cells {
		c, ok := row.Cells[tc.Name]
		if ok {
			if !c.Exclude {
				fields = append(fields, c.Name)
			}
		}
	}

	templateVars := make(map[string]interface{}, 0)
	templateVars["table"] = t
	templateVars["fields"] = fields

	var rowData = make([]interface{}, 0)
	var placeholderCursor = 1

	for _, col := range t.Cells {
		c, ok := row.Cells[col.Name]
		if ok {
			if !c.Exclude {
				switch c.Type {
				case CellBool, CellString, CellInt, CellFloat, CellDate, CellDatetime:
					value, err := c.GetValue()
					if err != nil {
						return err
					}
					rowData = append(rowData, value)
					placeholders = append(placeholders, "$"+strconv.Itoa(placeholderCursor))
				case CellBoolArray:
					v, err := c.BoolArray()
					if err == nil {
						for _, val := range v {
							rowData = append(rowData, val)
						}

						var p string
						for i := 0; i < len(v); i++ {
							p += "$" + strconv.Itoa(placeholderCursor)
							if i < len(v)-1 {
								p += ","
								placeholderCursor++
							}
						}

						placeholders = append(placeholders, "ARRAY["+p+"]::bool[]")
					}
				case CellStringArray:
					v, err := c.StringArray()
					if err == nil {
						for _, val := range v {
							rowData = append(rowData, val)
						}

						var p string
						for i := 0; i < len(v); i++ {
							p += "$" + strconv.Itoa(placeholderCursor)
							if i < len(v)-1 {
								p += ","
								placeholderCursor++
							}
						}

						placeholders = append(placeholders, "ARRAY["+p+"]::text[]")
					}
				case CellIntArray:
					v, err := c.IntArray()
					if err == nil {
						for _, val := range v {
							rowData = append(rowData, val)
						}

						var p string
						for i := 0; i < len(v); i++ {
							p += "$" + strconv.Itoa(placeholderCursor)
							if i < len(v)-1 {
								p += ","
								placeholderCursor++
							}
						}

						placeholders = append(placeholders, "ARRAY["+p+"]::integer[]")
					}
				case CellFloatArray:
					v, err := c.FloatArray()
					if err == nil {
						for _, val := range v {
							rowData = append(rowData, val)
						}

						var p string
						for i := 0; i < len(v); i++ {
							p += "$" + strconv.Itoa(placeholderCursor)
							if i < len(v)-1 {
								p += ","
								placeholderCursor++
							}
						}

						placeholders = append(placeholders, "ARRAY["+p+"]::float[]")
					}
				case CellDateArray:
					v, err := c.DateArray()
					if err == nil {
						for _, val := range v {
							rowData = append(rowData, val)
						}

						var p string
						for i := 0; i < len(v); i++ {
							p += "$" + strconv.Itoa(placeholderCursor)
							if i < len(v)-1 {
								p += ","
								placeholderCursor++
							}
						}

						placeholders = append(placeholders, "ARRAY["+p+"]::date[]")
					}
				case CellDatetimeArray:
					v, err := c.DateArray()
					if err == nil {
						for _, val := range v {
							rowData = append(rowData, val)
						}

						var p string
						for i := 0; i < len(v); i++ {
							p += "$" + strconv.Itoa(placeholderCursor)
							if i < len(v)-1 {
								p += ","
								placeholderCursor++
							}
						}

						placeholders = append(placeholders, "ARRAY["+p+"]::datetime[]")
					}
				}

				placeholderCursor++
			}
		}
	}

	templateVars["placeholders"] = placeholders

	var b bytes.Buffer

	err := tmpl.ExecuteTemplate(&b, "insert", templateVars)
	if err != nil {
		return errors.New("Failure to execute template")
	}

	_, err = db.Exec(b.String(), rowData...)
	if err != nil {
		spew.Dump(err)
		return errors.New("Failure to execute query")
	}

	return nil
}
