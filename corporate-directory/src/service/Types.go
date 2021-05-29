package service

import (
	"github.com/raviraj-srib/go-project/corporate-directory/src/model"
)

type DirectoryService interface {
	// Takes 2 employee id as input and return their nearest manager id in heirarchy
	GetClosestCommonManager(employeeId1 string, employeeId2 string) string

	//Add new onboarded employee to the existing manager
	AddEmployee(employee model.Node, managerId string)

	// Returns employee Details given its employee id
	GetEmployeeDetails(employeeId string) model.Node

	//Assign set of employees to manager
	AssignReportees(managerId string, reportees []model.Node)

	//move all the reporties to another manager, mostly incase of manager leaving the organisation
	MoveAllReporteeToOtherManager(curManagerId string, newManagerId string)

	//change employee manager, mostly incase of project/team change
	ChangeManager(employeeId string, oldManagerId string, newManagerId string)

	//Remove the employee from directory, incase of leaving the organisation
	RemoveEmployee(employeeId string)

	//Print the details of employee recursively starting at given employeeid
	PrintEmployeeHierarchy(employeeId string)
}

type directoryServiceImpl struct {
	ceo *model.Manager
}
