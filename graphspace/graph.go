package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/awalterschulze/gographviz"
)

type Graph struct {
	Description   string
	Format        string
	Text          string
	Width, Height int
	Output        string
}

func (g *Graph) GetId() string {
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%d-%d", g.Width, g.Height))
	io.WriteString(h, g.Format)
	text := g.Text
	if graph, err := gographviz.Read([]byte(g.Text)); err == nil {
		text = graph.String()
	}
	io.WriteString(h, text)
	io.WriteString(h, g.Output)
	io.WriteString(h, g.Description)
	b64 := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return strings.Replace(b64, "=", "", -1)

}

var Outputs = map[string]string{
	"png": "image/png",
	"jpg": "image/jpg",
	"svg": "image/svg+xml",
}
var SupportedFormats = map[string]string{
	"dot":       "dot",
	"neato":     "neato",
	"twopi":     "twopi",
	"circo":     "circo",
	"fdp":       "fdp",
	"sfdp":      "sfdp",
	"patchwork": "patchwork",
}

type Image struct {
	ContentType string
	Bytes       []byte
}

func GraphvizImage(g *Graph) (*Image, error) {
	var err error
	cmd_name, ok := SupportedFormats[g.Format]
	if ok == false {
		cmd_name = "dot"
	}
	cmdline := []string{
		cmd_name,
	}

	content_type, ok := Outputs[g.Output]
	if ok == false {
		g.Output = "png"
		content_type = Outputs[g.Output]
	}
	cmdline = append(cmdline, "-T"+g.Output)

	if g.Width > 0 && g.Height > 0 {
		cmdline = append(cmdline, fmt.Sprintf("-Gsize=%d,%d!", g.Width, g.Height))
		cmdline = append(cmdline, "-Gdpi=100")
	}

	if g.Text == "" {
		return nil, fmt.Errorf("empty graph")
	}

	_, err = gographviz.Read([]byte(g.Text))
	if err != nil {
		return nil, fmt.Errorf("graph: %s", err)
	}

	log.Tracef("%s", cmdline)

	response := bytes.NewBuffer(nil)
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdin = bytes.NewBuffer([]byte(g.Text))
	cmd.Stdout = response
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	ret := &Image{
		Bytes:       response.Bytes(),
		ContentType: content_type,
	}
	return ret, nil
}
