package main

import (
	"fmt"

	"github.com/raviraj-srib/go-project/corporate-directory/src/service"
	servicetest "github.com/raviraj-srib/go-project/corporate-directory/src/service_test"
)

const (
	INPUT_FILE_PATH = "../input/employee.txt"
)

func main() {
	fmt.Println("\n************DIRECTORY APPLICATION START************\n")

	//	app.StartApplication()
	servicetest.GenerateDummyEmployees(INPUT_FILE_PATH)
	dirService := service.GetDirectoryService()
	ceo := dirService.GetCeo()
	dirService.PrintEmployeeHierarchy(ceo.GetId())

	fmt.Println("\n************DIRECTORY APPLICATION END************\n")

}
