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

	// session middleware for handler
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.handlerHome))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.handlerSnippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.handlerSnippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.handlerSnippetCreatePost))

	// middleware chain
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
