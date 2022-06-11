package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xEtarusx/xplane-gateway-downloader/config"
	"os"
	"path"
	"strings"
)

func ActionUninstall(c *cli.Context) error {
	icao := strings.ToUpper(c.String("icao"))

	if icao == "" {
		fmt.Println("--icao parameter cannot be empty")
		return nil
	}

	// Check if airport is not installed locally
	if !config.GlobalConfig.IsAirportInstalled(icao) {
		fmt.Printf("Airport %s is not installed", icao)
		return nil
	}

	// delete scenery folder
	err := os.RemoveAll(path.Join(config.GlobalConfig.CustomSceneryFolder, GetSceneryPackFolderName(icao)))
	if err != nil {
		return err
	}

	// delete airport from local config
	config.GlobalConfig.AirportConfig = config.GlobalConfig.RemoveAirport(icao)

	return nil
}
