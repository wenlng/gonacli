package cmd

func RunCommand(path string, exec string) (msg string, err error) {
	cmdName := "/bin/bash"
	cmdArg := "-c"
	msg, err = runCommand(path, cmdName, cmdArg, exec)
	return
}
