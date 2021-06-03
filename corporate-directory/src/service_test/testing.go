package servicetest

import (
	"github.com/raviraj-srib/go-project/corporate-directory/src/fileops"
	"github.com/raviraj-srib/go-project/corporate-directory/src/logger"
	"github.com/raviraj-srib/go-project/corporate-directory/src/model"
	"github.com/raviraj-srib/go-project/corporate-directory/src/service"
)

func GenerateDummyEmployees(path string) {
	generateDummyEmployees(path)
}

func generateDummyEmployees(path string) {
	allEmpNames := fileops.ReadFile(path)
	empIndex := 0
	empName := allEmpNames[empIndex]
	serviceImpl := service.GetDirectoryService()
	ceoId := serviceImpl.GetCeo().GetId()
	logger.Debug("Ceo Details %s", ceoId)

	//Adding 2 -->> mgr -->> ceo
	mgr := &model.Manager{}
	mgr.Create(empName)
	serviceImpl.AddEmployee(mgr, ceoId)

	//Adding 3 -->> mgr -->> ceo
	mgr = &model.Manager{}
	mgr.Create(empName)
	serviceImpl.AddEmployee(mgr, ceoId)

	//Adding 4 -->> mgr -->> ceo
	mgr = &model.Manager{}
	mgr.Create(empName)
	serviceImpl.AddEmployee(mgr, ceoId)

	//Adding 5 -->> mgr -->> 2
	mgr = &model.Manager{}
	mgr.Create(empName)
	serviceImpl.AddEmployee(mgr, "2")

	//Adding 6 -->> mgr -->> 2
	emp := &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "2")

	//Adding 7 -->> mgr -->> 3
	mgr = &model.Manager{}
	mgr.Create(empName)
	serviceImpl.AddEmployee(mgr, "3")

	//Adding 8 -->> mgr -->> 4
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "4")

	//Adding 9 -->> mgr -->> 4
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "4")

	//Adding 10 -->> mgr -->> 4
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "4")

	//Adding 11 -->> mgr -->> 5
	mgr = &model.Manager{}
	mgr.Create(empName)
	serviceImpl.AddEmployee(mgr, "5")

	//Adding 12 -->> mgr -->> 11
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "11")

	//Adding 13 -->> mgr -->> 11
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "11")

	//Adding 14 -->> mgr -->> 7
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "7")

	//Adding 15 -->> mgr -->> 7
	mgr = &model.Manager{}
	mgr.Create(empName)
	serviceImpl.AddEmployee(mgr, "7")

	//Adding 16 -->> mgr -->> 15
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "15")

	//Adding 17 -->> mgr -->> 15
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "15")

	//Adding 18 -->> mgr -->> 15
	emp = &model.Engineer{}
	emp.Create(empName)
	serviceImpl.AddEmployee(emp, "15")

}
