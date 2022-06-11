package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xEtarusx/xplane-gateway-downloader/config"
	"github.com/xEtarusx/xplane-gateway-downloader/downloader"
	"os"
	"path"
	"strings"
)

func ActionInstall(c *cli.Context) error {
	icao := strings.ToUpper(c.String("icao"))

	if icao == "" {
		fmt.Println("--icao parameter cannot be empty")
		return nil
	}

	// Check if airport is already installed locally
	if config.GlobalConfig.IsAirportInstalled(icao) {
		fmt.Printf("Airport %s already installed", icao)
		return nil
	}

	// Check if the folder {ICAO}_Scenery_Pack already exist in CustomScenery folder
	if _, err := os.Stat(path.Join(config.GlobalConfig.CustomSceneryFolder, GetSceneryPackFolderName(icao))); !os.IsNotExist(err) {
		fmt.Printf("The folder %s already exists in your CustomScenery folder but not in the local config. Aborting\n", GetSceneryPackFolderName(icao))
		return nil
	}

	airport, err := downloader.GetAirportData(icao)
	if err != nil {
		return err
	}

	// Check if the user has set their x-plane version
	if config.GlobalConfig.IsXPlaneVersionSet() {
		// Check if the airport scenery is already included in the x-plane release
		if config.GlobalConfig.IsSceneryPackIncluded(airport.RecommendedSceneryId) {
			fmt.Printf("There is no newer version of this airport available. X-Plane %v contains the latest update.\n", config.GlobalConfig.XPlaneVersion)
			return nil
		}
	}

	scenery, err := downloader.GetSceneryData(airport.RecommendedSceneryId)
	if err != nil {
		return err
	}

	// extract the {ICAO}_Scenery_Pack.zip form the downloaded zip archive
	sceneryPackBytes, err := scenery.ExtractSceneryPackZip(icao)
	if err != nil {
		return err
	}

	// delete old scenery folder
	err = os.RemoveAll(path.Join(config.GlobalConfig.CustomSceneryFolder, GetSceneryPackFolderName(icao)))
	if err != nil {
		return err
	}

	// create new scenery folder
	err = WriteZipToFolder(sceneryPackBytes, config.GlobalConfig.CustomSceneryFolder)
	if err != nil {
		return err
	}

	fmt.Printf("%s was successfully installed\n", airport.ICAO)

	// Store the scenery approved date for information purpose in the airport
	airport.SceneryApprovedDate = scenery.DateApproved

	// Store airport in config
	config.GlobalConfig.SaveAirport(airport)

	return nil
}
