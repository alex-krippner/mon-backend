package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"mon-backend/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) addVocabHandlers(r *mux.Router) {

	r.Methods("POST").Path("/vocab").Handler(Endpoint{s.createVocab})
	r.Methods("GET").Path("/vocab/{id}").Handler(Endpoint{s.getVocab})
	r.Methods("GET").Path("/vocab").Handler(Endpoint{s.getAllVocab})
	r.Methods("PATCH").Path("/vocab").Handler(Endpoint{s.updateVocab})
	r.Methods("DELETE").Path("/vocab/{id}").Handler(Endpoint{s.deleteVocab})
}

func (s *APIServer) createVocab(w http.ResponseWriter, req *http.Request) error {
	var v storage.Vocab

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(body), &v)

	if err != nil {
		return err
	}

	vocab, err := s.storage.CreateVocab(req.Context(), storage.CreateVocabRequest{
		Vocab:            v.Vocab,
		Definitions:      v.Definitions,
		ExampleSentences: v.ExampleSentences,
		PartsOfSpeech:    v.PartsOfSpeech,
		Kanji:            v.Kanji,
		VocabRating:      v.VocabRating,
		Username:         v.Username,
	})

	if err != nil {
		return err
	}

	jsonResponse, err := json.Marshal(vocab)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	return nil

}

func (s *APIServer) getVocab(w http.ResponseWriter, req *http.Request) error {
	id := mux.Vars(req)["id"]
	vocab, err := s.storage.GetVocab(req.Context(), id)

	if err != nil {
		return err
	}

	jsonResponse, err := json.Marshal(vocab)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}

func (s *APIServer) getAllVocab(w http.ResponseWriter, req *http.Request) error {
	vocabslice, err := s.storage.GetAllVocab(req.Context())

	if err != nil {
		return err
	}
	jsonResponse, err := json.Marshal(vocabslice)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil

}

func (s *APIServer) updateVocab(w http.ResponseWriter, req *http.Request) error {
	var v storage.UpdateVocabRequest

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &v)

	vocab, err := s.storage.UpdateVocab(req.Context(), v)

	if err != nil {
		return err
	}

	jsonResponse, err := json.Marshal(vocab)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}

func (s *APIServer) deleteVocab(w http.ResponseWriter, req *http.Request) error {
	id := mux.Vars(req)["id"]
	err := s.storage.DeleteVocab(id)

	if err != nil {
		return err
	}

	deleteResponse := DeleteResponse{id}
	response, err := json.Marshal(deleteResponse)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write((response))

	return nil
}
