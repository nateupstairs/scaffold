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
		switch cell.Type {
		case CellBool, CellString, CellInt, CellFloat, CellTime:
			v, err := cell.GetValue()
			if err == nil {
				jsv, _ = sjson.SetBytes(jsv, cell.Name, v)
			}
		case CellBoolArray:
			v := cell.BoolArrayVal
			cjsv := []byte("[]")

			if v != nil {
				for _, vv := range v.Value {
					cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
				}
			}

			jsv, _ = sjson.SetBytes(jsv, cell.Name, cjsv)
		case CellStringArray:
			v := cell.StringArrayVal
			cjsv := []byte("[]")

			if v != nil {
				for _, vv := range v.Value {
					cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
				}
			}

			jsv, _ = sjson.SetBytes(jsv, cell.Name, cjsv)
		case CellIntArray:
			v := cell.IntArrayVal
			cjsv := []byte("[]")

			if v != nil {
				for _, vv := range v.Value {
					cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
				}
			}

			jsv, _ = sjson.SetBytes(jsv, cell.Name, cjsv)
		case CellFloatArray:
			v := cell.FloatArrayVal
			cjsv := []byte("[]")

			if v != nil {
				for _, vv := range v.Value {
					cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
				}
			}

			jsv, _ = sjson.SetBytes(jsv, cell.Name, cjsv)
		}
	}

	return jsv
}

// AsJSON gets rows data as json bytes
func (r *Rows) AsJSON() ([]byte, error) {
	jsv := []byte("{\"records\":[]}")

	for _, row := range r.Rows {
		var err error

		v := row.AsJSON()

		jsv, err = sjson.SetRawBytes(jsv, "records.-1", v)
		if err != nil {
			return jsv, err
		}
	}

	return jsv, nil
}

// MarshalJSON to more flexibly deal with variant data
func (r *Row) MarshalJSON() ([]byte, error) {
	b := []byte("{}")

	for _, cell := range r.Cells {
		v, err := cell.GetValue()
		if err == nil {
			b, _ = sjson.SetBytes(b, cell.Name, v)
		} else {
			return b, err
		}
	}

	return b, nil
}
