package mapreduce

// Employee ...
type Employee struct {
	Name   string
	Age    uint
	Salary int
}

// EmployeeMap ...
func EmployeeMap(employees []*Employee, fn func(e *Employee) *Employee) []*Employee {
	newEmployees := make([]*Employee, len(employees))
	for i, v := range employees {
		newEmployees[i] = fn(v)
	}

	return newEmployees
}

// EmployeeReduceCount ...
func EmployeeReduceCount(employees []*Employee, fn func(e *Employee) bool) int {
	count := 0
	for _, v := range employees {
		if fn(v) {
			count++
		}
	}

	return count
}

// EmployeeReduceSum ...
func EmployeeReduceSum(employees []*Employee, fn func(e *Employee) int) int {
	var sum = 0
	for _, v := range employees {
		sum += fn(v)
	}

	return sum
}

// EmployeeFilter ...
func EmployeeFilter(employees []*Employee, fn func(e *Employee) bool) []*Employee {
	newEmployees := make([]*Employee, 0)
	for _, v := range employees {
		if fn(v) {
			newEmployees = append(newEmployees, v)
		}
	}

	return newEmployees
}
