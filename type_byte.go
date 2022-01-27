package scaffold

import (
	"errors"

	"github.com/lib/pq"
)

// SQLBytes representation of SQL
type SQLBytes struct {
	Valid bool
	Value []byte
}

// SQLBytesArray representation of SQL
type SQLBytesArray struct {
	Valid bool
	Value [][]byte
}

// NewSQLByte makes a SQLByte
func NewSQLBytes() *SQLBytes {
	x := new(SQLBytes)
	x.Valid = false

	return x
}

// NewSQLByteArray makes a SQLByteArray
func NewSQLByteArray() *SQLBytesArray {
	x := new(SQLBytesArray)
	x.Valid = true
	x.Value = make([][]byte, 0)

	return x
}

// Raw Byte->Raw
func (x *SQLBytes) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Raw BytesArray->Raw
func (x *SQLBytesArray) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Target gets the scannable target for SQLByte
func (x *SQLBytes) Target() interface{} {
	return x
}

// Target gets the scannable target for SQLByteArray
func (x *SQLBytesArray) Target() interface{} {
	return pq.Array(&x.Value)
}

// Scan interface->Byte
func (x *SQLBytes) Scan(data interface{}) error {
	switch data.(type) {
	case []byte:
		v, ok := data.([]byte)
		if ok {
			x.Valid = true
			x.Value = v
		}
	case nil:
		x.Valid = false
		x.Value = []byte("")
	default:
		return errors.New("Incompatible type")
	}
	return nil
}
