package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

func (h *GraphvizHandler) Proc(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		WriteError(w, r, err)
		return
	}

	id := r.Form.Get("id")

	req := &Request{
		Format: "dot",
		Output: "png",
	}
	g := &Graph{}

	if id != "" {

		if id == "example" {
			g = &Graph{
				Format: "dot",
				Text:   Example,
			}
			id = g.GetId()

		} else if id == "random" {
			g.Format = req.Format
			g.Text = RandomGraph()
			id = g.GetId()

		} else {

			g, err = h.backend.Get(id)
			if err != nil {
				WriteError(w, r, err)
				return
			}
		}

	} else {
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			WriteError(w, r, err)
			return
		}
		g.Text = req.Text
		g.Format = req.Format

		if req.Width != "" {
			val, err := strconv.ParseInt(req.Width, 10, 32)
			if err != nil {
				WriteError(w, r, err)
				return
			}
			g.Width = int(val)
		}
		if req.Height != "" {
			val, err := strconv.ParseInt(req.Height, 10, 32)
			if err != nil {
				WriteError(w, r, err)
				return
			}
			g.Height = int(val)
		}

		g.Output = req.Output
		g.Description = req.Description
	}

	response, err := h.builder.GraphvizImage(g)
	if err != nil {
		WriteError(w, r, err)
		return
	}

	image := base64.StdEncoding.EncodeToString(response.Bytes)

	id, err = h.backend.Create(g)
	if err != nil {
		log.Infof("create %s: %s", id, err)
	}
	if id == "" {
		id = g.GetId()
	}

	if g.Description == "" {
		g.Description = "graph"
	}

	ret := &Response{
		Id:          id,
		Format:      g.Format,
		Image:       image,
		Text:        g.Text,
		Output:      g.Output,
		ContentType: response.ContentType,
		Description: g.Description,
	}

	json.NewEncoder(w).Encode(ret)

}

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
