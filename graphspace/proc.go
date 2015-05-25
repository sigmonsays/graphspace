package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *GraphvizHandler) initCors(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	hdr := w.Header()
	for _, server := range h.Cors {
		if origin == server {
			hdr.Set("Access-Control-Allow-Origin", server)
			break
		}
	}
	hdr.Set("Access-Control-Allow-Methods", "POST")
	hdr.Set("Access-Control-Allow-Headers", "Content-Type")
}

func (h *GraphvizHandler) Proc(w http.ResponseWriter, r *http.Request) {

	if h.Cors != nil {
		h.initCors(w, r)
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

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
				Format:      "dot",
				Description: "OS process state",
				Text:        Example,
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
		log.Tracef("request %#v", req)
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

	if req.Button == "save" {
		id, err = h.backend.Create(g)
		if err != nil {
			log.Infof("create %s: %s", id, err)
		}
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
