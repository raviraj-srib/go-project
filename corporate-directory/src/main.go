package main

import (
	"fmt"

	"github.com/raviraj-srib/go-project/corporate-directory/src/service"
	servicetest "github.com/raviraj-srib/go-project/corporate-directory/src/service_test"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("\n************DIRECTORY APPLICATION START************\n")

	logrus.Errorf("Name %s Address %s", "ravi", "12155")

	//	app.StartApplication()
	servicetest.GenerateDummyEmployees()
	dirService := service.GetDirectoryService()
	ceo := dirService.GetCeo()
	dirService.PrintEmployeeHierarchy(ceo.GetId())

	fmt.Println("\n************DIRECTORY APPLICATION END************\n")

}
