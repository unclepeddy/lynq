package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func apply(t *Terraform) {
	log.Print("running terraform init:")
	initArgs := []string{
		"init",
		"-no-color",
	}
	if err := t.Exec(initArgs); err != nil {
		log.Printf("failed to run init: %v", err)
		return
	}

	log.Print("running terraform apply:")
	applyArgs := []string{
		"apply",
		"-no-color",
		"-auto-approve",
	}
	if err := t.Exec(applyArgs); err != nil {
		log.Printf("failed to run apply: %v", err)
		return
	}

	log.Print("successfully applied configuration")
}

func main() {
	var workingdir string
	var period int
	flag.StringVar(&workingdir, "dir", "", "working directory for terraform")
	flag.IntVar(&period, "t", 120, "loop period in seconds")
	flag.Parse()

	var opts []TerraformOption
	if workingdir != "" {
		opts = append(opts, WorkingDir(workingdir))
	}
	t, err := NewTerraform(opts...)
	if err != nil {
		log.Fatalf("unable to contruct new terraform client: %v", err)
	}

	log.Print("running terraform version:")
	versionArgs := []string{
		"version",
	}
	if err := t.Exec(versionArgs); err != nil {
		log.Fatalf("failed to run version: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

loop:
	for {
		select {
		case <-time.After(time.Duration(period) * time.Second):
			apply(t)
		case <-ch:
			break loop
		}
	}

	log.Print("gracefully terminating")
}
