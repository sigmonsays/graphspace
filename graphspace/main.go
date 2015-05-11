package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sigmonsays/go-apachelog"
	gologging "github.com/sigmonsays/go-logging"
	"github.com/sigmonsays/graphspace/data"
)

type Request struct {
	Format        string
	Text          string
	Width, Height string
	Output        string
	Description   string
}

type Response struct {
	Id          string `json:"id"`
	Format      string `json:"format"`
	Image       string `json:"image"`
	Text        string `json:"text"`
	Output      string `json:"output"`
	ContentType string `json:"content_type"`
	Description string `json:"description"`
}
type GraphvizHandler struct {
	backend *sqlGraphviz
	builder *GraphBuilder
}

func WriteResponse(w http.ResponseWriter, r *http.Request, reply interface{}) {
	err := json.NewEncoder(w).Encode(reply)
	if err != nil {
		log.Warnf("%s", err)
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	log.Warnf("%s", err)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "%s", err)
}

func NewGraphvizHandler(dbpath string, builder *GraphBuilder) (*GraphvizHandler, error) {
	backend, err := NewSqlGraphviz(dbpath)
	if err != nil {
		return nil, err
	}
	gr := &GraphvizHandler{
		backend: backend,
		builder: builder,
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

func main() {
	gologging.SetLogLevel("trace")
	addr := ":7001"
	datapath := "/tmp/graphspace"
	flag.StringVar(&addr, "addr", addr, "http server address")
	flag.StringVar(&datapath, "data", datapath, "data path")
	flag.Parse()

	os.MkdirAll(datapath, 0755)
	dbpath := filepath.Join(datapath, "graphspace.db")
	cachepath := filepath.Join(datapath, "cache")
	os.MkdirAll(cachepath, 0755)

	builder := NewGraphBuilder(cachepath)

	svc, err := NewGraphvizHandler(dbpath, builder)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.Index)
	mux.HandleFunc("/static/", svc.Static)
	mux.HandleFunc("/proc", svc.Proc)
	mux.HandleFunc("/image/", svc.Image)
	mux.HandleFunc("/api/list", svc.List)

	handler := apachelog.NewHandler(mux, os.Stderr)
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		panic(err)
	}
}
