package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

type Terraform struct {
	path, workdir string
	silent        bool
}

type TerraformOption func(*Terraform)

func (t *Terraform) Exec(args []string) error {
	var out bytes.Buffer

	cmd := exec.Command(t.path, args...)
	cmd.Dir = t.workdir
	cmd.Env = os.Environ()
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return err
	}

	if !t.silent {
		fmt.Print(out.String())
	}

	return nil
}

func Silent() TerraformOption {
	return func(t *Terraform) {
		t.silent = true
	}
}

func WorkingDir(workdir string) TerraformOption {
	return func(t *Terraform) {
		t.workdir = workdir
	}
}

func NewTerraform(options ...TerraformOption) (*Terraform, error) {
	path, err := exec.LookPath("terraform")
	if err != nil {
		return nil, errors.Wrap(err, "could not find terraform")
	}

	t := &Terraform{
		path:    path,
		workdir: ".",
		silent:  false,
	}
	for _, opt := range options {
		opt(t)
	}
	return t, nil
}
