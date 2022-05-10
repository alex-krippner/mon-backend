package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	str "strings"

	"mon-backend/storage"
)

func (s *APIServer) createKanji(w http.ResponseWriter, req *http.Request) error {
	var k storage.Kanji

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), &k)

	kanji, err := s.storage.CreateKanji(req.Context(), storage.CreateKanjiRequest{
		Kanji:       k.Kanji,
		OnReading:   k.OnReading,
		KunReading:  k.KunReading,
		KanjiRating: k.KanjiRating,
		Username:    k.Username,
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

func getIdFromURL(req *http.Request) string {

	splitURL := str.Split(req.URL.String(), "/")
	return splitURL[len(splitURL)-1]
}

func (s *APIServer) deleteKanji(w http.ResponseWriter, req *http.Request) error {

	id := getIdFromURL(req)
	err := s.storage.DeleteKanji(id)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

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

func (s *APIServer) getKanji(w http.ResponseWriter, req *http.Request) error {
	id := getIdFromURL(req)
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

func (s *APIServer) listKanji(w http.ResponseWriter, req *http.Request) error {
	kanjiSlice, err := s.storage.ListKanji(req.Context())
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
