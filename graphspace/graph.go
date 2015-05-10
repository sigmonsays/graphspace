package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/awalterschulze/gographviz"
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

func GraphvizImage(g *Graph) ([]byte, error) {
	var err error
	cmd_name, ok := SupportedFormats[g.Format]
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
		return nil, fmt.Errorf("empty graph")
	}

	_, err = gographviz.Read([]byte(g.Text))
	if err != nil {
		return nil, fmt.Errorf("graph: %s", err)
	}

	response := bytes.NewBuffer(nil)
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdin = bytes.NewBuffer([]byte(g.Text))
	cmd.Stdout = response
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	return response.Bytes(), nil
}
