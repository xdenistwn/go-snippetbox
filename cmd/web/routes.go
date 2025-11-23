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
	protected := dynamic.Append(app.requireAuthentication)

	// public
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.handlerHome))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.handlerSnippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.handlerUserSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.handlerUserSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.handlerUserLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.handlerUserLoginPost))

	// protected
	mux.Handle("POST /user/logout", protected.ThenFunc(app.handlerUserLogoutPost))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.handlerSnippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.handlerSnippetCreatePost))

	// middleware chain
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
