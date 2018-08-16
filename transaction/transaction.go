package transaction


const (
	ADD = iota
	SUBTRACT
)

type Transaction struct {
	Type uint64
	Value int
}