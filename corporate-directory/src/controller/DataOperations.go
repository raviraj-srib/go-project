package controller

import (
	"math/rand"

	"github.com/raviraj-srib/go-project/corporate-directory/src/fileops"
	"github.com/raviraj-srib/go-project/corporate-directory/src/model"
	"github.com/raviraj-srib/go-project/corporate-directory/src/service"
	"github.com/raviraj-srib/go-project/corporate-directory/src/utils"
)

const (
	MIN_MUL_FACTOR = 7
	MAX_MUL_FACTOR = 12
)

/*
Currently read the names from the file
distributes the employee based on below strategy
number of total employee at current level = (level * mulFactor)
number of engineers at current level = totalNumberOfEmployeeAtCurrentLevel * level * mulfactor

mulFactor is varying at each level from MIN_MUL_FACTOR to MAX_MUL_FACTOR
*/

//TODO: Refactor this method
func populateEmployeeData() {
	curEmpIndex := 0
	curLevel := 1

	queue := utils.Queue{}
	allEmpNames := fileops.ReadFile()
	totalEmployeeCount := len(allEmpNames)
	isDataExhausted := false

	queue.Add(service.GetDirSvcImpl().GetCeo(), curLevel)

	for !queue.IsEmpty() {
		data := queue.Remove()
		totalEmployeeAtCurrentLevel := data.GetLevel()*rand.Intn(MIN_MUL_FACTOR-MIN_MUL_FACTOR) + MIN_MUL_FACTOR
		totalEngineerCount := int(totalEmployeeAtCurrentLevel * totalEmployeeAtCurrentLevel / 100)
		totalManagerCount := totalEmployeeAtCurrentLevel - totalEngineerCount
		curLevel++

		manager := data.GetManagerField()

		for ; totalManagerCount > 0; totalManagerCount-- {
			newMgr := &model.Manager{}
			curEmpIndex++
			if curEmpIndex >= totalEmployeeCount {
				isDataExhausted = true
				break
			}
			newMgr.Create(allEmpNames[curEmpIndex])
			manager.AddReportee(newMgr)
			queue.Add(newMgr, curLevel)
		}

		if isDataExhausted {
			break
		}

		for ; totalEngineerCount > 0; totalEngineerCount-- {
			eng := &model.Engineer{}
			curEmpIndex++
			if curEmpIndex >= totalEmployeeCount {
				isDataExhausted = true
				break
			}
			eng.Create(allEmpNames[curEmpIndex])
			manager.AddReportee(eng)
		}

		if isDataExhausted {
			break
		}
	}
}
