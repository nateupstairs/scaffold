package scaffold

import (
	"time"

	"github.com/tidwall/sjson"
)

// Row structure containing cells
type Row struct {
	Cells map[string]*Cell
}

// Rows structure that contains an array of rows and the column names
type Rows struct {
	Rows []*Row
	Cols []string
}

// AsJSON gets row data as json bytes
func (r *Row) AsJSON(cols []string) []byte {
	jsv := []byte("{}")

	for _, col := range cols {
		cell, ok := r.Cells[col]

		if ok {
			switch cell.Type {
			case CellBool, CellString, CellInt, CellFloat:
				v, err := cell.GetValue()
				if err == nil {
					jsv, _ = sjson.SetBytes(jsv, cell.Name, v)
				}
			case CellDate:
				v, err := cell.Date()
				if err == nil {
					jsv, _ = sjson.SetBytes(jsv, cell.Name, v.Format("2006-01-02"))
				}
			case CellDatetime:
				v, err := cell.Date()
				if err == nil {
					jsv, _ = sjson.SetBytes(jsv, cell.Name, v.Format(time.RFC3339))
				}
			case CellBoolArray:
				v, err := cell.BoolArray()
				if err == nil {
					cjsv := []byte("[]")

					if v != nil {
						for _, vv := range v {
							cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
						}
					}

					jsv, _ = sjson.SetRawBytes(jsv, cell.Name, cjsv)
				}
			case CellStringArray:
				v, err := cell.StringArray()
				if err == nil {
					cjsv := []byte("[]")

					if v != nil {
						for _, vv := range v {
							cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
						}
					}

					jsv, _ = sjson.SetRawBytes(jsv, cell.Name, cjsv)
				}
			case CellIntArray:
				v, err := cell.IntArray()
				if err == nil {
					cjsv := []byte("[]")

					if v != nil {
						for _, vv := range v {
							cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
						}
					}

					jsv, _ = sjson.SetRawBytes(jsv, cell.Name, cjsv)

				}
			case CellFloatArray:
				v, err := cell.FloatArray()
				if err == nil {
					cjsv := []byte("[]")

					if v != nil {
						for _, vv := range v {
							cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
						}
					}

					jsv, _ = sjson.SetRawBytes(jsv, cell.Name, cjsv)
				}
			case CellDateArray:
				v, err := cell.DateArray()
				if err == nil {
					cjsv := []byte("[]")

					if v != nil {
						for _, vv := range v {
							cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
						}
					}

					jsv, _ = sjson.SetRawBytes(jsv, cell.Name, cjsv)
				}
			case CellDatetimeArray:
				v, err := cell.DatetimeArray()
				if err == nil {
					cjsv := []byte("[]")

					if v != nil {
						for _, vv := range v {
							cjsv, _ = sjson.SetBytes(cjsv, "-1", vv)
						}
					}

					jsv, _ = sjson.SetRawBytes(jsv, cell.Name, cjsv)
				}
			}

		}
	}

	return jsv
}

// AsJSON gets rows data as json bytes
func (r *Rows) AsJSON() ([]byte, error) {
	jsv := []byte("{\"records\":[]}")

	for _, row := range r.Rows {
		var err error

		v := row.AsJSON(r.Cols)

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
