package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

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
	r.URL.Path = "/static/index.html"
	h.Static(w, r)
}

func (h *GraphvizHandler) Static(w http.ResponseWriter, r *http.Request) {
	asset := ""
	if len(r.URL.Path) > 1 {
		asset = r.URL.Path[1:]
	}
	buf, err := data.Asset(asset)
	if err == nil {
		w.Write(buf)
	}

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

	format := r.Form.Get("format")
	id := r.Form.Get("id")

	var graph_string string
	var g *Graph

	if id != "" {

		if id == "example" {
			g = &Graph{
				Id:   "example",
				Text: Example,
			}
			graph_string = g.Text

		} else if id == "random" {
			graph_string = RandomGraph()
			g = &Graph{
				Text: graph_string,
			}

		} else {

			g, err = h.backend.Get(id)
			if err != nil {
				log.Warnf("get: %s: %s", id, err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			graph_string = g.Text
			format = g.Format
		}

	} else {
		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Warnf("read: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		graph_string = string(req)
		g = &Graph{
			Text: graph_string,
		}
	}

	cmd_name, ok := SupportedFormats[format]
	if ok == false {
		cmd_name = "dot"
	}
	cmdline := []string{
		cmd_name, "-Tpng",
	}

	response := bytes.NewBuffer(nil)
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdin = bytes.NewBuffer([]byte(graph_string))
	cmd.Stdout = response
	err = cmd.Run()
	if err != nil {
		log.Warnf("%s [%s] - %s", cmdline, graph_string, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	image := base64.StdEncoding.EncodeToString(response.Bytes())

	g = &Graph{
		Format: format,
		Text:   graph_string,
	}

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
		Format: format,
		Image:  image,
		Text:   graph_string,
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
