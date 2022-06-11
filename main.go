package main

import (
	"github.com/urfave/cli/v2"
	"github.com/xEtarusx/xplane-gateway-downloader/app"
	"github.com/xEtarusx/xplane-gateway-downloader/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config.GlobalConfig = cfg

	a := &cli.App{
		Name:        "X-Plane Gateway Downloader",
		HelpName:    "xplane-gateway-downloader",
		Usage:       "Download airports from X-Plane Gateway with ease",
		Description: "Update airport sceneries from the X-Plane Gateway",
		Commands: []*cli.Command{
			{
				Name:   "install",
				Usage:  "Install a new airport scenery pack",
				Action: app.ActionInstall,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "icao",
						Aliases: []string{"i"},
						Usage:   "Install an airport by `ICAO` code",
					},
				},
			},
			{
				Name:   "update",
				Usage:  "Update all installed airport scenery packs",
				Action: app.ActionUpdate,
			},
			{
				Name:   "uninstall",
				Usage:  "Uninstall an installed airport scenery pack",
				Action: app.ActionUninstall,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "icao",
						Aliases: []string{"i"},
						Usage:   "Uninstall an airport by `ICAO` code",
					},
				},
			},
			{
				Name:   "config",
				Usage:  "Configure the application",
				Action: app.ActionConfig,
				Flags: []cli.Flag{
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

	err = a.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	err = config.SaveConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
}
