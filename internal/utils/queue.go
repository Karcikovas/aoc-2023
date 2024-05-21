package utils

// TODO: Rework Queue implementation
// TODO: Add unit test to cover this
type Queue []interface{}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	h := *q
	var el interface{}
	l := len(h)
	el, *q = h[0], h[1:l]
	return el
}

func (q *Queue) Len() int {
	return len(*q)
}

func NewQueue() *Queue {
	return &Queue{}
}
