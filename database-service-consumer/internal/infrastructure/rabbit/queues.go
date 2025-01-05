package rabbit

type Queue string

func (q Queue) String() string {
	return string(q)
}

const (
	PersonQueue Queue = "person"
)
