package ports

import (
	"mon-backend/app"
	"mon-backend/app/handler"
	"mon-backend/domain/reading"
	"mon-backend/server/httperr"
	"net/http"

	"github.com/go-chi/render"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) CreateReading(w http.ResponseWriter, r *http.Request) {
	postReading := PostReading{}
	if err := render.Decode(r, &postReading); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	newReading := handler.NewReading{
		Japanese:    postReading.Japanese,
		Translation: postReading.Translation,
		Title:       postReading.Title,
		Username:    postReading.Username,
	}

	readings, err := h.app.Handlers.ReadingHandler.TranslateReading(r.Context(), newReading)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.Respond(w, r, readings)
}

func (h HttpServer) GetReadings(w http.ResponseWriter, r *http.Request, username string) {
	readings, err := h.app.Handlers.ReadingHandler.GetAllReading(r.Context(), username)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, readings)

}

func (h HttpServer) DeleteReading(w http.ResponseWriter, r *http.Request, readingId string) {
	deletedReading := DeletedReading{
		Id: readingId,
	}

	err := h.app.Handlers.ReadingHandler.DeleteReading(r.Context(), readingId)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, deletedReading)
}

func (h HttpServer) UpdateReading(w http.ResponseWriter, r *http.Request) {
	patchReading := reading.Reading{}
	if err := render.Decode(r, &patchReading); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	updatedReading, err := h.app.Handlers.ReadingHandler.UpdateReading(r.Context(), patchReading)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, updatedReading)
}
