package servicetest

import (
	"testing"

	"github.com/raviraj-srib/go-project/corporate-directory/src/service"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	emp1 string
	emp2 string
	mgr  string
}

type Tests []TestCase

var tests = Tests{
	TestCase{emp1: "2", emp2: "3", mgr: "1"},
	TestCase{emp1: "2", emp2: "3", mgr: "1"},
	TestCase{emp1: "2", emp2: "3", mgr: "1"},
	TestCase{emp1: "2", emp2: "3", mgr: "1"},
	TestCase{emp1: "2", emp2: "3", mgr: "1"},
}

const (
	INPUT_FILE_PATH = "../../input/employee.txt"
)

func TestCommonClosestManager(t *testing.T) {
	generateDummyEmployees(INPUT_FILE_PATH)
	dirService := service.GetDirectoryService()
	dirService.PrintCompleteEmployeeHierarchy()

	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(dirService.GetClosestCommonManager(test.emp1, test.emp2), test.mgr)
	}

}
