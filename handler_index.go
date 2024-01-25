package main

import (
	"net/http"
	"text/template"
)

func (apiCfg apiConfig) handlerRenderindex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error rendering index")
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error rendering index")
		return
	}
}
