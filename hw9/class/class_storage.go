package class

import (
	"course/hw9/student"
	"course/hw9/teacher"
	"errors"
)

type Storage struct {
	classes []Class
}

type Provider interface {
	GetClassStatistic(class string) Statistic
}

func NewClassStorage() *Storage {
	studentsStorage := student.NewStudentStorage()
	return &Storage{
		classes: []Class{
			{
				Title: "A",
				Teacher: teacher.Teacher{
					Username: "teacher1",
					Password: "teacher1",
				},
				Students: studentsStorage.GetStudentsByClass("A"),
			},
			{
				Title: "B",
				Teacher: teacher.Teacher{
					Username: "teacher2",
					Password: "teacher2",
				},
				Students: studentsStorage.GetStudentsByClass("B"),
			},
			{
				Title: "C",
				Teacher: teacher.Teacher{
					Username: "teacher3",
					Password: "teacher3",
				},
				Students: studentsStorage.GetStudentsByClass("C"),
			},
		},
	}
}

func (s *Storage) GetClassByTitle(title string) Class {
	for _, class := range s.classes {
		if class.Title == title {
			return class
		}
	}
	return Class{}
}

func (s *Storage) GetTeacherByClass(classTitle string) teacher.Teacher {
	class := s.GetClassByTitle(classTitle)
	return class.Teacher
}

func (s *Storage) GetTeacherByUsername(username string) (teacher.Teacher, error) {
	for _, class := range s.classes {
		if class.Teacher.Username == username {
			return class.Teacher, nil
		}
	}
	return teacher.Teacher{}, errors.New("teacher not found")
}

func (s *Storage) IsTeacherResponsibility(classTitle string, teacherUsername string) bool {
	class := s.GetClassByTitle(classTitle)
	return class.Teacher.Username == teacherUsername
}

func (s *Storage) GetClassStatistic(classTitle string) Statistic {
	class := s.GetClassByTitle(classTitle)

	if len(class.Students) == 0 {
		return Statistic{Class: classTitle}
	}

	boysCount := 0
	girlsCount := 0
	averageAge := 0.0
	sumAge := 0
	for _, studentEntity := range class.Students {
		if studentEntity.Gender == "male" {
			boysCount++
		} else {
			girlsCount++
		}
		sumAge += studentEntity.Age
	}
	averageAge = float64(sumAge) / float64(len(class.Students))

	return Statistic{
		Class:      classTitle,
		AverageAge: averageAge,
		BoysCount:  boysCount,
		GirlsCount: girlsCount,
	}
}
