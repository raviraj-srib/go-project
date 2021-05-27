package main

import (
	"fmt"

	"github.com/raviraj-srib/go-project/corporate-directory/src/controller"
	"github.com/raviraj-srib/go-project/corporate-directory/src/service"
)

func main() {
	fmt.Println("************DIRECTORY APPLICATION************")
	controller.StartApplication()

	dirSvcImpl := service.GetDirSvcImpl()
	ceo := dirSvcImpl.GetCeo()
	dirSvcImpl.PrintEmployeeHierarchy(ceo.GetId())
	testFeature(dirSvcImpl)

}

func testFeature(dirSvcImpl service.DirectoryService) {
	fmt.Println("Closet Manager for Employees 10 and 18 is :", dirSvcImpl.GetClosestCommonManager("10", "18"))
	fmt.Println("Closet Manager for Employees 70 and 89 is :", dirSvcImpl.GetClosestCommonManager("70", "89"))
}
