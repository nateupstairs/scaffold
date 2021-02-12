package scaffold

import (
	"errors"
	"time"

	"github.com/lib/pq"
)

// SQLDate representation of SQL
type SQLDate struct {
	Valid bool
	Value time.Time
}

// SQLDateArray representation of SQL
type SQLDateArray struct {
	Valid bool
	Value []time.Time
}

// NewSQLDate makes a SQLDate
func NewSQLDate() *SQLDate {
	x := new(SQLDate)
	x.Valid = false

	return x
}

// NewSQLDateArray makes a SQLDateArray
func NewSQLDateArray() *SQLDateArray {
	x := new(SQLDateArray)
	x.Valid = true
	x.Value = make([]time.Time, 0)

	return x
}

// Raw Date->Raw
func (x *SQLDate) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Raw DateArray->Raw
func (x *SQLDateArray) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Target gets the scannable target for SQLDate
func (x *SQLDate) Target() interface{} {
	return x
}

// Target gets the scannable target for SQLDateArray
func (x *SQLDateArray) Target() interface{} {
	return pq.Array(&x.Value)
}

// Scan interface->Date
func (x *SQLDate) Scan(data interface{}) error {
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
