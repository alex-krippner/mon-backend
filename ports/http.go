package ports

import (
	"mon-backend/app"
	"mon-backend/app/handler"
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
	}

	readings, err := h.app.Handlers.ReadingHandler.TranslateReading(r.Context(), newReading)

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.Respond(w, r, readings)
}
