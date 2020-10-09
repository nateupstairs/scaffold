package scaffold

// Filter a query
type Filter struct {
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
