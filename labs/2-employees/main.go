package main

import (
	"errors"
	"fmt"
	"os"
)

type kindOfEmployee int

const (
	salary = iota
	hourly
)

type employee struct {
	kind kindOfEmployee
}

type hourlyEmployee struct {
	employee
	hourlyRate  int
	hoursWorked int
}

func (he *hourlyEmployee) setEmployeeType() {
	he.kind = hourly
}

func (he hourlyEmployee) calculate() int {
	return he.hourlyRate * he.hoursWorked
}

func (he hourlyEmployee) String() string {
	return fmt.Sprintf("Hourly Employee (Rate: %d, Hours Worked: %d)", he.hourlyRate, he.hoursWorked)
}

type salariedEmployee struct {
	employee
	annualSalary int
}

func (se *salariedEmployee) setEmployeeType() {
	se.kind = salary
}

func (se salariedEmployee) calculate() int {
	return se.annualSalary / 26
}

func (se salariedEmployee) String() string {
	return fmt.Sprintf("Salaried Employee (Annual Salary: %d)", se.annualSalary)
}

type iEmployee interface {
	calculate() int
}

func addEmployee(list []iEmployee, nums ...int) ([]iEmployee, error) {
	if len(nums) == 1 {
		e := salariedEmployee{annualSalary: nums[0]}
		e.setEmployeeType()
		list = append(list, e)
		return list, nil
	}

	if len(nums) == 2 {
		e := hourlyEmployee{hoursWorked: nums[0], hourlyRate: nums[1]}
		e.setEmployeeType()
		list = append(list, e)
		return list, nil
	}

	return nil, errors.New("more than 2 values included in variadic parameter")
}

func main() {
	var employees []iEmployee

	employees, err := addEmployee(employees, 30_000)
	exitOnError(err)

	employees, err = addEmployee(employees, 40, 40)
	exitOnError(err)

	employees, err = addEmployee(employees, 35, 30)
	exitOnError(err)

	func() {
		for _, e := range employees {
			fmt.Println(e)
		}
	}()

	fmt.Println("First employee payroll: ", calculatePayroll(employees[0]))
	fmt.Println("All employee payroll: ", calculatePayroll(employees...))
}

func calculatePayroll(employees ...iEmployee) int {
	payroll := 0
	for _, e := range employees {
		payroll += e.calculate()
	}
	return payroll
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
