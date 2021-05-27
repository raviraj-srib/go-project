package service

import (
	"github.com/raviraj-srib/go-project/corporate-directory/src/logger"
	"github.com/raviraj-srib/go-project/corporate-directory/src/model"
)

func (service directoryServiceImpl) GetClosestCommonManager(empId1 string, empId2 string) string {
	logger.Error("Method not yet implemented")
	return ""

}

func (service directoryServiceImpl) AddEmployee(emp model.Node, mgrId string) {
	logger.Error("Method not yet implemented")
}

func (service directoryServiceImpl) GetEmployeeDetails(empId string) model.Node {
	return service.searchEmployee(empId)
}

func (service directoryServiceImpl) AssignReportee(mgrId string, reportees model.Node) {
	logger.Error("Method not yet implemented")

}

func (service directoryServiceImpl) AssignReportees(mgrId string, reportees []model.Node) {
	logger.Error("Method not yet implemented")

}

func (service directoryServiceImpl) MoveAllReporteeToOtherManager(curmgrId string, newmgrId string) {
	logger.Error("Method not yet implemented")
}

func (service directoryServiceImpl) ChangeManager(empId string, oldmgrId string, newmgrId string) {
	logger.Error("Method not yet implemented")
}

func (service directoryServiceImpl) PromoteToManager(empId string, reportees []model.Employee) {
	logger.Error("Method not yet implemented")
}

func (service directoryServiceImpl) RemoveEmployee(empId string) {
	logger.Error("Method not yet implemented")
}

func (service directoryServiceImpl) PrintEmployeeHierarchy(empId string) {
	emp := service.searchEmployee(empId)
	printTreeInBFS(emp)
}
