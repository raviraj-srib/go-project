package dataops

import (
	"math/rand"

	"github.com/raviraj-srib/go-project/corporate-directory/src/fileops"
	"github.com/raviraj-srib/go-project/corporate-directory/src/logger"
	"github.com/raviraj-srib/go-project/corporate-directory/src/model"
	"github.com/raviraj-srib/go-project/corporate-directory/src/service"
	"github.com/raviraj-srib/go-project/corporate-directory/src/utils"
)

const (
	MIN_MUL_FACTOR = 3
	MAX_MUL_FACTOR = 7
)

/*
Currently read the names from the file
distributes the employee based on below strategy
number of total employee at current level = (level * mulFactor)
number of engineers at current level = totalNumberOfEmployeeAtCurrentLevel * level * mulfactor

mulFactor is varying at each level from MIN_MUL_FACTOR to MAX_MUL_FACTOR
*/

//TODO: Refactor this method
func PopulateEmployeeData() {
	curEmpIndex := 0
	curLevel := 1

	queue := utils.Queue{}
	allEmpNames := fileops.ReadFile()
	totalEmployeeCount := len(allEmpNames)
	isDataExhausted := false

	queue.Add(service.GetDirSvcImpl().GetCeo(), curLevel)

	for !queue.IsEmpty() {
		data := queue.Remove()
		manager := data.GetManagerField()

		mulFactor := rand.Intn(MAX_MUL_FACTOR-MIN_MUL_FACTOR) + MIN_MUL_FACTOR
		empCount := data.GetLevel() * mulFactor

		//14% of 14 emp
		engCount := data.GetLevel()
		mgrCount := empCount - engCount
		logger.Debug("Id: %s CurLevel: %d MulFactor: %d  EmpCount:  %d  MgrCount:  %d EngCount: %d", manager.GetId(), data.GetLevel(), mulFactor, empCount, mgrCount, engCount)
		curLevel++
		for ; mgrCount > 0; mgrCount-- {
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

		for ; engCount > 0; engCount-- {
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
