package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/awalterschulze/gographviz"
	"github.com/sigmonsays/go-apachelog"
	gologging "github.com/sigmonsays/go-logging"
	"github.com/sigmonsays/graphspace/data"
)

var SupportedFormats = map[string]string{
	"dot":       "dot",
	"neato":     "neato",
	"twopi":     "twopi",
	"circo":     "circo",
	"fdp":       "fdp",
	"sfdp":      "sfdp",
	"patchwork": "patchwork",
}

type GraphvizHandler struct {
	backend *sqlGraphviz
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
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Write(buf)

}

type Request struct {
	Format        string
	Text          string
	Width, Height string
}

type Response struct {
	Id     string `json:"id"`
	Format string `json:"format"`
	Image  string `json:"image"`
	Text   string `json:"text"`
}

func (h *GraphvizHandler) Proc(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Warnf("parse: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.Form.Get("id")

	req := &Request{
		Format: "dot",
	}
	g := &Graph{}

	if id != "" {

		if id == "example" {
			g = &Graph{
				Format: "dot",
				Text:   Example,
			}

		} else if id == "random" {
			g.Format = req.Format
			g.Text = RandomGraph()

		} else {

			g, err = h.backend.Get(id)
			if err != nil {
				log.Warnf("get: %s: %s", id, err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

	} else {
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			log.Warnf("decode: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		g.Text = req.Text
		g.Format = req.Format

		if req.Width != "" {
			val, err := strconv.ParseInt(req.Width, 10, 32)
			if err != nil {
				log.Warnf("width: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			g.Width = int(val)
		}
		if req.Height != "" {
			val, err := strconv.ParseInt(req.Height, 10, 32)
			if err != nil {
				log.Warnf("height: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			g.Height = int(val)
		}
	}

	cmd_name, ok := SupportedFormats[req.Format]
	if ok == false {
		cmd_name = "dot"
	}
	cmdline := []string{
		cmd_name, "-Tpng",
	}

	if g.Width > 0 && g.Height > 0 {
		cmdline = append(cmdline, fmt.Sprintf("-Gsize=%d,%d!", g.Width, g.Height))
		cmdline = append(cmdline, "-Gdpi=100")
	}

	if g.Text == "" {
		log.Warnf("empty graph")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = gographviz.Read([]byte(g.Text))
	if err != nil {
		log.Warnf("graph: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := bytes.NewBuffer(nil)
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdin = bytes.NewBuffer([]byte(g.Text))
	cmd.Stdout = response
	err = cmd.Run()
	if err != nil {
		log.Warnf("%s [%s] - %s", cmdline, g.Text, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	image := base64.StdEncoding.EncodeToString(response.Bytes())

	if id == "" {
		id, err = h.backend.Create(g)
		if err != nil {
			log.Warnf("%s - %s", cmdline, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	ret := &Response{
		Id:     id,
		Format: g.Format,
		Image:  image,
		Text:   g.Text,
	}

	json.NewEncoder(w).Encode(ret)

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

	handler := apachelog.NewHandler(mux, os.Stderr)
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		panic(err)
	}
}
