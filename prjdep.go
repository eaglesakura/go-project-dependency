package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp();
	app.Name = "prjdep / Project Dependency";
	app.Usage = "Project dependency sync tool";
	app.Version = "0.0.2";
	app.Commands = []cli.Command{
		{
			Name:"init",
			Usage:"Initialize repository dependencies to `dependencies.json`",
			Action:cmdInit,
		},
		{
			Name:"restore",
			Usage:"Restore repository dependencies from `dependencies.json`",
			Action:cmdRestore,
		},
	};
	app.Run(os.Args);
}

const dependenciesFileName = "dependencies.json";

// `prjdep init` コマンドのハンドリングを行なう
func cmdInit(ctx *cli.Context) error {
	dependencies, err := NewDependencies();
	if err != nil {
		return err;
	}

	return dependencies.ToFile(dependenciesFileName);
}

// `prjdep restore` コマンドのハンドリングを行なう
func cmdRestore(ctx *cli.Context) error {
	dependencies, err := NewDependenciesFromFile(dependenciesFileName);
	if err != nil {
		return err;
	}
	return dependencies.Restore();
}