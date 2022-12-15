//go:build windows
// +build windows

package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"syscall"
)

func runCommand(path, name string, arg string, execStr string) (msg string, err error) {
	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c cd `+path+` && %s`, execStr), HideWindow: true}

	output, err2 := cmd.CombinedOutput()

	msg = string(output)
	if err2 != nil {
		err = errors.New(string(output))
	}

	return
}
