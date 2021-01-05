package scaffold

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/lib/pq"
)

// Table structure
type Table struct {
	Name    string
	Cells   []*Cell
	CellMap map[string]int
}

// GetCellIndex gets index for cell name in table definition
func (t *Table) GetCellIndex(needle string) (int, error) {
	v, found := t.CellMap[needle]
	if found {
		return v, nil
	}

	return 0, errors.New("Invalid cell name")
}

// NewRow creates a row that conforms to the table definition
func (t *Table) NewRow() *Row {
	row := new(Row)
	row.Cells = make([]*Cell, len(t.Cells))

	for indx, proto := range t.Cells {
		cell := new(Cell)

		cell.Name = proto.Name
		cell.Type = proto.Type
		cell.SQL = proto.SQL
		row.Cells[indx] = cell
	}

	return row
}

// SetCell sets a row's cell by name
func (t *Table) SetCell(row *Row, needle string, val interface{}) {
	indx, err := t.GetCellIndex(needle)
	if err != nil {
		return
	}

	cellDef := t.Cells[indx]
	cell := row.Cells[indx]
	if cell == nil {
		c := new(Cell)

		c.Name = cellDef.Name
		c.Type = cellDef.Type
		c.SQL = cellDef.SQL

		row.Cells[indx] = c
	}

	_ = cell.SetValue(val)
	return
}

// GetCell gets a row's data by cell name
func (t *Table) GetCell(row *Row, needle string) (interface{}, error) {
	indx, err := t.GetCellIndex(needle)
	if err != nil {
		return nil, errors.New("Invalid cell")
	}

	cell := row.Cells[indx]

	return cell.GetValue()
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

		for _, c := range row.Cells {
			switch c.Type {
			case CellBool, CellString, CellInt, CellFloat, CellTime:
				scanList = append(scanList, c)
			case CellBoolArray:
				xx := NewSQLBoolArray()

				c.BoolArrayVal = xx
				scanList = append(scanList, pq.Array(&xx.Value))
			case CellStringArray:
				xx := NewSQLStringArray()

				c.StringArrayVal = xx
				scanList = append(scanList, pq.Array(&xx.Value))
			case CellIntArray:
				xx := NewSQLIntArray()

				c.IntArrayVal = xx
				scanList = append(scanList, pq.Array(&xx.Value))
			case CellFloatArray:
				xx := NewSQLFloatArray()

				c.FloatArrayVal = xx
				scanList = append(scanList, pq.Array(&xx.Value))
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

	for _, c := range row.Cells {
		if !c.Exclude {
			fields = append(fields, c.Name)
		}
	}

	templateVars := make(map[string]interface{}, 0)
	templateVars["table"] = t
	templateVars["fields"] = fields

	var rowData = make([]interface{}, 0)
	var placeholderCursor = 1

	for _, c := range row.Cells {
		if !c.Exclude {
			switch c.Type {
			case CellBool, CellString, CellInt, CellFloat, CellTime:
				value, err := c.GetValue()
				if err != nil {
					return err
				}
				rowData = append(rowData, value)
				placeholders = append(placeholders, "$"+strconv.Itoa(placeholderCursor))
			case CellBoolArray:
				for _, val := range c.BoolArrayVal.Value {
					rowData = append(rowData, val)
				}

				var p string
				for i := 0; i < len(c.BoolArrayVal.Value); i++ {
					p += "$" + strconv.Itoa(placeholderCursor)
					if i < len(c.BoolArrayVal.Value)-1 {
						p += ","
						placeholderCursor++
					}
				}

				placeholders = append(placeholders, "ARRAY["+p+"]::bool[]")
			case CellStringArray:
				for _, val := range c.StringArrayVal.Value {
					rowData = append(rowData, val)
				}

				var p string
				for i := 0; i < len(c.StringArrayVal.Value); i++ {
					p += "$" + strconv.Itoa(placeholderCursor)
					if i < len(c.StringArrayVal.Value)-1 {
						p += ","
						placeholderCursor++
					}
				}

				placeholders = append(placeholders, "ARRAY["+p+"]::text[]")
			case CellIntArray:
				for _, val := range c.IntArrayVal.Value {
					rowData = append(rowData, val)
				}

				var p string
				for i := 0; i < len(c.IntArrayVal.Value); i++ {
					p += "$" + strconv.Itoa(placeholderCursor)
					if i < len(c.IntArrayVal.Value)-1 {
						p += ","
						placeholderCursor++
					}
				}

				placeholders = append(placeholders, "ARRAY["+p+"]::integer[]")
			case CellFloatArray:
				for _, val := range c.FloatArrayVal.Value {
					rowData = append(rowData, val)
				}

				var p string
				for i := 0; i < len(c.FloatArrayVal.Value); i++ {
					p += "$" + strconv.Itoa(placeholderCursor)
					if i < len(c.FloatArrayVal.Value)-1 {
						p += ","
						placeholderCursor++
					}
				}

				placeholders = append(placeholders, "ARRAY["+p+"]::float[]")
			}

			placeholderCursor++
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
		return errors.New("Failure to execute query")
	}

	return nil
}
