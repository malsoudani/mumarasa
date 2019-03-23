package main

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next *Node
}

func (l *List) First() (*Node) {
	return l.head
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node // the head is assigned only if the linked list is empty
	} else {
		l.tail.next = node 	// there is an existing element in the list meaning that you have an element
	}						// that is pointed to by both the head and the tail or more than one element that has a head and a tail
							// here you are saying to keep the head in its place and move the tail to point to the next element aka the newly pushed on
	l.tail = node // notice that in a linkedList you always want to assign the tail regardless of weather the head is assigned
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	n := l.First()
	for {
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
		println(n.value)
	}
}