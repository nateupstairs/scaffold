package scaffold

import (
	"github.com/tidwall/sjson"
)

// Row structure containing cells
type Row struct {
	Cells []*Cell
}

// Rows structure that contains an array of rows and the column names
type Rows struct {
	Rows []*Row
	Cols []string
}

// AsJSON gets row data as json bytes
func (r *Row) AsJSON() []byte {
	jsv := []byte("{}")

	for _, cell := range r.Cells {
		v, err := cell.GetValue()
		if err == nil {
			jsv, _ = sjson.SetBytes(jsv, cell.Name, v)
		}
	}

	return jsv
}

// AsJSON gets rows data as json bytes
func (r *Rows) AsJSON() []byte {
	jsv := []byte("{\"records\":[]}")

	for _, row := range r.Rows {
		v := row.AsJSON()

		jsv, _ = sjson.SetRawBytes(jsv, "records.-1", v)
	}

	return jsv
}
