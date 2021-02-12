package scaffold

import (
	"errors"

	"github.com/lib/pq"
)

// SQLBool representation of SQL
type SQLBool struct {
	Valid bool
	Value bool
}

// SQLBoolArray representation of SQL
type SQLBoolArray struct {
	Valid bool
	Value []bool
}

// NewSQLBool makes a SQLBool
func NewSQLBool() *SQLBool {
	x := new(SQLBool)
	x.Valid = false

	return x
}

// NewSQLBoolArray makes a SQLBoolArray
func NewSQLBoolArray() *SQLBoolArray {
	x := new(SQLBoolArray)
	x.Valid = true
	x.Value = make([]bool, 0)

	return x
}

// Raw Bool->Raw
func (x *SQLBool) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Raw BoolArray->Raw
func (x *SQLBoolArray) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Target gets the scannable target for SQLBool
func (x *SQLBool) Target() interface{} {
	return x
}

// Target gets the scannable target for SQLBoolArray
func (x *SQLBoolArray) Target() interface{} {
	return pq.Array(&x.Value)
}

// Scan interface->Bool
func (x *SQLBool) Scan(data interface{}) error {
	switch data.(type) {
	case bool:
		v, ok := data.(bool)
		if ok {
			x.Valid = true
			x.Value = v
		}
	case nil:
		x.Valid = false
		x.Value = false
	default:
		return errors.New("Incompatible type")
	}
	return nil
}
