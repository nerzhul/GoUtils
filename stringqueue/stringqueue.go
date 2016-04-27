package stringqueue

// Imported and adapted from https://gist.github.com/moraes/2141121

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []string
	size  int
	head  int
	tail  int
	count int
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]string, size),
		size:  size,
	}
}

// Push adds a node to the queue.
func (q *Queue) Push(n string) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]string, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() string {
	if q.count == 0 {
		return ""
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}
