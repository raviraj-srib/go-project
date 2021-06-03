package service

import (
	"github.com/raviraj-srib/go-project/corporate-directory/src/logger"
	"github.com/raviraj-srib/go-project/corporate-directory/src/model"
)

func (service directoryServiceImpl) GetCeo() *model.Manager {
	return service.ceo
}

func (service directoryServiceImpl) GetClosestCommonManager(empId1 string, empId2 string) string {
	return service.getClosetManagerUsingLCA(empId1, empId2)
}

func (service directoryServiceImpl) AddEmployee(emp model.Node, mgrId string) {
	mgr, ok := service.getManagerFromEmpId(mgrId)
	if ok {
		mgr.AddReportee(emp)
		logger.Debug("Employee with id : %s added successfully to manager: %s", emp.GetId(), mgrId)
	} else {
		logger.Error("Manager not found with id: %s", mgrId)
	}
}

func (service directoryServiceImpl) GetEmployeeDetails(empId string) model.Node {
	return service.searchEmployee(empId)
}

func (service directoryServiceImpl) AssignReportees(mgrId string, reportee []model.Node) {
	mgr, ok := service.getManagerFromEmpId(mgrId)
	if ok {
		for _, emp := range reportee {
			mgr.AddReportee(emp)
			logger.Debug("Employee with id : %s added successfully to manager: %s", emp.GetId(), mgrId)
		}
	} else {
		logger.Error("Manager not found with id: %s", mgrId)
	}
}

func (service directoryServiceImpl) MoveAllReporteeToOtherManager(curmgrId string, newmgrId string) {
	curMgr, isCurManagerPresent := service.getManagerFromEmpId(curmgrId)
	newMgr, isNewManagerPresent := service.getManagerFromEmpId(newmgrId)
	if isCurManagerPresent && isNewManagerPresent {
		repotree := curMgr.GetReportee()
		for empId, emp := range repotree {
			curMgr.RemoveReportee(empId)
			newMgr.AddReportee(emp)
		}
	} else {
		logger.Error("Managers/Emp not present")
	}
}

func (service directoryServiceImpl) ChangeManager(empId string, oldmgrId string, newmgrId string) {
	oldMgr, isOldManagerPresent := service.getManagerFromEmpId(oldmgrId)
	newMgr, isNewManagerPresent := service.getManagerFromEmpId(newmgrId)
	emp, isEmployeePresent := service.getManagerFromEmpId(empId)
	if isOldManagerPresent && isNewManagerPresent && isEmployeePresent {
		oldMgr.RemoveReportee(empId)
		newMgr.AddReportee(emp)
	} else {
		logger.Error("Managers/Emp not present")
	}
}

func (service directoryServiceImpl) RemoveEmployee(empId string) {
	logger.Error("Method not yet implemented")
}

func (service directoryServiceImpl) PrintCompleteEmployeeHierarchy() {
	printCompleteData(service.ceo, 0)
}
func (service directoryServiceImpl) PrintEmployeeHierarchy(empId string) {
	emp := service.searchEmployee(empId)
	printCompleteData(emp, 0)
}
