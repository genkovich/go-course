package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Handler struct {
	storage *Storage
}

func (h *Handler) GetTasksByDate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request")

	fmt.Println("GET params were:", r.URL.Query())

	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks := h.storage.GetTasksByDate(dateTime)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(tasks)
}

func NewHandler() Handler {
	taskStorage := NewTaskStorage()
	return Handler{storage: taskStorage}
}
