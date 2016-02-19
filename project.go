package main

import (
	"os"

	"github.com/nildev/lib/log"
	"github.com/codegangsta/cli"
	"github.com/nildev/project/setup"
	"github.com/nildev/project/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "project"
	app.Usage = "Tool for project setup and maintainance"
	app.Version = version.Version
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "verbosity",
			Value: 0,
			Usage: "logging level",
		},
	}
	app.Before = func(c *cli.Context) error {
		// setup logging here
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "setup",
			Aliases: []string{"s"},
			Usage:   "setup new project based on given template",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "destDir",
					Value: "",
					Usage: "full path to directory where project should be initialized",
				},
				cli.StringFlag{
					Name:  "configFile",
					Value: "",
					Usage: "path to config file required to setup project, temmplate repo should provide sample one",
				},
				cli.StringFlag{
					Name:  "templateRepo",
					Value: "github.com/nildev/goodbey-world-template",
					Usage: "full path to template repo, for example: github.com/nildev/goodbey-world-template",
				},
				cli.BoolFlag{
					Name:  "override",
					Usage: "if destDir exists override its content",
				},
			},
			Action: func(c *cli.Context) {
				if c.String("destDir") == "" {
					log.Fatalf("Please provide path to configFile")
				}
				if c.String("configFile") == "" {
					log.Fatalf("Please provide path to configFile")
				}

				if c.String("templateRepo") == "" {
					log.Fatalf("Please provide path to templateRepo")
				}

				init := setup.NewInitializer()
				cfgLoader := setup.NewConfigLoader(c.String("configFile"))
				cfg := cfgLoader.Read(c.String("configFile"))

				init.Setup(cfg, c.String("destDir"), c.String("templateRepo"))
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Debugf(" Could not run project cmd, %s", err)
	}
}
