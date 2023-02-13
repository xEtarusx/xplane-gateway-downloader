package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xEtarusx/xplane-gateway-downloader/config"
	"github.com/xEtarusx/xplane-gateway-downloader/downloader"
	"os"
	"path"
)

func ActionUpdate(c *cli.Context) error {
	if len(config.GlobalConfig.AirportConfig) == 0 {
		fmt.Println("No airports installed")
		return nil
	}

	// Loop through all local airport sceneries
	for icao, localAirport := range config.GlobalConfig.AirportConfig {

		// Get the latest version from gateway for airport
		airport, err := downloader.GetAirportData(icao)
		if err != nil {
			return err
		}

		// Check if the recommendedSceneryId in the local config is the same as from the gateway
		if airport.RecommendedSceneryId == localAirport.RecommendedSceneryId {
			fmt.Printf("%s is already up to date\n", airport.ICAO)
			// Skip updating this airport, continue with next one
			continue
		}

		// Download the latest airport scenery pack
		scenery, err := downloader.GetSceneryData(airport.RecommendedSceneryId)
		if err != nil {
			return err
		}

		// extract the {ICAO}_Scenery_Pack.zip form the downloaded zip archive
		sceneryPackBytes, err := scenery.ExtractSceneryPackZip(icao)
		if err != nil {
			return err
		}

		// delete old scenery folder from CustomScenery folder
		err = os.RemoveAll(path.Join(config.GlobalConfig.CustomSceneryFolder, GetSceneryPackFolderName(icao)))
		if err != nil {
			return err
		}

		// create new scenery folder
		err = WriteZipToFolder(sceneryPackBytes, config.GlobalConfig.CustomSceneryFolder)
		if err != nil {
			return err
		}

		fmt.Printf("%s was successfully updated\n", airport.ICAO)

		// Store the scenery approved date for information purpose in the airport
		airport.SceneryApprovedDate = scenery.DateApproved

		// Store airport in config
		config.GlobalConfig.SaveAirport(airport)
	}

	return nil
}
