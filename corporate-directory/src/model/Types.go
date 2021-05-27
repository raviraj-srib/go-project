package model

type Node interface {
	GetId() string
	Create(name string)
	IsManager() bool
}

type Engineer struct {
	Employee
}

type Manager struct {
	Employee
	reportees map[string]Node
}

type Employee struct {
	id   string
	info EmpInfo
}

type EmpInfo struct {
	name    string
	salary  Salary
	dob     Date
	doj     Date
	address Address
}

type Salary struct {
	base       float32
	hra        float32
	sAllowance float32
	epf        float32
	gradutiy   float32
	bonus      float32
}

type Date struct {
	day   int
	month int
	year  int
}

type Address struct {
	addLine1 string
	addLine2 string
	district string
	state    string
	country  string
	pinCode  string
}
