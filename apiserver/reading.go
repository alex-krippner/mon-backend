package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mon-backend/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) addReadingHandlers(r *mux.Router) {

	r.Methods("POST").Path("/reading").Handler(Endpoint{s.createReading})
	r.Methods("GET").Path("/reading").Handler(Endpoint{s.getAllReading})
	r.Methods("PATCH").Path("/reading").Handler(Endpoint{s.updateReading})
}

func (s *APIServer) createReading(w http.ResponseWriter, req *http.Request) error {
	var r storage.Reading

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &r)
	reading, err := s.storage.CreateReading(req.Context(), storage.CreateReadingRequest{
		Translation: r.Translation,
		Japanese:    r.Japanese,
		Title:       r.Title,
	})

	if err != nil {
		log.Print(err.Error())
		return err
	}

	jsonResponse, err := json.Marshal(reading)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	return nil
}

func (s *APIServer) getAllReading(w http.ResponseWriter, req *http.Request) error {
	readingSlice, err := s.storage.GetAllReading(req.Context())
	if err != nil {
		return err
	}
	jsonResponse, err := json.Marshal(readingSlice)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}

func (s *APIServer) updateReading(w http.ResponseWriter, req *http.Request) error {
	var r storage.Reading

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &r)

	reading, err := s.storage.UpdateReading(req.Context(), r)

	if err != nil {
		log.Print(err)
		return err
	}

	jsonResponse, err := json.Marshal(reading)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	return nil
}
