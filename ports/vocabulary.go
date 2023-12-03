package ports

import (
	"errors"
	"mon-backend/app/handler"
	"mon-backend/domain/vocabulary"
	"mon-backend/server/httperr"
	"net/http"

	"github.com/go-chi/render"
)

func (h HttpServer) AddVocab(w http.ResponseWriter, r *http.Request, params AddVocabParams) {
	postVocab := PostVocab{}

	if err := render.Decode(r, &postVocab); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	newVocab := handler.NewVocab{
		Vocab:            postVocab.Vocab,
		Definitions:      postVocab.Definitions,
		ExampleSentences: postVocab.ExampleSentences,
		PartsOfSpeech:    postVocab.PartsOfSpeech,
		Kanji:            postVocab.Kanji,
		VocabRating:      postVocab.VocabRating,
		Username:         postVocab.Username,
	}

	vocab, err := h.app.Handlers.VocabHandler.AddVocab(r.Context(), newVocab)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.Respond(w, r, vocab)
}

func (h HttpServer) GetVocab(w http.ResponseWriter, r *http.Request, params GetVocabParams) {

	if params.Limit == nil {
		httperr.RespondWithSlugError(errors.New("limit query parameter not set"), w, r)
		return
	}

	vocab, err := h.app.Handlers.VocabHandler.GetVocab(r.Context(), *params.Username, *params.Limit)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, vocab)

}

func (h HttpServer) UpdateVocab(w http.ResponseWriter, r *http.Request, params UpdateVocabParams) {
	patchVocab := vocabulary.Vocab{}
	if err := render.Decode(r, &patchVocab); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	updatedVocab, err := h.app.Handlers.VocabHandler.UpdateVocab(r.Context(), patchVocab)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, updatedVocab)
}

func (h HttpServer) DeleteVocab(w http.ResponseWriter, r *http.Request, params DeleteVocabParams) {

	err := h.app.Handlers.VocabHandler.DeleteVocab(*params.Id)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, params.Id)
}
