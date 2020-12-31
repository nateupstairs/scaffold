package scaffold

import (
	"errors"
	"time"
)

// CellType for vague match against SQL types
type CellType int

// CellType const
const (
	CellBool CellType = iota
	CellBoolArray
	CellString
	CellStringArray
	CellInt
	CellIntArray
	CellFloat
	CellFloatArray
	CellTime
)

type SQLBool struct {
	Valid bool
	Value bool
}

type SQLBoolArray struct {
	Value []bool
}

type SQLString struct {
	Valid bool
	Value string
}

type SQLStringArray struct {
	Value []string
}

type SQLInt struct {
	Valid bool
	Value int64
}

type SQLIntArray struct {
	Value []int64
}

type SQLFloat struct {
	Valid bool
	Value float64
}

type SQLFloatArray struct {
	Value []float64
}

type SQLTime struct {
	Valid bool
	Value time.Time
}

func NewSQLBoolArray() *SQLBoolArray {
	x := new(SQLBoolArray)
	x.Value = make([]bool, 0)

	return x
}

func NewSQLStringArray() *SQLStringArray {
	x := new(SQLStringArray)
	x.Value = make([]string, 0)

	return x
}

func NewSQLIntArray() *SQLIntArray {
	x := new(SQLIntArray)
	x.Value = make([]int64, 0)

	return x
}

func NewSQLFloatArray() *SQLFloatArray {
	x := new(SQLFloatArray)
	x.Value = make([]float64, 0)

	return x
}

func (x *SQLBool) Scan(data interface{}) error {
	switch data.(type) {
	case bool:
		v, ok := data.(bool)
		if ok {
			x.Valid = true
			x.Value = v
		}
	default:
		return errors.New("Incompatible type")
	}
	return nil
}

func (x *SQLString) Scan(data interface{}) error {
	switch data.(type) {
	case string:
		v, ok := data.(string)
		if ok {
			x.Valid = true
			x.Value = v
		}
	default:
		return errors.New("Incompatible type")
	}
	return nil
}

func (x *SQLInt) Scan(data interface{}) error {
	switch data.(type) {
	case int64:
		v, ok := data.(int64)
		if ok {
			x.Valid = true
			x.Value = v
		}
	default:
		return errors.New("Incompatible type")
	}
	return nil
}

func (x *SQLFloat) Scan(data interface{}) error {
	switch data.(type) {
	case float64:
		v, ok := data.(float64)
		if ok {
			x.Valid = true
			x.Value = v
		}
	default:
		return errors.New("Incompatible type")
	}
	return nil
}

func (x *SQLTime) Scan(data interface{}) error {
	switch data.(type) {
	case time.Time:
		v, ok := data.(time.Time)
		if ok {
			x.Valid = true
			x.Value = v
		}
	default:
		return errors.New("Incompatible type")
	}
	return nil
}

// Cell type container
type Cell struct {
	Name           string
	SQL            string
	Type           CellType
	Exclude        bool
	BoolVal        *SQLBool
	BoolArrayVal   *SQLBoolArray
	StringVal      *SQLString
	StringArrayVal *SQLStringArray
	IntVal         *SQLInt
	IntArrayVal    *SQLIntArray
	FloatVal       *SQLFloat
	FloatArrayVal  *SQLFloatArray
	TimeVal        *SQLTime
}

// GetValue from cell
func (c *Cell) GetValue() (interface{}, error) {
	switch c.Type {
	case CellBool:
		if c.BoolVal.Valid {
			return c.BoolVal.Value, nil
		}
		return c.BoolVal.Value, errors.New("Invalid value")
	case CellBoolArray:
		if c.BoolArrayVal != nil {
			return c.BoolArrayVal.Value, nil
		}
		return nil, errors.New("Invalid value")
	case CellString:
		if c.StringVal.Valid {
			return c.StringVal.Value, nil
		}
		return c.StringVal.Value, errors.New("Invalid value")
	case CellStringArray:
		if c.StringArrayVal != nil {
			return c.StringArrayVal.Value, nil
		}
		return nil, errors.New("Invalid value")
	case CellInt:
		if c.IntVal.Valid {
			return c.IntVal.Value, nil
		}
		return c.IntVal.Value, errors.New("Invalid value")
	case CellIntArray:
		if c.IntArrayVal != nil {
			return c.IntArrayVal.Value, nil
		}
		return nil, errors.New("Invalid value")
	case CellFloat:
		if c.FloatVal.Valid {
			return c.FloatVal.Value, nil
		}
		return c.FloatVal.Value, errors.New("Invalid value")
	case CellFloatArray:
		if c.FloatArrayVal != nil {
			return c.FloatArrayVal.Value, nil
		}
		return nil, errors.New("Invalid value")
	case CellTime:
		if c.TimeVal.Valid {
			return c.TimeVal.Value, nil
		}
		return c.TimeVal.Value, errors.New("Invalid value")
	}

	return nil, errors.New("Missing cell type")
}

// SetValue to cell
func (c *Cell) SetValue(val interface{}) error {
	return c.Scan(val)
}

// ScanValue to cell
func (c *Cell) Scan(val interface{}) error {
	switch c.Type {
	case CellBool:
		if c.BoolVal == nil {
			c.BoolVal = new(SQLBool)
		}

		c.BoolVal.Scan(val)
	case CellString:
		if c.BoolVal == nil {
			c.BoolVal = new(SQLBool)
		}

		c.StringVal = new(SQLString)
		c.StringVal.Scan(val)
	case CellInt:
		if c.IntVal == nil {
			c.IntVal = new(SQLInt)
		}

		c.IntVal.Scan(val)
	case CellFloat:
		if c.FloatVal == nil {
			c.FloatVal = new(SQLFloat)
		}

		c.FloatVal.Scan(val)
	case CellTime:
		if c.TimeVal == nil {
			c.TimeVal = new(SQLTime)
		}

		c.TimeVal.Scan(val)
	default:
		return errors.New("Incorrect scanning data type")
	}

	return nil
}
