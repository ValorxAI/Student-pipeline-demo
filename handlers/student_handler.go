package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jayashridesh/student-api/models"
	"github.com/jayashridesh/student-api/services"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{
		service: service,
	}
}

func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	students := h.service.GetAll()
	fmt.Println(students)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(students)
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/students/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	student, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if student.Name == "" || student.Email == "" || student.Course == "" {
		http.Error(w, "Name, email and course are required", http.StatusBadRequest)
		return
	}

	createdStudent := h.service.Create(student)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(createdStudent)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/students/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
