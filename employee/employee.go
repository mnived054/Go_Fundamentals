package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	ID   int
	Name string
	Age  int
}

func fetchEmployeeDetail(id int, wg *sync.WaitGroup, ch chan<- Employee) {
	defer wg.Done()

	employee := Employee{ID: id, Name: fmt.Sprintf("Employee %d", id), Age: 30 + (id % 10)}

	ch <- employee
}

func fetchAllEmployeeDetails(n int) []Employee {
	var wg sync.WaitGroup
	ch := make(chan Employee, n)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go fetchEmployeeDetail(i, &wg, ch)
	}

	wg.Wait()
	close(ch)

	employees := make([]Employee, 0, n)
	for emp := range ch {
		employees = append(employees, emp)
	}

	return employees
}

func main() {

	employees := fetchAllEmployeeDetails(100)

	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", emp.ID, emp.Name, emp.Age)
	}
}
