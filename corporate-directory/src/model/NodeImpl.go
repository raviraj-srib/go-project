package model

import (
	"fmt"
	"strconv"
)

var employeeId = 0

func (m *Manager) Create(name string) {
	m.fillEmployeeDetails(name)
	m.reportees = make(map[string]Node)
}

func (m *Manager) GetId() string {
	return m.id
}

func (m *Manager) IsManager() bool {
	return true
}

func (m *Manager) AddReportee(node Node) {
	m.reportees[node.GetId()] = node
}

func (m *Manager) GetReportee() map[string]Node {
	return m.reportees
}

func (e *Engineer) Create(name string) {
	e.fillEmployeeDetails(name)
}

func (e *Engineer) GetId() string {
	return e.id
}

func (m *Engineer) IsManager() bool {
	return false
}

func (e *Employee) fillEmployeeDetails(name string) {
	e.id = e.generateEmployId()
	e.info = createEmployInfo(name)
}

//Generate EmployeeId across Organizaition
//TODO: generate unique employeeId with consideration
//to employee details
func (e *Employee) generateEmployId() string {
	employeeId++
	return strconv.Itoa(employeeId)
}

func (e *Employee) String() {
	fmt.Println("Employee {Id:", e.id, " Name: ", e.GetName(), "}")
}

func (e *Employee) GetName() string {
	return e.info.name

}

//Currently only name is being populated
func createEmployInfo(name string) EmpInfo {
	return EmpInfo{name: name}
}
