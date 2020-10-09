package scaffold

import (
	"github.com/tidwall/sjson"
)

type Row struct {
	Cells []*Cell
}

type Rows struct {
	Rows []*Row
	Cols []string
}

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

func (r *Rows) AsJSON() []byte {
	jsv := []byte("{\"records\":[]}")

	for _, row := range r.Rows {
		v := row.AsJSON()

		jsv, _ = sjson.SetRawBytes(jsv, "records.-1", v)
	}

	return jsv
}
