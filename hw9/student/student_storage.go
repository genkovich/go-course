package student

import "errors"

type Student struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Class     string `json:"class"`
}

type ClassStatistic struct {
	Class      string  `json:"class"`
	AverageAge float64 `json:"average_age"`
	BoysCount  int     `json:"boys_count"`
	GirlsCount int     `json:"girls_count"`
}

type Provider interface {
	GetStudentsByClass(class string) []Student
	GetStudentById(id int) Student
	GetClassStatistic(class string) ClassStatistic
}

type Storage struct {
	students []Student
}

func NewStudentStorage() *Storage {
	return &Storage{
		students: []Student{
			{Id: 1, FirstName: "John", LastName: "Doe", Age: 14, Gender: "male", Class: "A"},
			{Id: 2, FirstName: "Anna", LastName: "Smith", Age: 15, Gender: "female", Class: "A"},
			{Id: 3, FirstName: "Paul", LastName: "Brown", Age: 14, Gender: "male", Class: "A"},
			{Id: 4, FirstName: "Laura", LastName: "Johnson", Age: 13, Gender: "female", Class: "B"},
			{Id: 5, FirstName: "Harry", LastName: "Miller", Age: 15, Gender: "male", Class: "B"},
			{Id: 6, FirstName: "Sophie", LastName: "Davis", Age: 13, Gender: "female", Class: "B"},
			{Id: 7, FirstName: "George", LastName: "Wilson", Age: 14, Gender: "male", Class: "B"},
			{Id: 8, FirstName: "Emma", LastName: "Taylor", Age: 15, Gender: "female", Class: "C"},
			{Id: 9, FirstName: "Charlie", LastName: "Anderson", Age: 13, Gender: "male", Class: "C"},
			{Id: 10, FirstName: "Mia", LastName: "Thomas", Age: 14, Gender: "female", Class: "C"},
			{Id: 11, FirstName: "Oliver", LastName: "Jackson", Age: 15, Gender: "male", Class: "C"},
			{Id: 12, FirstName: "Isabella", LastName: "White", Age: 13, Gender: "female", Class: "C"},
			{Id: 13, FirstName: "Max", LastName: "Harris", Age: 15, Gender: "male", Class: "C"},
			{Id: 14, FirstName: "Lily", LastName: "Martin", Age: 14, Gender: "female", Class: "C"},
			{Id: 15, FirstName: "Noah", LastName: "Thompson", Age: 14, Gender: "male", Class: "C"},
			{Id: 16, FirstName: "Ella", LastName: "Robinson", Age: 15, Gender: "female", Class: "C"},
		},
	}
}

func (s *Storage) GetStudentsByClass(class string) []Student {
	students := make([]Student, 0)
	for _, student := range s.students {
		if student.Class == class {
			students = append(students, student)
		}
	}
	return students
}

func (s *Storage) GetStudentById(id int) (Student, error) {
	// map is better, but I understand that after geretaing random students :)
	for _, student := range s.students {
		if student.Id == id {
			return student, nil
		}
	}
	return Student{}, errors.New("student not found")
}

func (s *Storage) GetClassStatistic(class string) ClassStatistic {
	students := s.GetStudentsByClass(class)

	if len(students) == 0 {
		return ClassStatistic{Class: class}
	}

	boysCount := 0
	girlsCount := 0
	averageAge := 0.0
	sumAge := 0
	for _, student := range students {
		if student.Gender == "male" {
			boysCount++
		} else {
			girlsCount++
		}
		sumAge += student.Age
	}
	averageAge = float64(sumAge) / float64(len(students))

	return ClassStatistic{
		Class:      class,
		AverageAge: averageAge,
		BoysCount:  boysCount,
		GirlsCount: girlsCount,
	}
}
