package main

import (
	"github.com/urfave/cli"
	"os"
	"./repo"
)

func main() {
	app := cli.NewApp();
	app.Name = "nova";
	app.Usage = "Golang dependency tool";
	app.Version = "0.0.1";
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
	dependencies, err := repo.NewDependencies();
	if err != nil {
		return err;
	}

	return dependencies.ToFile(dependenciesFileName);
}

// `prjdep restore` コマンドのハンドリングを行なう
func cmdRestore(ctx *cli.Context) error {
	dependencies, err := repo.FromFile(dependenciesFileName);
	if err != nil {
		return err;
	}
	return dependencies.Restore();
}