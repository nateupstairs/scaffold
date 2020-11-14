package scaffold

import (
	"bytes"
)

// Filter a query
type Filter struct {
	Operator   string
	Field      string
	Comparison string
	Value      string
}

// Order a query
type Order struct {
	Field     string
	Direction string
}

// Query structure to auto-build filter during sql query
type Query struct {
	Filters []Filter
	Orders  []Order
	Limit   int
	Offset  int
}

// ToBytes converts a query to a byte array text representation
func (q *Query) ToBytes() ([]byte, error) {
	var b bytes.Buffer

	templateVars := make(map[string]interface{}, 0)
	templateVars["query"] = q

	err := tmpl.ExecuteTemplate(&b, "query", templateVars)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
