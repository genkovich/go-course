package task

import "time"

type Task struct {
	Id    int
	Title string
	Date  time.Time
}

type Provider interface {
	GetTasksByDate(date time.Time) []Task
}

type Storage struct {
	tasks []Task
}

func NewTaskStorage() *Storage {
	return &Storage{
		tasks: []Task{
			Task{Id: 1, Title: "Task 1", Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			Task{Id: 2, Title: "Task 2", Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)},
			Task{Id: 3, Title: "Task 3", Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)},
			Task{Id: 4, Title: "Task 4", Date: time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC)},
			Task{Id: 5, Title: "Task 5", Date: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC)},
			Task{Id: 6, Title: "Task 6", Date: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC)},
			Task{Id: 7, Title: "Task 7", Date: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC)},
		},
	}
}

func (s *Storage) GetTasksByDate(date time.Time) []Task {
	tasks := make([]Task, 0)
	for _, task := range s.tasks {
		if task.Date == date {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
