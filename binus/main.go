package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	linked_list "github.com/vldcreation/go-ressources/dataStructure/linkedlist"
)

// const
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var menus = []string{}

var list = linked_list.NewSingleLinkedListEmployee()

var menuHandler map[int]func()

func registerMenu(pos int, name string, fn func()) {
	if pos > len(menus) {
		menus = append(menus, make([]string, pos-len(menus))...)
	}
	menus[pos-1] = name
	menuHandler[pos] = fn
}

func init() {
	menuHandler = make(map[int]func())
	registerMenu(1, "Add new Employee", add)
	registerMenu(2, "Display Employee List", display)
	registerMenu(3, "Delete Employee", delete)
	registerMenu(4, "Delete All Employee", deleteAll)
	registerMenu(5, "Exit", exit)
}

func add() {
	emp := linked_list.Employee{}
	i := 0
	maks := 5
	for i < maks {
		println("Add new Employee\n")
		print("Employee ID:")
		fmt.Scanf("%d", &emp.EmployeeID)
		if emp.EmployeeID > 99999 {
			println(Red + "Employee ID maksimal 5 angka" + Reset)
			return
		}

		print("Full Name:")
		emp.FullName = scanln()
		if len(emp.FullName) > 30 {
			println(Red + "Full Name maksimal 30 karakter" + Reset)
			return
		}

		print("Place of Birth:")
		emp.PlaceOfBirth = scanln()
		if len(emp.PlaceOfBirth) > 30 {
			println(Red + "Place of Birth maksimal 30 karakter" + Reset)
			return
		}

		print("Date of Birth:")
		fmt.Scanln(&emp.DateOfBirth)
		if _, err := time.Parse("2006-01-02", emp.DateOfBirth); err != nil {
			println(Red + "Date of Birth format salah, format yang benar adalah yyyy-mm-dd" + Reset)
			print(err.Error())
			return
		}

		print("Position:")
		emp.Position = scanln()
		if len(emp.Position) > 50 {
			println(Red + "Position maksimal 50 karakter" + Reset)
			return
		}

		println(Green + "Success add new employee. press 1 to continue input or any key to top menu" + Reset)
		ch := ""
		fmt.Scanln(&ch)
		list.Append(emp)
		i++
		if ch != "1" {
			break
		}

	}
}

func display() {
	sortType := ""
	print("Sort by Employee ID (asc/desc):")
	fmt.Scanf("%s", &sortType)
	list.Sort(sortType)

	println("\n=====Employee List=====")
	list.Display()
}

func delete() {
	print("Employee ID:")
	var id int
	fmt.Scanf("%d", &id)
	emp := linked_list.Employee{EmployeeID: id} // delete by employee id
	list.Delete(emp)
}

func deleteAll() {
	list.DeleteAll()
}

func exit() {
	println("Thanks for using this program, bye!")
	os.Exit(0)
}

func menu() {
	for k, v := range menus {
		println(k+1, v)
	}
}

func main() {
	for {
		menu()
		var input int
		print("Choose menu:")
		fmt.Scanf("%d", &input)
		if fn, ok := menuHandler[input]; ok {
			fn()
		} else {
			println("Menu not found")
		}
	}
}

func scanln() string {
	text := ""
	scanner := bufio.NewScanner(os.Stdin)

	// Scan for the next line
	if scanner.Scan() {
		text = scanner.Text() // Get the text from the scanned line
	}

	// Check for any scanning error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}

	return text
}
