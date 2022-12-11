package cli

import (
	"flag"
	"fmt"
	"log"
	"node_addon_go/task"
	"os"
)

type CLI struct{}

func printUsage() {
	fmt.Println("Usage:")

	fmt.Println("\tbuild -- Compile the golang source file of the export api")
	fmt.Println("\tgenerate -- Generate napi c/c++ code of golang and addon bridge")
	fmt.Println("\tnodegyp -- Compile addon bindings of nodejs")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	isValidArgs()

	// gonacli build
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	// gonacli generate
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	// gonacli make
	makeCmd := flag.NewFlagSet("make", flag.ExitOnError)

	// gonacli build --config xxx.json
	buildCofig := buildCmd.String("config", "goaddon.json", "Addon api export configuration file")
	// gonacli build --args '-ldflags "-s -w"'
	buildArg := buildCmd.String("args", "'-ldflags \"-s -w\"'", "Golang compilation arguments")
	// gonacli build --upx
	//buildUpx := buildCmd.Bool("upx", false, "Call the upx compression command")
	// gonacli build --xgo
	//buildXgo := buildCmd.Bool("xgo", false, "Call the xgo compression command")
	// gonacli generate --config xxx.json
	generateConfig := generateCmd.String("config", "goaddon.json", "Addon api export configuration file")
	// gonacli make --args "xxx"
	makeArg := makeCmd.String("args", "", "Nodegyp compilation arguments")
	// gonacli make --config xxx.json
	makeConfig := makeCmd.String("config", "goaddon.json", "Addon api export configuration file")
	makeMpn := makeCmd.Bool("npm-i", false, "Install npm dependencies")

	switch os.Args[1] {
	case "build":
		err := buildCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "generate":
		err := generateCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "make":
		err := makeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if buildCmd.Parsed() {
		task.RunBuildTask(*buildCofig, *buildArg)
	}

	if generateCmd.Parsed() {
		task.RunGenerateTask(*generateConfig)
	}

	if makeCmd.Parsed() {
		task.RunMakeTask(*makeConfig, *makeArg, *makeMpn)
	}
}
