package service

import (
	"strings"

	"github.com/raviraj-srib/go-project/corporate-directory/src/logger"
	"github.com/raviraj-srib/go-project/corporate-directory/src/model"
)

func (service directoryServiceImpl) getClosetManagerUsingLCA(empId1 string, empId2 string) string {
	lca := getLCA(service.GetCeo(), empId1, empId2)
	if nil != lca {
		return lca.GetId()
	} else {
		logger.Error("Not able to find both employee in database, emp1: %s emp2: %s", empId1, empId2)
		return "-1"
	}
}

func (service directoryServiceImpl) getManagerFromEmpId(empId string) (*model.Manager, bool) {
	mgr := service.searchEmployee(empId)
	if mgr == nil || mgr.IsManager() {
		logger.Error("Manager not found with id: %s", empId)
		return nil, false
	} else {
		oldManager, ok := mgr.(*model.Manager)
		return oldManager, ok
	}
}

func (service directoryServiceImpl) searchEmployee(searchEmpId string) model.Node {
	ceo := service.GetCeo()
	return searchEmployeeDFS(ceo, searchEmpId)
}

func searchEmployeeDFS(curEmp model.Node, searchEmpId string) model.Node {
	if curEmp.GetId() == searchEmpId {
		return curEmp
	}

	if curEmp.IsManager() {
		mgr, _ := curEmp.(*model.Manager)
		reportees := mgr.GetReportee()

		emp, ok := reportees[searchEmpId]
		if ok {
			return emp
		}

		for _, emp := range reportees {
			foundEmp := searchEmployeeDFS(emp, searchEmpId)
			if nil != foundEmp {
				return foundEmp
			}
		}
	}
	return nil
}

func getLCA(curEmp model.Node, empId1 string, empId2 string) model.Node {
	curEmpId := curEmp.GetId()

	if curEmpId == empId1 || curEmpId == empId2 {
		return curEmp
	}

	var lca model.Node

	if curEmp.IsManager() {
		mgr, _ := curEmp.(*model.Manager)
		reportees := mgr.GetReportee()

		emp1, isEmp1Present := reportees[empId1]
		if isEmp1Present {
			return emp1
		}

		emp2, isEmp2Present := reportees[empId2]
		if isEmp2Present {
			return emp2
		}
		count := 0
		for _, emp := range reportees {
			temp := getLCA(emp, empId1, empId2)
			if nil != temp {
				lca = temp
				count++
			}
		}
		if count == 2 {
			return curEmp
		}
	}
	return lca
}

//TODO: Revisit for approach
func printCompleteData(emp model.Node) {
	if emp.IsManager() {
		manager, _ := emp.(*model.Manager)
		logger.Debug("ManagerId: %s  -->> Reportees: %s", manager.GetId(), printAllKeys(manager.GetReportee()))
		for _, reportee := range manager.GetReportee() {
			printCompleteData(reportee)
		}
	} else {
		logger.Debug("EngineerId: %s", emp.GetId())
	}

}

func printAllKeys(empMap map[string]model.Node) string {
	var data strings.Builder

	for key, _ := range empMap {
		data.WriteString(key + " ")
	}

	return data.String()
}
