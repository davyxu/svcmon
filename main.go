package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// https://stackoverflow.com/questions/8875038/redirect-stdout-pipe-of-child-process-in-go
func monsvc(filename string) error {

	cmd := exec.Command(filename)
	outpipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	errpipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()

	if err != nil {
		return err
	}

	defer cmd.Wait()

	go io.Copy(os.Stdout, outpipe)
	go io.Copy(os.Stderr, errpipe)

	return nil
}

func main() {

	err := monsvc("montest.exe")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
