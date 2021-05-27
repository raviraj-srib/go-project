package utils

import "github.com/raviraj-srib/go-project/corporate-directory/src/model"

type Queue struct {
	list []Data
}

type Data struct {
	mgr   *model.Manager
	level int
}
