//go:build windows
// +build windows

package cmd

import (
	"fmt"
	"os/exec"
	"syscall"
)

func runCommand(path, name string, arg string, execStr string) (msg string, err error) {
	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c cd `+path+` && %s`, execStr), HideWindow: true}
	out, e := cmd.Output()
	msg = fmt.Sprintf("%s", string(out))
	err = e
	return
}
