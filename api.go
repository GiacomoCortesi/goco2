package goco2

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// APIServer is the server exposing the microservice REST API
type APIServer struct {
	svc Service
}

func NewAPIServer(svc Service) *APIServer {
	return &APIServer{svc: svc}
}

// Start runs the microservice on the specified listenAddr
// Start is blocking
func (s *APIServer) Start(listenAddr string) error {
	http.HandleFunc("/saving", s.GetCO2SavingHandler())
	http.HandleFunc("/intervention", s.AddInterventionHandler())

	// serve static files for the basic open AI UI
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", HomePage)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *APIServer) GetCO2SavingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		saving, err := s.svc.GetCO2Saving(context.TODO())
		if err != nil {
			_ = writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		}
		_ = writeJSON(w, http.StatusOK, saving)
	}
}

func (s *APIServer) AddInterventionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var intervention Intervention
		err := readJSON(r, &intervention)
		if err != nil {
			_ = writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		err = s.svc.AddIntervention(context.TODO(), intervention)
		if err != nil {
			_ = writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
			return
		}
		_ = writeJSON(w, http.StatusNoContent, nil)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	type PageVariables struct {
		Title string
	}
	pageVariables := PageVariables{
		Title: "Microservice GUI",
	}
	tmplFile := "static/index.html"
	tmpl, err := template.New("index.html").ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func readJSON(r *http.Request, v any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("invalid content type, must be: application/json")
	}
	return json.NewDecoder(r.Body).Decode(v)
}
