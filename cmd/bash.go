//go:build !windows
// +build !windows

package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func runCommand(path, name string, arg string, execStr string) (msg string, err error) {
	cmd := exec.Command(name, arg, execStr)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = path
	err = cmd.Run()
	//clog.Info(cmd.Args)
	msg = fmt.Sprintf("%s", err) + ": " + fmt.Sprintf("%s", stderr.String())
	if err != nil {
		err = errors.New(msg)
	}
	//log.Println(out.String())
	return
}
