package main

import (
	"fmt"

	"github.com/raviraj-srib/go-project/corporate-directory/src/app"
	"github.com/raviraj-srib/go-project/corporate-directory/src/service"
)

func main() {
	fmt.Println("\n************DIRECTORY APPLICATION START************\n")

	app.StartApplication()

	dirSvcImpl := service.GetDirSvcImpl()
	ceo := dirSvcImpl.GetCeo()
	dirSvcImpl.PrintEmployeeHierarchy(ceo.GetId())
	testFeature(dirSvcImpl)

	fmt.Println("\n************DIRECTORY APPLICATION END************\n")

}

func testFeature(dirSvcImpl service.DirectoryService) {
	fmt.Println("Closet Manager for Employees 10 and 18 is :", dirSvcImpl.GetClosestCommonManager("10", "18"))
	fmt.Println("Closet Manager for Employees 70 and 89 is :", dirSvcImpl.GetClosestCommonManager("70", "89"))
}
