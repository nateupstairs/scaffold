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
	CellDate
	CellDateArray
	CellDatetime
	CellDatetimeArray
	CellByte
	CellByteArray
)

// SQLCell explains that base properties of a cell
type SQLCell interface {
	Target() interface{}
	Raw() (interface{}, error)
}

// Cell type container
type Cell struct {
	Name    string
	SQL     string
	Type    CellType
	Exclude bool
	Data    SQLCell
}

// CellTarget for a cell
func (c *Cell) CellTarget() interface{} {
	return c.Data.Target()
}

// GetValue from cell
func (c *Cell) GetValue() (interface{}, error) {
	return c.Data.Raw()
}

// Bool from cell
func (c *Cell) Bool() (bool, error) {
	v, err := c.GetValue()
	if err != nil {
		return false, err
	}

	vv, ok := v.(bool)
	if !ok {
		return false, err
	}

	return vv, nil
}

// SetBool to cell
func (c *Cell) SetBool(x bool) error {
	if c.Type != CellBool {
		return errors.New("set incorrect type")
	}

	d := NewSQLBool()
	d.Scan(x)

	c.Data = d

	return nil
}

// BoolArray from cell
func (c *Cell) BoolArray() ([]bool, error) {
	v, err := c.GetValue()
	if err != nil {
		return []bool{}, err
	}

	vv, ok := v.([]bool)
	if !ok {
		return []bool{}, err
	}

	return vv, nil
}

// String from cell
func (c *Cell) String() (string, error) {
	v, err := c.GetValue()
	if err != nil {
		return "", err
	}

	vv, ok := v.(string)
	if !ok {
		return "", err
	}

	return vv, nil
}

// SetString to cell
func (c *Cell) SetString(x string) error {
	if c.Type != CellString {
		return errors.New("set incorrect type")
	}

	d := NewSQLString()
	d.Scan(x)

	c.Data = d

	return nil
}

// StringArray from cell
func (c *Cell) StringArray() ([]string, error) {
	v, err := c.GetValue()
	if err != nil {
		return []string{}, err
	}

	vv, ok := v.([]string)
	if !ok {
		return []string{}, err
	}

	return vv, nil
}

// Int from cell
func (c *Cell) Int() (int64, error) {
	v, err := c.GetValue()
	if err != nil {
		return 0, err
	}

	vv, ok := v.(int64)
	if !ok {
		return 0, err
	}

	return vv, nil
}

// SetInt to cell
func (c *Cell) SetInt(x int64) error {
	if c.Type != CellInt {
		return errors.New("set incorrect type")
	}

	d := NewSQLInt()
	d.Scan(x)

	c.Data = d

	return nil
}

// IntArray from cell
func (c *Cell) IntArray() ([]int64, error) {
	v, err := c.GetValue()
	if err != nil {
		return []int64{}, err
	}

	vv, ok := v.([]int64)
	if !ok {
		return []int64{}, err
	}

	return vv, nil
}

// Float from cell
func (c *Cell) Float() (float64, error) {
	v, err := c.GetValue()
	if err != nil {
		return 0, err
	}

	vv, ok := v.(float64)
	if !ok {
		return 0, err
	}

	return vv, nil
}

// SetFloat to cell
func (c *Cell) SetFloat(x float64) error {
	if c.Type != CellFloat {
		return errors.New("set incorrect type")
	}

	d := NewSQLFloat()
	d.Scan(x)

	c.Data = d

	return nil
}

// FloatArray from cell
func (c *Cell) FloatArray() ([]float64, error) {
	v, err := c.GetValue()
	if err != nil {
		return []float64{}, err
	}

	vv, ok := v.([]float64)
	if !ok {
		return []float64{}, err
	}

	return vv, nil
}

// Date from cell
func (c *Cell) Date() (time.Time, error) {
	v, err := c.GetValue()
	if err != nil {
		return time.Time{}, err
	}

	vv, ok := v.(time.Time)
	if !ok {
		return time.Time{}, err
	}

	return vv, nil
}

// SetDate to cell
func (c *Cell) SetDate(x time.Time) error {
	if c.Type != CellDate {
		return errors.New("set incorrect type")
	}

	d := NewSQLDate()
	d.Scan(x)

	c.Data = d

	return nil
}

// DateArray from cell
func (c *Cell) DateArray() ([]time.Time, error) {
	v, err := c.GetValue()
	if err != nil {
		return []time.Time{}, err
	}

	vv, ok := v.([]time.Time)
	if !ok {
		return []time.Time{}, err
	}

	return vv, nil
}

// Datetime from cell
func (c *Cell) Datetime() (time.Time, error) {
	v, err := c.GetValue()
	if err != nil {
		return time.Time{}, err
	}

	vv, ok := v.(time.Time)
	if !ok {
		return time.Time{}, err
	}

	return vv, nil
}

// SetDatetime to cell
func (c *Cell) SetDatetime(x time.Time) error {
	if c.Type != CellDate {
		return errors.New("set incorrect type")
	}

	d := NewSQLDatetime()
	d.Scan(x)

	c.Data = d

	return nil
}

// DatetimeArray from cell
func (c *Cell) DatetimeArray() ([]time.Time, error) {
	v, err := c.GetValue()
	if err != nil {
		return []time.Time{}, err
	}

	vv, ok := v.([]time.Time)
	if !ok {
		return []time.Time{}, err
	}

	return vv, nil
}
