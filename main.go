package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ValorxAI/Student-pipeline-demo/handlers"
	"github.com/ValorxAI/Student-pipeline-demo/services"
)

func main() {
	service := services.NewStudentService()
	handler := handlers.NewStudentHandler(service)

	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetStudents(w, r)

		case http.MethodPost:
			handler.CreateStudent(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetStudent(w, r)

		case http.MethodDelete:
			handler.DeleteStudent(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Student API running on port %s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"status":"healthy"}`))
}
