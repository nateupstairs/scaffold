package scaffold

import (
	"database/sql"
	"errors"
)

// CellType for vague match against SQL types
type CellType int

// CellType const
const (
	CellBool CellType = iota
	CellString
	CellInt
	CellFloat
	CellTime
)

// Cell type container
type Cell struct {
	Name      string
	SQL       string
	Type      CellType
	Exclude   bool
	BoolVal   *sql.NullBool
	StringVal *sql.NullString
	IntVal    *sql.NullInt64
	FloatVal  *sql.NullFloat64
	TimeVal   *sql.NullTime
}

// GetValue from cell
func (c *Cell) GetValue() (interface{}, error) {
	switch c.Type {
	case CellBool:
		if c.BoolVal == nil {
			return nil, errors.New("Missing cell value")
		}
		return c.BoolVal.Value()
	case CellString:
		if c.StringVal == nil {
			return nil, errors.New("Missing cell value")
		}
		return c.StringVal.Value()
	case CellInt:
		if c.IntVal == nil {
			return nil, errors.New("Missing cell value")
		}
		return c.IntVal.Value()
	case CellFloat:
		if c.FloatVal == nil {
			return nil, errors.New("Missing cell value")
		}
		return c.FloatVal.Value()
	case CellTime:
		if c.TimeVal == nil {
			return nil, errors.New("Missing cell value")
		}
		return c.TimeVal.Value()
	}

	return nil, errors.New("Missing cell type")
}
