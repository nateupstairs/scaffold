package scaffold

import "time"

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
