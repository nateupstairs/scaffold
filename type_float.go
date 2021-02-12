package scaffold

import (
	"errors"

	"github.com/lib/pq"
)

// SQLFloat representation of SQL
type SQLFloat struct {
	Valid bool
	Value float64
}

// SQLFloatArray representation of SQL
type SQLFloatArray struct {
	Valid bool
	Value []float64
}

// NewSQLFloat makes a SQLFloat
func NewSQLFloat() *SQLFloat {
	x := new(SQLFloat)
	x.Valid = false

	return x
}

// NewSQLFloatArray makes a SQLFloatArray
func NewSQLFloatArray() *SQLFloatArray {
	x := new(SQLFloatArray)
	x.Valid = true
	x.Value = make([]float64, 0)

	return x
}

// Raw Float->Raw
func (x *SQLFloat) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Raw FloatArray->Raw
func (x *SQLFloatArray) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Target gets the scannable target for SQLFloat
func (x *SQLFloat) Target() interface{} {
	return x
}

// Target gets the scannable target for SQLFloatArray
func (x *SQLFloatArray) Target() interface{} {
	return pq.Array(&x.Value)
}

// Scan interface->Int
func (x *SQLFloat) Scan(data interface{}) error {
	switch data.(type) {
	case float64:
		v, ok := data.(float64)
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
