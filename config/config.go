package config

type Arg struct {
	Name      string `env:"Name" default:"arg"`
	Type      string `env:"Type" default:"string"`
	IsRequire bool   `env:"IsRequire" default:"false"`
}

type Export struct {
	Name       string `env:"Name" default:""`
	Args       []Arg  `env:"Args" default:""`
	ReturnType string `env:"ReturnType" default:""`
	JsCallName string `env:"JsCallName" default:"func"`
	JsCallMode string `env:"JsCallMode" default:"sync"`
}

type Config struct {
	Name    string   `env:"Name" required:"true" default:"addon"`
	OutPut  string   `env:"OutPut" required:"true" default:"./addon/"`
	Root    string   `env:"Root" required:"true" default:"./"`
	Sources []string `env:"Sources" required:"true" default:""`
	Target  string   `env:"Target" required:"true" default:"addon.cc"`
	Exports []Export `env:"Exports" required:"true" default:""`
}
