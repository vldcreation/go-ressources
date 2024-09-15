package linked_list

import "fmt"

// Case1
// Push Employee data to Single Linked List
// Display Employee data
// Delete Employee data by Employee ID
// Sort Employee data by Employee ID
// Delete all Employee data

type Employee struct {
	EmployeeID   int
	FullName     string
	PlaceOfBirth string
	DateOfBirth  string
	Position     string
}

func (e Employee) String() string {
	format := "Employee ID: %d\nNama Lengkap: %s\nTempat Lahir: %s\nTanggal Lahir: %s\nJabatan: %s\n"
	return fmt.Sprintf(format, e.EmployeeID, e.FullName, e.PlaceOfBirth, e.DateOfBirth, e.Position)
}

type NodeEmployee struct {
	data Employee
	next *NodeEmployee
}

type singleLinkedListEmployee struct {
	head *NodeEmployee
}

type LinkedListEmployee interface {
	Source() singleLinkedListEmployee
	LinkedList
}

func NewSingleLinkedListEmployee() LinkedListEmployee {
	return &singleLinkedListEmployee{}
}

func (s *singleLinkedListEmployee) Source() singleLinkedListEmployee {
	return *s
}

func (s *singleLinkedListEmployee) Append(data any) {
	dataEmployee, ok := data.(Employee)
	if !ok {
		panic("Data yang diinputkan bukan Employee")
	}

	newNode := &NodeEmployee{data: dataEmployee}
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

func (s *singleLinkedListEmployee) Count() int {
	count := 0
	current := s.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (s *singleLinkedListEmployee) Delete(data any) {
	dataEmployee, ok := data.(Employee)
	if !ok {
		panic("Data yang diinputkan bukan Employee")
	}

	if s.head == nil {
		return
	}

	if s.head.data.EmployeeID == dataEmployee.EmployeeID {
		s.head = s.head.next
		return
	}

	current := s.head
	for current.next != nil {
		if current.next.data.EmployeeID == dataEmployee.EmployeeID {
			current.next = current.next.next
			return
		}

		current = current.next
	}
}

func (s *singleLinkedListEmployee) Display() {
	current := s.head
	for current != nil {
		println(current.data.String(), " \n")
		current = current.next
	}
}

func (s *singleLinkedListEmployee) Sort(sortType string) {
	if sortType == "asc" {
		s.sortAsc()
	} else if sortType == "desc" {
		s.sortDesc()
	} else {
		println("Invalid sort type, please use 'asc' or 'desc'. \n by default use asc then")
	}
}

func (s *singleLinkedListEmployee) sortAsc() {
	current := s.head
	for current != nil {
		next := current.next
		for next != nil {
			if current.data.EmployeeID > next.data.EmployeeID {
				temp := current.data
				current.data = next.data
				next.data = temp
			}
			next = next.next
		}
		current = current.next
	}
}

func (s *singleLinkedListEmployee) sortDesc() {
	current := s.head
	for current != nil {
		next := current.next
		for next != nil {
			if current.data.EmployeeID < next.data.EmployeeID {
				temp := current.data
				current.data = next.data
				next.data = temp
			}
			next = next.next
		}
		current = current.next
	}
}

func (s *singleLinkedListEmployee) DeleteAll() {
	s.head = nil
}
