package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mon-backend/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) addKanjiHandlers(r *mux.Router) {

	r.Methods("POST").Path("/kanji").Handler(Endpoint{s.createKanji})
	r.Methods("GET").Path("/kanji/{id}").Handler(Endpoint{s.getKanji})
	r.Methods("GET").Path("/kanji").Handler(Endpoint{s.getAllKanji})
	r.Methods("PATCH").Path("/kanji").Handler(Endpoint{s.updateKanji})
	r.Methods("DELETE").Path("/kanji/{id}").Handler(Endpoint{s.deleteKanji})
}

func (s *APIServer) createKanji(w http.ResponseWriter, req *http.Request) error {
	var k storage.Kanji

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &k)
	kanji, err := s.storage.CreateKanji(req.Context(), storage.CreateKanjiRequest{
		Kanji:            k.Kanji,
		ExampleSentences: k.ExampleSentences,
		ExampleWords:     k.ExampleWords,
		OnReading:        k.OnReading,
		KunReading:       k.KunReading,
		KanjiRating:      k.KanjiRating,
		Username:         k.Username,
	})

	if err != nil {
		log.Print(err)
		return err
	}

	jsonResponse, err := json.Marshal(kanji)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	return nil
}

func (s *APIServer) getKanji(w http.ResponseWriter, req *http.Request) error {
	id := mux.Vars(req)["id"]
	kanji, err := s.storage.GetKanji(req.Context(), id)

	if err != nil {
		return err
	}

	jsonResponse, err := json.Marshal(kanji)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}

func (s *APIServer) getAllKanji(w http.ResponseWriter, req *http.Request) error {
	kanjiSlice, err := s.storage.GetAllKanji(req.Context())
	if err != nil {
		return err
	}
	jsonResponse, err := json.Marshal(kanjiSlice)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}

func (s *APIServer) updateKanji(w http.ResponseWriter, req *http.Request) error {

	var k storage.UpdateKanjiRequest

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &k)

	kanji, err := s.storage.UpdateKanji(req.Context(), k)

	if err != nil {
		log.Print(err)
		return err
	}

	jsonResponse, err := json.Marshal(kanji)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

	return nil
}

func (s *APIServer) deleteKanji(w http.ResponseWriter, req *http.Request) error {

	id := mux.Vars(req)["id"]
	err := s.storage.DeleteKanji(id)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
