package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Summary struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

var summaries []Summary

func main() {
	http.HandleFunc("/summaries", summariesHandler)
	http.HandleFunc("/summaries/", summaryHandler)

	log.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func summariesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getSummaries(w)
	case http.MethodPost:
		createSummary(w, r)
	}
}

func getSummaries(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summaries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createSummary(w http.ResponseWriter, r *http.Request) {
	var summary Summary

	if err := json.NewDecoder(r.Body).Decode(&summary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	summary.ID = len(summaries) + 1
	summaries = append(summaries, summary)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(summary)
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) != 2 || parts[0] != "summaries" {
		http.Error(w, "Invalid URL", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid summary ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getSummary(w, id)
	case http.MethodPut:
		updateSummary(w, r, id)
	case http.MethodDelete:
		deleteSummary(w, id)
	}
}

func getSummary(w http.ResponseWriter, id int) {
	for _, s := range summaries {
		if s.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(s)
			return
		}
	}
	http.Error(w, "Summary not found", http.StatusNotFound)
}

func updateSummary(w http.ResponseWriter, r *http.Request, id int) {
	var summary Summary

	if err := json.NewDecoder(r.Body).Decode(&summary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for i, s := range summaries {
		if s.ID == id {
			summaries[i].Description = summary.Description
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(summaries[i])
			return
		}
	}
	http.Error(w, "Summary not found", http.StatusNotFound)
}

func deleteSummary(w http.ResponseWriter, id int) {
	index := -1

	for i, s := range summaries {
		if s.ID == id {
			index = i
		}
	}

	if index == -1 {
		http.Error(w, "Summary not found", http.StatusNotFound)
		return
	}

	summaries = append(summaries[:index], summaries[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}
