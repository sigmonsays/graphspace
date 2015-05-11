package main

import (
	"fmt"
	"net/http"
	"path"
)

// get an image
func (h *GraphvizHandler) Image(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		WriteError(w, r, err)
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		WriteError(w, r, err)
		return
	}

	g, err := h.backend.Get(id)
	if err != nil {
		WriteError(w, r, err)
		return
	}

	// let the extension used in the filename override the output type
	ext := path.Ext(r.URL.Path)
	if len(ext) > 0 {
		ext = ext[1:]
	}
	if _, ok := Outputs[ext]; ok {
		g.Output = ext
	}

	response, err := h.builder.GraphvizImage(g)
	if err != nil {
		WriteError(w, r, err)
		return
	}

	hdr := w.Header()
	hdr.Set("Content-Type", response.ContentType)
	hdr.Set("Content-Length", fmt.Sprintf("%d", len(response.Bytes)))
	w.Write(response.Bytes)
}
