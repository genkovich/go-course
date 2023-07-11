package class

import (
	"course/hw9/student"
	"course/hw9/teacher"
)

type Class struct {
	Title    string
	Teacher  teacher.Teacher
	Students []student.Student
}

type Statistic struct {
	Class      string  `json:"class"`
	AverageAge float64 `json:"average_age"`
	BoysCount  int     `json:"boys_count"`
	GirlsCount int     `json:"girls_count"`
}
