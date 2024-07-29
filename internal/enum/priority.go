package enum

type Priority int

const (
	HighPriority Priority = iota + 1
	MediumPriority
	LowPriority
)

func (p Priority) PriorityString() string {
	return [...]string{"high", "medium", "low"}[p-1]
}
