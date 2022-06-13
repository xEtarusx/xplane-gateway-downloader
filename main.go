package main

import (
	"github.com/urfave/cli/v2"
	"github.com/xEtarusx/xplane-gateway-downloader/app"
	"log"
	"os"
)

func main() {
	action := app.Action{}

	a := &cli.App{
		Name:        "X-Plane Gateway Downloader",
		HelpName:    "xplane-gateway-downloader",
		Usage:       "Download airports from X-Plane Gateway with ease",
		Description: "Update airport sceneries from the X-Plane Gateway",
		Commands: []*cli.Command{
			{
				Name:  "install",
				Usage: "Install a new airport scenery pack",
				Action: func(c *cli.Context) error {
					return action.Process(c, app.ActionInstall)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "The `path` to the config.json",
						Value:   "config.json",
					},
					&cli.StringFlag{
						Name:    "icao",
						Aliases: []string{"i"},
						Usage:   "Install an airport by `ICAO` code",
					},
				},
			},
			{
				Name:  "update",
				Usage: "Update all installed airport scenery packs",
				Action: func(c *cli.Context) error {
					return action.Process(c, app.ActionUpdate)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "The `path` to the config.json",
						Value:   "config.json",
					},
				},
			},
			{
				Name:  "uninstall",
				Usage: "Uninstall an installed airport scenery pack",
				Action: func(c *cli.Context) error {
					return action.Process(c, app.ActionUninstall)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "The `path` to the config.json",
						Value:   "config.json",
					},
					&cli.StringFlag{
						Name:    "icao",
						Aliases: []string{"i"},
						Usage:   "Uninstall an airport by `ICAO` code",
					},
				},
			},
			{
				Name:  "config",
				Usage: "Configure the application",
				Action: func(c *cli.Context) error {
					return action.Process(c, app.ActionConfig)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "The `path` to the config.json",
						Value:   "config.json",
					},
					&cli.StringFlag{
						Name:    "custom-scenery-folder",
						Aliases: []string{"csf"},
						Usage:   "The `path` to CustomScenery folder of x-plane",
					},
					&cli.StringFlag{
						Name:    "x-plane-version",
						Aliases: []string{"v"},
						Usage:   "Set the current `version` of x-plane",
					},
				},
			},
		},
	}

	err := a.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
