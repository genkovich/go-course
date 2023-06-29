package student

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	storage *Storage
}

func NewHandler() Handler {
	return Handler{storage: NewStudentStorage()}
}

func (h *Handler) GetStudentsByClass(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request")
	vars := mux.Vars(r)

	fmt.Println("router params were:", vars)
	class := vars["class"]
	if class == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	students := h.storage.GetStudentsByClass(class)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(students)
}

func (h *Handler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request")
	vars := mux.Vars(r)

	fmt.Println("router params were:", vars)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID format")
		return
	}

	student, err := h.storage.GetStudentById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(student)
}

func (h *Handler) GetClassStatistic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request")
	vars := mux.Vars(r)

	fmt.Println("router params were:", vars)
	class := vars["class"]
	if class == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	statistic := h.storage.GetClassStatistic(class)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(statistic)
}
