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

type MarksIterator struct {
	index int
	marks []StudentMarks
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
						Marks: []float64{4.2, 5, 4.0, 4.1},
					},
					{
						Student: Student{
							Name:    "Alex",
							Surname: "Smith",
						},
						Marks: []float64{4.0, 5, 4.5, 3.8},
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

	math := journal.GetMarksIteratorBySubject("Math")
	fmt.Printf("Average mark of Math  %.2f\n", math.averageMark())
	fmt.Printf("Sum of marks of Math  %.2f\n", math.sumOfMarks())

	eng := journal.GetMarksIteratorBySubject("English")
	fmt.Printf("Average mark of English  %.2f\n", eng.averageMark())
	fmt.Printf("Sum of marks of English  %.2f\n", eng.sumOfMarks())

	biology := journal.GetMarksIteratorBySubject("Biology")
	fmt.Printf("Average mark of Biology  %.2f\n", biology.averageMark())
	fmt.Printf("Sum of marks of Biology  %.2f\n", biology.sumOfMarks())

}

func (j Journal) GetMarksIteratorBySubject(subjectTitle string) *MarksIterator {
	for _, subject := range j.Subjects {
		if subjectTitle == subject.Title {
			return subject.createIterator()
		}
	}

	return nil
}

func (subject *Subject) createIterator() *MarksIterator {
	return &MarksIterator{
		index: -1, // -1 because we need to call Next() before first iteration
		marks: subject.Marks,
	}
}

func (m *MarksIterator) HasNext() bool {
	if m == nil {
		return false
	}

	return m.index < len(m.marks)-1
}

func (m *MarksIterator) Next() *StudentMarks {
	if !m.HasNext() {
		return nil
	}

	m.index++
	return &m.marks[m.index]
}

func (m *MarksIterator) Rewind() *StudentMarks {
	if m == nil {
		return nil
	}
	m.index = -1
	return m.Next()
}

func (m *MarksIterator) averageMark() float64 {
	m.Rewind()
	var sum float64
	var count int
	for m.HasNext() {
		mark := m.Next()
		for _, studentMark := range mark.Marks {
			sum += studentMark
			count++
		}
	}
	return sum / float64(count)
}

func (m *MarksIterator) sumOfMarks() float64 {
	m.Rewind()
	var sum float64
	for m.HasNext() {
		mark := m.Next()
		for _, studentMark := range mark.Marks {
			sum += studentMark
		}
	}
	m.Rewind()
	return sum
}
