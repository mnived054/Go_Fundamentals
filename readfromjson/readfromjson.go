package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Address struct {
	Colony  string `json:"Colony"`
	Pincode int    `json:"Pincode"`
	Street  string `json:"Street"`
}

type Employee struct {
	Name       string    `json:"name"`
	Profession string    `json:"profession"`
	Address    []Address `json:"Address"`
}

type EmployeeMap struct {
	Employee1 []Employee `json:"Employee1"`
	Employee2 []Employee `json:"Employee2"`
}

type Employees struct {
	Employees []EmployeeMap `json:"Employees"`
}

func main() {

	data, err := ioutil.ReadFile("readfromjson/read.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var employees Employees
	err = json.Unmarshal(data, &employees)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, empMap := range employees.Employees {
		fmt.Println("Employee Map:")
		if len(empMap.Employee1) > 0 {
			fmt.Println("  Employee1:")
			for _, emp := range empMap.Employee1 {
				printEmployee(emp)
			}
		}
		if len(empMap.Employee2) > 0 {
			fmt.Println("  Employee2:")
			for _, emp := range empMap.Employee2 {
				printEmployee(emp)
			}
		}
	}

	var employeeName string
	fmt.Print("Enter the Employee name to fetch complete details:")
	fmt.Scan(&employeeName)
	found := printEmployeeByName(employees, employeeName)
	if !found {
		fmt.Printf("Employee with name %s not found.\n", employeeName)
	}

}

func printEmployee(emp Employee) {
	fmt.Printf("    Name: %s\n", emp.Name)
	fmt.Printf("    Profession: %s\n", emp.Profession)
	for _, addr := range emp.Address {
		fmt.Printf("      Address: %s, %s, %d\n\n", addr.Street, addr.Colony, addr.Pincode)
	}
}

func printEmployeeByName(employees Employees, name string) bool {
	for _, empMap := range employees.Employees {
		for _, empList := range []struct {
			Name string
			List []Employee
		}{
			{"Employee1", empMap.Employee1},
			{"Employee2", empMap.Employee2},
		} {
			for _, emp := range empList.List {
				if emp.Name == name {
					printEmployee(emp)
					return true
				}
			}
		}
	}
	return false
}
