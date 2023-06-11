package mark_journal

import "fmt"

type Journal struct {
	Subjects []Subject
}

type Subject struct {
	Title string
	Marks []StudentMarks
}

type StudentMarks struct {
	Student
	Marks []float64
}

type Student struct {
	Name    string
	Surname string
}

func JournalStart() {
	journal := Journal{
		Subjects: []Subject{
			{
				Title: "Math",
				Marks: []StudentMarks{
					{
						Student: Student{
							Name:    "John",
							Surname: "Doe",
						},
						Marks: []float64{5, 2.2, 5, 4.3},
					},
					{
						Student: Student{
							Name:    "Alex",
							Surname: "Smith",
						},
						Marks: []float64{4.5, 3.2, 4.1, 3.9},
					},
					{
						Student: Student{
							Name:    "Anna",
							Surname: "Brown",
						},
						Marks: []float64{4.0, 4.0, 4.5, 4.6},
					},
				},
			},
			{
				Title: "English",
				Marks: []StudentMarks{
					{
						Student: Student{
							Name:    "John",
							Surname: "Doe",
						},
						Marks: []float64{4.2, 3.2, 4.0, 4.1},
					},
					{
						Student: Student{
							Name:    "Alex",
							Surname: "Smith",
						},
						Marks: []float64{4.0, 3.0, 4.5, 3.8},
					},
					{
						Student: Student{
							Name:    "Anna",
							Surname: "Brown",
						},
						Marks: []float64{5.0, 4.5, 4.8, 4.9},
					},
				},
			},
			{
				Title: "Physics",
				Marks: []StudentMarks{
					{
						Student: Student{
							Name:    "John",
							Surname: "Doe",
						},
						Marks: []float64{3.5, 3.2, 3.7, 3.6},
					},
					{
						Student: Student{
							Name:    "Alex",
							Surname: "Smith",
						},
						Marks: []float64{4.5, 3.2, 4.1, 3.9},
					},
					{
						Student: Student{
							Name:    "Anna",
							Surname: "Brown",
						},
						Marks: []float64{4.0, 4.0, 4.5, 4.6},
					},
				},
			},
		},
	}

	fmt.Println("Average mark of Math", averageSubjectMark(journal.Subjects[0].Marks))
	fmt.Println("Average mark of English", averageSubjectMark(journal.Subjects[1].Marks))

	fmt.Println("Average mark John Doe by Math", averageMark(journal.Subjects[0].Marks[0].Marks))
	fmt.Println("Average mark Alex Smith by Math", averageMark(journal.Subjects[0].Marks[1].Marks))

}

func averageMark(marks []float64) float64 {
	var sum float64
	for _, mark := range marks {
		sum += mark
	}
	return sum / float64(len(marks))
}

func averageSubjectMark(studentMarks []StudentMarks) float64 {
	var sum float64
	for _, mark := range studentMarks {
		sum += averageMark(mark.Marks)
	}
	return sum / float64(len(studentMarks))
}
