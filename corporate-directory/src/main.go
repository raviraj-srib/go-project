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

}
