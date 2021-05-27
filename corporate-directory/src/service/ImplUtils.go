package service

import "github.com/raviraj-srib/go-project/corporate-directory/src/model"

func (service directoryServiceImpl) GetCeo() *model.Manager {
	return service.ceo
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

func printTreeInBFS(emp model.Node) {

}
