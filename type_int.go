package scaffold

import (
	"errors"

	"github.com/lib/pq"
)

// SQLInt representation of SQL
type SQLInt struct {
	Valid bool
	Value int64
}

// SQLIntArray representation of SQL
type SQLIntArray struct {
	Valid bool
	Value []int64
}

// NewSQLInt makes a SQLInt
func NewSQLInt() *SQLInt {
	x := new(SQLInt)
	x.Valid = false

	return x
}

// NewSQLIntArray makes a SQLIntArray
func NewSQLIntArray() *SQLIntArray {
	x := new(SQLIntArray)
	x.Valid = true
	x.Value = make([]int64, 0)

	return x
}

// Raw Int->Raw
func (x *SQLInt) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Raw IntArray->Raw
func (x *SQLIntArray) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Target gets the scannable target for SQLInt
func (x *SQLInt) Target() interface{} {
	return x
}

// Target gets the scannable target for SQLIntArray
func (x *SQLIntArray) Target() interface{} {
	return pq.Array(&x.Value)
}

// Scan interface->Int
func (x *SQLInt) Scan(data interface{}) error {
	switch data.(type) {
	case int64:
		v, ok := data.(int64)
		if ok {
			x.Valid = true
			x.Value = v
		}
	case nil:
		x.Valid = false
		x.Value = 0
	default:
		return errors.New("Incompatible type")
	}
	return nil
}
