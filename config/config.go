package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xEtarusx/xplane-gateway-downloader/types"
	"os"
)

type Config struct {
	CustomSceneryFolder             string                   `json:"custom-scenery-folder"`
	XPlaneVersion                   string                   `json:"x-plane-version"`
	AirportConfig                   map[string]types.Airport `json:"airports"`
	ReleasedSceneryPacksWithVersion map[string][]int         `json:"released-scenery-packs-with-version"`
}

var GlobalConfig Config

// createEmptyConfig creates an empty GlobalConfig and saves it to file
func createEmptyConfig(file string) {
	GlobalConfig = Config{
		CustomSceneryFolder:             "",
		XPlaneVersion:                   "",
		AirportConfig:                   map[string]types.Airport{},
		ReleasedSceneryPacksWithVersion: map[string][]int{},
	}

	SaveConfig(file)
}

// LoadConfig loads the config from config.json if available or creates a new config.json file
func LoadConfig(file string) error {
	// If there is no config file in the given location, create a new one
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Config file '%s' not found. Creating empty one.\n", file)
		createEmptyConfig(file)
		return nil
	}

	var config Config

	configFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	// Save config.json content into GlobalConfig
	GlobalConfig = config

	return err
}

// SaveConfig Save the config into a file
func SaveConfig(file string) error {
	content, err := json.MarshalIndent(GlobalConfig, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, content, os.ModePerm)
}

func (c Config) Valid() error {
	// Check if custom scenery folder was set
	if c.CustomSceneryFolder == "" {
		return errors.New("X-Plane custom scenery folder is not set. Use 'xplane-gateway-downloader config -csf \"D:\\path-to\\X-Plane 12\\Custom Scenery\"' to set the path")
	}

	// Check if x-plane version was set
	if c.XPlaneVersion == "" {
		return errors.New("X-Plane version is not set. Use 'xplane-gateway-downloader config -v XX.XX' to set the version")
	}

	return nil
}

// IsAirportInstalled Check if an airport is installed with the provided icao code
func (c Config) IsAirportInstalled(icao string) bool {
	_, found := c.AirportConfig[icao]
	return found
}

func (c Config) IsXPlaneVersionSet() bool {
	if c.XPlaneVersion != "" {
		return true
	}
	return false
}

func (c Config) IsSceneryPackIncluded(sceneryId int) bool {
	if _, ok := c.ReleasedSceneryPacksWithVersion[c.XPlaneVersion]; !ok {
		fmt.Printf("Could not find list of released scenery packs for version %s. Please set the version via the config command.\n", c.XPlaneVersion)
		return false
	}

	for _, id := range c.ReleasedSceneryPacksWithVersion[c.XPlaneVersion] {
		if id == sceneryId {
			return true
		}
	}

	return false
}

// SaveAirport Save a new airport in the local config
func (c Config) SaveAirport(airport types.Airport) {
	c.AirportConfig[airport.ICAO] = airport
}

// RemoveAirport Removes an airport by icao code and return the new map
func (c Config) RemoveAirport(icao string) map[string]types.Airport {
	newAirports := make(map[string]types.Airport)

	for keyIcao, airport := range c.AirportConfig {
		// Do not save the provided icao airport
		if icao == keyIcao {
			continue
		}

		newAirports[keyIcao] = airport
	}

	return newAirports
}
