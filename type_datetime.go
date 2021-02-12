package scaffold

import (
	"errors"
	"time"

	"github.com/lib/pq"
)

// SQLDatetime representation of SQL
type SQLDatetime struct {
	Valid bool
	Value time.Time
}

// SQLDatetimeArray representation of SQL
type SQLDatetimeArray struct {
	Valid bool
	Value []time.Time
}

// NewSQLDatetime makes a SQLDatetime
func NewSQLDatetime() *SQLDatetime {
	x := new(SQLDatetime)
	x.Valid = false

	return x
}

// NewSQLDatetimeArray makes a SQLDatetimeArray
func NewSQLDatetimeArray() *SQLDatetimeArray {
	x := new(SQLDatetimeArray)
	x.Valid = true
	x.Value = make([]time.Time, 0)

	return x
}

// Raw Date->Raw
func (x *SQLDatetime) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Raw DateArray->Raw
func (x *SQLDatetimeArray) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Target gets the scannable target for SQLDatetime
func (x *SQLDatetime) Target() interface{} {
	return x
}

// Target gets the scannable target for SQLDatetimeArray
func (x *SQLDatetimeArray) Target() interface{} {
	return pq.Array(&x.Value)
}

// Scan interface->Date
func (x *SQLDatetime) Scan(data interface{}) error {
	switch data.(type) {
	case time.Time:
		v, ok := data.(time.Time)
		if ok {
			x.Valid = true
			x.Value = v
		}
	case nil:
		x.Valid = false
		x.Value = time.Time{}
	default:
		return errors.New("Incompatible type")
	}
	return nil
}
