package service

import "github.com/raviraj-srib/go-project/corporate-directory/src/model"

var dirService DirectoryService

func init() {
	ceo := &model.Manager{}
	ceo.Create("Claire")
	dirService = directoryServiceImpl{ceo: ceo}
}

//TODO: Singleton implementation
func GetDirectoryService() DirectoryService {
	return dirService
}
