package hw9

import (
	"context"
	"course/hw9/class"
	"course/hw9/student"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	studentStorage *student.Storage
	classStorage   *class.Storage
}

func NewHandler() Handler {
	return Handler{
		studentStorage: student.NewStudentStorage(),
		classStorage:   class.NewClassStorage(),
	}
}

func (h *Handler) GetStudentsByClass(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request")
	vars := mux.Vars(r)

	fmt.Println("router params were:", vars)
	classTitle := vars["class"]
	if classTitle == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !h.classStorage.IsTeacherResponsibility(classTitle, r.Context().Value("teacher").(string)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	students := h.studentStorage.GetStudentsByClass(classTitle)

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

	student, err := h.studentStorage.GetStudentById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !h.classStorage.IsTeacherResponsibility(student.Class, r.Context().Value("teacher").(string)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(student)
}

func (h *Handler) GetClassStatistic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request")
	vars := mux.Vars(r)

	fmt.Println("router params were:", vars)
	classTitle := vars["class"]
	if classTitle == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !h.classStorage.IsTeacherResponsibility(classTitle, r.Context().Value("teacher").(string)) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	statistic := h.classStorage.GetClassStatistic(classTitle)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(statistic)
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		t, err := h.classStorage.GetTeacherByUsername(username)

		if err != nil || t.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "teacher", t.Username)
		newReq := r.WithContext(ctx)

		next.ServeHTTP(w, newReq)
	})
}
