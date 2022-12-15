package buildtask

import (
	"github.com/wenlng/gonacli/clog"
	"github.com/wenlng/gonacli/cmd"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
)

func installDep(cfgs config.Config) bool {
	path := tools.FormatDirPath(cfgs.OutPut)

	clog.Info("Staring install dependencies ...")
	// "bindings" "node-addon-api"
	msg, err := cmd.RunCommand(
		"./",
		"cd "+path+" && npm install bindings node-addon-api -S",
	)
	if err != nil {
		//clog.Warning("Please check whether the \"npm\" command is executed correctly.")
		clog.Error(err)
		return false
	}
	clog.Info(msg)
	clog.Info("Install dependencies done ~")
	return true
}
