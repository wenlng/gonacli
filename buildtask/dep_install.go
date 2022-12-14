package buildtask

import (
	"github.com/wenlng/gonacli/clog"
	"github.com/wenlng/gonacli/cmd"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
)

func installDep(cfgs config.Config) bool {
	path := tools.FormatDirPath(cfgs.OutPut)

	clog.Info("Staring install npm dependencies ...")
	// "bindings" "node-addon-api"
	msg, err := cmd.RunCommand(
		"./",
		"cd "+path+" && npm install bindings node-addon-api -S",
	)
	if err != nil {
		clog.Error(err)
		return false
	}
	clog.Info(msg)
	clog.Info("Install npm dependencies done ~")
	return true
}
