package services

import (
	"errors"
	"sync"

	"github.com/jayashridesh/student-api/models"
)

type StudentService struct {
	students []models.Student
	mu       sync.Mutex
}

func NewStudentService() *StudentService {
	return &StudentService{
		students: []models.Student{
			{
				ID:     1,
				Name:   "Rahul",
				Age:    20,
				Course: "B.Com",
				Email:  "rahul@example.com",
			},
			{
				ID:     2,
				Name:   "Priya",
				Age:    21,
				Course: "BBA",
				Email:  "priya@example.com",
			},
		},
	}
}

func (s *StudentService) GetAll() []models.Student {
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]models.Student(nil), s.students...)
}

func (s *StudentService) GetByID(id int) (*models.Student, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, student := range s.students {
		if student.ID == id {
			result := student
			return &result, nil
		}
	}

	return nil, errors.New("student not found")
}

func (s *StudentService) Create(student models.Student) models.Student {
	s.mu.Lock()
	defer s.mu.Unlock()

	student.ID = len(s.students) + 1
	s.students = append(s.students, student)

	return student
}

func (s *StudentService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, student := range s.students {
		if student.ID == id {
			s.students = append(s.students[:i], s.students[i+1:]...)
			return nil
		}
	}

	return errors.New("student not found")
}
