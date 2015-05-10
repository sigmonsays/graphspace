package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/sigmonsays/go-apachelog"
	gologging "github.com/sigmonsays/go-logging"
	"github.com/sigmonsays/graphspace/data"
)

type Request struct {
	Format        string
	Text          string
	Width, Height string
	Output        string
}

type Response struct {
	Id          string `json:"id"`
	Format      string `json:"format"`
	Image       string `json:"image"`
	Text        string `json:"text"`
	Output      string `json:"output"`
	ContentType string `json:"content_type"`
}
type GraphvizHandler struct {
	backend *sqlGraphviz
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	log.Warnf("%s", err)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%s", err)
}

func NewGraphvizHandler() (*GraphvizHandler, error) {
	backend, err := NewSqlGraphviz("/tmp/graphspace.db")
	if err != nil {
		return nil, err
	}
	gr := &GraphvizHandler{
		backend: backend,
	}
	return gr, nil
}

func (h *GraphvizHandler) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/static/index.html"
	}
	h.Static(w, r)

}

func (h *GraphvizHandler) Static(w http.ResponseWriter, r *http.Request) {
	asset := ""
	if len(r.URL.Path) > 1 {
		asset = r.URL.Path[1:]
	}
	buf, err := data.Asset(asset)
	if err != nil {
		WriteError(w, r, err)
		return
	}
	w.Write(buf)

}

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
	}

	response, err := GraphvizImage(g)
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

	ret := &Response{
		Id:          id,
		Format:      g.Format,
		Image:       image,
		Text:        g.Text,
		Output:      g.Output,
		ContentType: response.ContentType,
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

	response, err := GraphvizImage(g)
	if err != nil {
		WriteError(w, r, err)
		return
	}

	hdr := w.Header()
	hdr.Set("Content-Type", response.ContentType)
	hdr.Set("Content-Length", fmt.Sprintf("%d", len(response.Bytes)))
	w.Write(response.Bytes)
}

func main() {
	gologging.SetLogLevel("trace")
	addr := ":7001"
	flag.StringVar(&addr, "addr", addr, "http server address")
	flag.Parse()

	svc, err := NewGraphvizHandler()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.Index)
	mux.HandleFunc("/static/", svc.Static)
	mux.HandleFunc("/proc", svc.Proc)
	mux.HandleFunc("/image/", svc.Image)

	handler := apachelog.NewHandler(mux, os.Stderr)
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		panic(err)
	}
}
