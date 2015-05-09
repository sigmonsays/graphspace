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
	"strconv"
	"strings"

	"github.com/sigmonsays/go-apachelog"
	gologging "github.com/sigmonsays/go-logging"
	"github.com/sigmonsays/graphspace/data"
)

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
	Id    int64  `json:"id"`
	Image string `json:"image"`
	Text  string `json:"text"`
}

func (h *GraphvizHandler) Proc(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Warnf("parse: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_id := r.Form.Get("id")
	var id int64
	if len(_id) > 0 {
		id, err = strconv.ParseInt(_id, 10, 64)
		if err != nil {
			log.Warnf("atoi: %s: %s", _id, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	var graph_string string

	if id > 0 {
		graph_string, err = h.backend.Get(id)
		if err != nil {
			log.Warnf("get: %s: %s", id, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	} else {
		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Warnf("read: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		graph_string = string(req)
	}
	cmdline := []string{
		"dot", "-Tpng",
	}
	log.Tracef("request %s", strings.Replace(graph_string, "\n", " ", -1))

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

	if id == 0 {
		id, err = h.backend.Create(graph_string)
		if err != nil {
			log.Warnf("%s [%s] - %s", cmdline, graph_string, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	ret := &Response{
		Id:    id,
		Image: image,
		Text:  graph_string,
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
