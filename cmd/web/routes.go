package main

import (
	"net/http"

	"github.com/justinas/alice"
	"snippetbox.stwn.dev/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// handle static files
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	// Add a new GET /ping route.
	mux.HandleFunc("GET /ping", ping)

	// session middleware for handler
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	// public
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.handlerHome))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.handlerSnippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.handlerUserSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.handlerUserSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.handlerUserLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.handlerUserLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	// protected
	mux.Handle("POST /user/logout", protected.ThenFunc(app.handlerUserLogoutPost))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.handlerSnippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.handlerSnippetCreatePost))

	// middleware chain
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
