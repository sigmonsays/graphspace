package main

import (
	"net/http"
)

type DeleteResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *GraphvizHandler) Delete(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		WriteError(w, r, err)
	}

	id := r.Form.Get("id")

	err = h.backend.Delete(id)
	if err != nil {
		WriteError(w, r, err)
	}

	res := &DeleteResponse{
		Message: "deleted",
	}

	WriteResponse(w, r, res)
}
