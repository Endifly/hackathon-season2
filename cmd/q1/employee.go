package main

import "time"

// Employee
type Employee struct {
	EmployeeID int
	Passport   string
	FirstName  string
	LastName   string
	Gender     int
	Birthday   time.Time
}
