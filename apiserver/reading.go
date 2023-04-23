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
}

func (s *APIServer) createReading(w http.ResponseWriter, req *http.Request) error {
	var r storage.Reading

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &r)
	reading, err := s.storage.CreateReading(req.Context(), storage.CreateReadingRequest{
		EnglishTranslation: r.EnglishTranslation,
		Japanese:           r.Japanese,
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
