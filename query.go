package scaffold

type Filter struct {
	Field      string
	Comparison string
	Value      string
}

type Order struct {
	Field     string
	Direction string
}

type Query struct {
	Filters []Filter
	Orders  []Order
	Limit   int
	Offset  int
}
