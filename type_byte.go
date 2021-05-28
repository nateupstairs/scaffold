package scaffold

import (
	"errors"

	"github.com/lib/pq"
)

// SQLString representation of SQL
type SQLByte struct {
	Valid bool
	Value []byte
}

// SQLStringArray representation of SQL
type SQLByteArray struct {
	Valid bool
	Value [][]byte
}

// NewSQLByte makes a SQLByte
func NewSQLByte() *SQLByte {
	x := new(SQLByte)
	x.Valid = false

	return x
}

// NewSQLByteArray makes a SQLByteArray
func NewSQLByteArray() *SQLByteArray {
	x := new(SQLByteArray)
	x.Valid = true
	x.Value = make([][]byte, 0)

	return x
}

// Raw Byte->Raw
func (x *SQLByte) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Raw ByteArray->Raw
func (x *SQLByteArray) Raw() (interface{}, error) {
	if !x.Valid {
		return x.Value, errors.New("Invalid value")
	}

	return x.Value, nil
}

// Target gets the scannable target for SQLByte
func (x *SQLByte) Target() interface{} {
	return x
}

// Target gets the scannable target for SQLByteArray
func (x *SQLByteArray) Target() interface{} {
	return pq.Array(&x.Value)
}

// Scan interface->Byte
func (x *SQLByte) Scan(data interface{}) error {
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
