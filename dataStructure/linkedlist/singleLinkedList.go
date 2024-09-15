package linked_list

import "fmt"

type Node struct {
	data any
	next *Node
}

type singleLinkedList struct {
	head *Node
}

type circularLinkedList struct {
	head *Node
}

type LinkedList interface {
	Append(data any)
	Display()
	Delete(data any)
	Count() int
	Sort(typeSort string)
	DeleteAll()
}

type SingleLinkedList interface {
	Source() singleLinkedList
	LinkedList
}

type CircularLinkedList interface {
	Source() circularLinkedList
	LinkedList
}

func NewSingleLinkedList() SingleLinkedList {
	return &singleLinkedList{}
}

func (s *singleLinkedList) Source() singleLinkedList {
	return *s
}

func (s *singleLinkedList) Append(data any) {
	newNode := &Node{data: data}
	if s.head == nil {
		s.head = newNode
		return
	}

	last := s.head
	for last.next != nil {
		last = last.next
	}

	last.next = newNode
}

func (s *singleLinkedList) Display() {
	current := s.head
	for current != nil {
		println(current.data, " -> ")
		current = current.next
	}
}

func (s *singleLinkedList) Delete(data any) {
	if s.head == nil {
		return
	}

	if s.head.data == data {
		s.head = s.head.next
		return
	}

	current := s.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}

		current = current.next
	}
}

func (s *singleLinkedList) Count() int {
	return s.count()
}

func (s *singleLinkedList) count() int {
	count := 0
	current := s.head
	for current != nil {
		count++
		current = current.next
	}

	return count
}

func (c *singleLinkedList) Sort(typeSort string) {
	fmt.Println("Noop")
}

func (s *singleLinkedList) DeleteAll() {
	s.head = nil
}

func NewCircularLinkedList() CircularLinkedList {
	return &circularLinkedList{}
}

func (c *circularLinkedList) Source() circularLinkedList {
	return *c
}

func (c *circularLinkedList) Append(data any) {
	newNode := &Node{data: data}
	if c.head == nil {
		c.head = newNode
		c.head.next = c.head
		return
	}

	last := c.head
	for last.next != c.head {
		last = last.next
	}

	last.next = newNode
	newNode.next = c.head
}

func (c *circularLinkedList) Display() {
	current := c.head
	for current != nil {
		println(current.data, " -> ")
		current = current.next
		if current == c.head {
			break
		}
	}
}

func (c *circularLinkedList) Delete(data any) {
	if c.head == nil {
		return
	}

	if c.head.data == data {
		c.head = c.head.next
		return
	}

	current := c.head
	for current.next != c.head {
		if current.next.data == data {
			current.next = current.next.next
			return
		}

		current = current.next
	}
}

func (c *circularLinkedList) Count() int {
	return c.count()
}

func (c *circularLinkedList) count() int {
	count := 0
	current := c.head
	for current != nil {
		count++
		current = current.next
		if current == c.head {
			break
		}
	}

	return count
}

func (c *circularLinkedList) Sort(typeSort string) {
	fmt.Println("Noop")
}

func (c *circularLinkedList) DeleteAll() {
	c.head = nil
}
