package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func selectWithPeco(config *Config, flags *Flags) []string {
	candidates := config.buildCandidates(flags)
	pecoResult := runPeco(candidates)
	return strings.Split(pecoResult, "\n")
}

func runPeco(candidates string) string {
	pecoBin := "peco"
	_, err := exec.LookPath(pecoBin)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(pecoBin)
	cmd.Stdin = strings.NewReader(candidates)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}
