package linked_list

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestNewSingleLinkedList(t *testing.T) {
	list := NewSingleLinkedList()
	if list == nil {
		t.Error("NewSingleLinkedList() should return a pointer to SingleLinkedList")
	}

	if list.Source().head != nil {
		t.Error("NewSingleLinkedList() should return an empty SingleLinkedList")
	}

	list.Append(1)
	if list.Source().head == nil {
		t.Error("Append() should add a new node to the list")
	}

	if list.Source().head.data != 1 {
		t.Error("Append() should add a new node with the correct data")
	}

	if list.Count() != 1 {
		t.Error("Append() should increment the count")
	}

	list.Append(2)
	if list.Source().head.next == nil {
		t.Error("Append() should add a new node to the list")
	}

	if list.Source().head.next.data != 2 {
		t.Error("Append() should add a new node with the correct data")
	}

	if list.Count() != 2 {
		t.Error("Append() should increment the count")
	}

	list.Delete(1)
	if list.Source().head.data != 2 {
		t.Error("Delete() should remove the correct node")
	}

	list.Display()
}

func TestNewCircularLinkedList(t *testing.T) {
	list := NewCircularLinkedList()
	if list == nil {
		t.Error("NewCircularLinkedList() should return a pointer to CircularLinkedList")
	}

	if list.Source().head != nil {
		t.Error("NewCircularLinkedList() should return an empty CircularLinkedList")
	}

	list.Append(1)
	if list.Source().head == nil {
		t.Error("Append() should add a new node to the list")
	}

	if list.Source().head.data != 1 {
		t.Error("Append() should add a new node with the correct data")
	}

	if list.Count() != 1 {
		t.Error("Append() should increment the count")
	}

	list.Append(2)
	if list.Source().head.next == nil {
		t.Error("Append() should add a new node to the list")
	}

	if list.Source().head.next.data != 2 {
		t.Error("Append() should add a new node with the correct data")
	}

	if list.Count() != 2 {
		t.Error("Append() should increment the count")
	}

	list.Delete(1)
	if list.Source().head.data != 2 {
		t.Error("Delete() should remove the correct node")
	}

	list.Display()
}

func TestSingleListEmployee(t *testing.T) {
	list := NewSingleLinkedListEmployee()
	if list == nil {
		t.Error("NewSingleLinkedListEmployee() should return a pointer to SingleLinkedListEmployee")
	}

	if list.Source().head != nil {
		t.Error("NewSingleLinkedListEmployee() should return an empty SingleLinkedListEmployee")
	}

	list.Append(Employee{EmployeeID: 1, FullName: "Employee 1"})
	if list.Source().head == nil {
		t.Error("Append() should add a new node to the list")
	}

	if list.Source().head.data.EmployeeID != 1 {
		t.Error("Append() should add a new node with the correct data")
	}

	if list.Count() != 1 {
		t.Error("Append() should increment the count")
	}

	list.Append(Employee{EmployeeID: 2, FullName: "Employee 2"})
	if list.Source().head.next == nil {
		t.Error("Append() should add a new node to the list")
	}

	if list.Source().head.next.data.EmployeeID != 2 {
		t.Error("Append() should add a new node with the correct data")
	}

	if list.Count() != 2 {
		t.Error("Append() should increment the count")
	}

	list.Sort("desc")
	list.Display()
	// list.Delete(Employee{EmployeeID: 1, FullName: "Employee 1"})
	// if list.Source().head.data.EmployeeID != 2 {
	// 	t.Error("Delete() should remove the correct node")
	// }

	// list.Display()

	// list.DeleteAll()
	// list.Delete(Employee{EmployeeID: 1})
	list.Sort("asc")
	list.Display()
}
