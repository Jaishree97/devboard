package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(HealthResponse{
		Status: "ok",
	})
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	databaseURL := os.Getenv("POSTGRES_URL")

	if databaseURL == "" {
		http.Error(w, "POSTGRES_URL is not configured", http.StatusInternalServerError)
		return
	}

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Printf("database connection failed: %v", err)
		http.Error(w, "database unavailable", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(
		context.Background(),
		"SELECT id, name FROM projects ORDER BY id",
	)
	if err != nil {
		log.Printf("query failed: %v", err)
		http.Error(w, "failed to query projects", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	projects := make([]Project, 0)

	for rows.Next() {
		var project Project

		if err := rows.Scan(&project.ID, &project.Name); err != nil {
			log.Printf("row scan failed: %v", err)
			http.Error(w, "failed to read projects", http.StatusInternalServerError)
			return
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		log.Printf("row iteration failed: %v", err)
		http.Error(w, "failed to read projects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/projects", projectsHandler)

	log.Println("DevBoard backend starting on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
