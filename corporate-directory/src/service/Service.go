package service

import "github.com/raviraj-srib/go-project/corporate-directory/src/model"

var dirSvcImpl directoryServiceImpl

func init() {
	ceo := &model.Manager{}
	ceo.Create("Claire")
	dirSvcImpl = directoryServiceImpl{ceo: ceo}
}

//TODO: Singleton implementation
func GetDirSvcImpl() directoryServiceImpl {
	return dirSvcImpl
}
