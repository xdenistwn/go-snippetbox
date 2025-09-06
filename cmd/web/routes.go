package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// handle static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.handlerHome)
	mux.HandleFunc("GET /snippet/view/{id}", app.handlerSnippetView)
	mux.HandleFunc("GET /snippet/create", app.handlerSnippetCreate)
	mux.HandleFunc("POST /snippet/create", app.handlerSnippetCreatePost)

	// middleware chain
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
