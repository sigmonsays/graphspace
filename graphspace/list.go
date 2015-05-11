package main

import (
	"net/http"
)

type listEntry struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

func (h *GraphvizHandler) List(w http.ResponseWriter, r *http.Request) {

	ls, err := h.backend.ListRecent(10)
	if err != nil {
		WriteError(w, r, err)
	}

	WriteResponse(w, r, ls)
}
