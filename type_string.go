package scaffold

import (
	"errors"

	"github.com/lib/pq"
)

// SQLString representation of SQL
type SQLString struct {
	Valid bool
	Value string
}

// SQLStringArray representation of SQL
type SQLStringArray struct {
	Valid bool
	Value []string
}

// NewSQLString makes a SQLString
func NewSQLString() *SQLString {
	x := new(SQLString)
	x.Valid = false

	return x
}

// NewSQLStringArray makes a SQLStringArray
func NewSQLStringArray() *SQLStringArray {
	x := new(SQLStringArray)
	x.Valid = true
	x.Value = make([]string, 0)

	return x
}

// Raw String->Raw
func (x *SQLString) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Raw StringArray->Raw
func (x *SQLStringArray) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Target gets the scannable target for SQLString
func (x *SQLString) Target() interface{} {
	return x
}

// Target gets the scannable target for SQLStringArray
func (x *SQLStringArray) Target() interface{} {
	return pq.Array(&x.Value)
}

// Scan interface->String
func (x *SQLString) Scan(data interface{}) error {
	switch data.(type) {
	case string:
		v, ok := data.(string)
		if ok {
			x.Valid = true
			x.Value = v
		}
	case nil:
		x.Valid = false
		x.Value = ""
	default:
		return errors.New("Incompatible type")
	}
	return nil
}
