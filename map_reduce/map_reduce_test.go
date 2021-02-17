package mapreduce

import (
	"fmt"
	"testing"
)

var employees = []*Employee{
	{"Levin", 28, 50000},
	{"Alice", 18, 25000},
	{"David", 35, 30000},
}

func TestMap(t *testing.T) {
	employeesSalaryIncrease := EmployeeMap(employees, func(e *Employee) *Employee {
		e.Salary = e.Salary + 10000
		return e
	})

	fmt.Println("increase salary for every employee:")
	for _, v := range employeesSalaryIncrease {
		fmt.Println(*v)
	}
}

func TestReduce(t *testing.T) {
	count := EmployeeReduceCount(employees, func(e *Employee) bool {
		return e.Age < 30
	})

	fmt.Printf("count of employees younger than 30: %d\n", count)

	sum := EmployeeReduceSum(employees, func(e *Employee) int {
		return e.Salary
	})

	fmt.Printf("sum of employees's salary: %d\n", sum)
}

func TestFilter(t *testing.T) {
	employeesYoungerThan30 := EmployeeFilter(employees, func(e *Employee) bool {
		return e.Age < 30
	})

	fmt.Println("filter employees that younger than 30:")
	for _, v := range employeesYoungerThan30 {
		fmt.Println(*v)
	}
}
