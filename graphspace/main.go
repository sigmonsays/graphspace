package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/sigmonsays/go-apachelog"
	gologging "github.com/sigmonsays/go-logging"
	"github.com/sigmonsays/graphspace/data"

)

type GraphvizHandler struct {
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

func (h *GraphvizHandler) Proc(w http.ResponseWriter, r *http.Request) {

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmdline := []string{
		"dot", "-Tpng",
	}
	log.Tracef("request %s", strings.Replace(string(req), "\n", " ", -1))

	response := bytes.NewBuffer(nil)
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdin = bytes.NewBuffer(req)
	cmd.Stdout = response
	err = cmd.Run()
	if err != nil {
		log.Warnf("%s [%s] - %s", cmdline, req, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	str := base64.StdEncoding.EncodeToString(response.Bytes())
	fmt.Fprintf(w, str)

}

func main() {
	gologging.SetLogLevel("trace")
	addr := ":7001"
	flag.StringVar(&addr, "addr", addr, "http server address")
	flag.Parse()

	svc := &GraphvizHandler{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.Index)
	mux.HandleFunc("/static/", svc.Static)
	mux.HandleFunc("/proc", svc.Proc)

	handler := apachelog.NewHandler(mux, os.Stderr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		panic(err)
	}
}
