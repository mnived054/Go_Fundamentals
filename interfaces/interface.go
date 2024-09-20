package main

import (
	"fmt"
)

type PermenantEmp struct { //struct 1 signature
	Name           string
	Age, Pay, Perk int
}

type ContractEmp struct { //struct 2 signature
	Name     string
	Age, Pay int
}

type EmployeeSalary interface {
	CaluculateSalary() int
}

func (pe *PermenantEmp) CaluculateSalary() int {
	return pe.Pay + pe.Perk
}

func (ce *ContractEmp) CaluculateSalary() int {
	return ce.Pay
}

func TotalExpense(Salary []EmployeeSalary) {
	var total = 0
	for _, emp := range Salary {
		total += emp.CaluculateSalary()
	}
	fmt.Println("Total expenses :", total)
}
func main() {
	var pe_1 = PermenantEmp{Name: "Nived", Age: 23, Perk: 1580, Pay: 50000}
	var pe_2 = PermenantEmp{Name: "Nikhil", Age: 25, Perk: 1650, Pay: 54200}
	var ce_1 = ContractEmp{Name: "Rahul", Age: 21, Pay: 30000}
	employee := []EmployeeSalary{&pe_1, &pe_2, &ce_1}
	TotalExpense(employee)
}
