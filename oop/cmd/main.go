package main

import (
	"examples/oop/employee"
)

func main() {
	e := employee.Employee {
		FirstName: "Tobias",
		LastName: "Gleiter",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()

}