package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.stwn.dev/internal/models"
)

func (app *application) handlerHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, "home.html", templateData{
		Snippets: snippets,
	})
}
func (app *application) handlerSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	app.render(w, r, http.StatusOK, "view.html", templateData{
		Snippet: snippet,
	})
}
func (app *application) handlerSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}
func (app *application) handlerSnippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "0 snail"
	content := "lorem ipsum try one try two"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
