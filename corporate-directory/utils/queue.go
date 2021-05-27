package utils

import "github.com/raviraj-srib/go-project/corporate-directory/src/model"

func (d Data) GetManagerField() *model.Manager {
	return d.mgr
}

func (d Data) GetLevel() int {
	return d.level
}
func (q *Queue) Add(mgr *model.Manager, level int) {
	q.list = append(q.list, Data{mgr: mgr, level: level})

}

func (q *Queue) Remove() Data {
	data := q.list[0]
	q.list = q.list[1:]
	return data
}

func (q *Queue) IsEmpty() bool {
	return len(q.list) == 0
}
