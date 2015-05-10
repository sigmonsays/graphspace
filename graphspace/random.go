package main

// make a random graph using gvgen
// http://www.graphviz.org/pdf/gvgen.1.pdf

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
)

func RandomGraph() string {
	response := bytes.NewBuffer(nil)
	cmdline := []string{"gvgen", "-ng"}

	// takes a single integer
	single_flags := []string{
		"-c",
		"-h",
		"-k",
		"-m",
		"-p",
		"-R",
		"-s",
		"-S",
		"-t",
		"-w",
	}
	// takes 2 integers
	/*
		double_flags := []string{
			"-C",
			"-g",
			"-G",
			"-b",
			"-B",
			"-t",
			"-T",
		}
	*/

	intval1 := rand.Intn(8) + 2
	idx := rand.Intn(len(single_flags))
	flag := single_flags[idx]
	cmdline = append(cmdline, flag)
	cmdline = append(cmdline, fmt.Sprintf("%d", intval1))

	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdout = response
	err := cmd.Run()
	if err != nil {
		log.Warnf("%s: %s", cmdline, err)
	}
	log.Infof("command %s", cmdline)

	return response.String()
}
