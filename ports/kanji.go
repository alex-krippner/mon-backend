package ports

import (
	"mon-backend/app/handler"
	"mon-backend/domain/kanji"
	"mon-backend/server/httperr"
	"net/http"

	"github.com/go-chi/render"
)

func (h HttpServer) AddKanji(w http.ResponseWriter, r *http.Request) {
	postKanji := PostKanji{}
	if err := render.Decode(r, &postKanji); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	newKanji := handler.NewKanji{
		ExampleSentences: postKanji.ExampleSentences,
		ExampleWords:     postKanji.ExampleWords,
		Kanji:            postKanji.Kanji,
		KanjiRating:      postKanji.KanjiRating,
		KunReading:       postKanji.KunReading,
		Meanings:         postKanji.Meanings,
		OnReading:        postKanji.OnReading,
		Username:         postKanji.Username,
	}

	kanji, err := h.app.Handlers.KanjiHandler.AddKanji(r.Context(), newKanji)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.Respond(w, r, kanji)
}

func (h HttpServer) GetKanjis(w http.ResponseWriter, r *http.Request, username string) {
	kanjis, err := h.app.Handlers.KanjiHandler.GetAllKanji(r.Context(), username)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, kanjis)
}

func (h HttpServer) GetKanji(w http.ResponseWriter, r *http.Request, kanjiId string) {
	kanji, err := h.app.Handlers.KanjiHandler.GetKanji(r.Context(), kanjiId)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, kanji)
}

func (h HttpServer) UpdateKanji(w http.ResponseWriter, r *http.Request) {
	patchKanji := kanji.Kanji{}
	if err := render.Decode(r, &patchKanji); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	updatedKanji, err := h.app.Handlers.KanjiHandler.UpdateKanji(r.Context(), patchKanji)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, updatedKanji)
}

func (h HttpServer) DeleteKanji(w http.ResponseWriter, r *http.Request, kanjiId string) {
	deletedKanji := DeletedKanji{
		Id: kanjiId,
	}

	err := h.app.Handlers.KanjiHandler.DeleteKanji(kanjiId)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, deletedKanji)
}
