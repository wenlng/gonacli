package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/wenlng/gonacli/clog"
	"github.com/wenlng/gonacli/tools"
	"os/exec"
)

func runCommand(path, name string, arg ...string) (msg string, err error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = path
	err = cmd.Run()
	clog.Info(cmd.Args)
	if err != nil {
		msg = fmt.Sprint(err) + ": " + stderr.String()
		err = errors.New(msg)
	}
	//log.Println(out.String())
	return
}

func RunCommand(path string, exec string) (msg string, err error) {
	cmdName := "/bin/bash"
	cmdArg := "-c"
	if tools.IsWindowsOs() {
		cmdName = "cmd"
		cmdArg = "/c"
	}

	msg, err = runCommand(path, cmdName, cmdArg, exec)
	return
}
