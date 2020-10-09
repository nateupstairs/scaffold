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
}

// GetValue from cell
func (c *Cell) GetValue() (interface{}, error) {
	switch c.Type {
	case CellBool:
		return c.BoolVal.Value()
	case CellString:
		return c.StringVal.Value()
	case CellInt:
		return c.IntVal.Value()
	case CellFloat:
		return c.FloatVal.Value()
	}

	return nil, errors.New("Missing cell type")
}
