package scaffold

import (
	"bytes"
	"database/sql"
	"errors"
	"log"
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

	switch cell.Type {
	case CellBool:
		row.Cells[indx].BoolVal = new(sql.NullBool)
		row.Cells[indx].BoolVal.Scan(val)
	case CellString:
		row.Cells[indx].StringVal = new(sql.NullString)
		row.Cells[indx].StringVal.Scan(val)
	case CellInt:
		row.Cells[indx].IntVal = new(sql.NullInt64)
		row.Cells[indx].IntVal.Scan(val)
	case CellFloat:
		row.Cells[indx].FloatVal = new(sql.NullFloat64)
		row.Cells[indx].FloatVal.Scan(val)
	case CellTime:
		row.Cells[indx].TimeVal = new(sql.NullTime)
		row.Cells[indx].TimeVal.Scan(val)
	}
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
func (t *Table) GetRows(q Query) *Rows {
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
		log.Fatal(err)
	}

	rows, err := db.Query(b.String())
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
		row := t.NewRow()
		scanList := make([]interface{}, 0)

		for _, c := range row.Cells {
			switch c.Type {
			case CellBool:
				scanList = append(scanList, &c.BoolVal)
			case CellString:
				scanList = append(scanList, &c.StringVal)
			case CellInt:
				scanList = append(scanList, &c.IntVal)
			case CellFloat:
				scanList = append(scanList, &c.FloatVal)
			case CellTime:
				scanList = append(scanList, &c.TimeVal)
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

// Insert inserts into a table
func (t *Table) Insert(row *Row) {
	fields := make([]string, 0)

	for _, c := range row.Cells {
		if !c.Exclude {
			fields = append(fields, c.Name)
		}
	}

	templateVars := make(map[string]interface{}, 0)
	templateVars["table"] = t
	templateVars["fields"] = fields

	var rowData = make([]interface{}, 0)

	for _, c := range row.Cells {
		if !c.Exclude {
			switch c.Type {
			case CellBool:
				rowData = append(rowData, c.BoolVal)
			case CellString:
				rowData = append(rowData, c.StringVal)
			case CellInt:
				rowData = append(rowData, c.IntVal)
			case CellFloat:
				rowData = append(rowData, c.FloatVal)
			case CellTime:
				rowData = append(rowData, c.TimeVal)
			}
		}
	}

	var b bytes.Buffer

	err := tmpl.ExecuteTemplate(&b, "insert", templateVars)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(b.String(), rowData...)
	if err != nil {
		log.Fatal(err)
	}
}
