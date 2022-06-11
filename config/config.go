package config

import (
	"encoding/json"
	"fmt"
	"github.com/xEtarusx/xplane-gateway-downloader/types"
	"io/ioutil"
	"os"
)

type Config struct {
	CustomSceneryFolder             string                   `json:"custom-scenery-folder"`
	XPlaneVersion                   string                   `json:"x-plane-version"`
	AirportConfig                   map[string]types.Airport `json:"airports"`
	ReleasedSceneryPacksWithVersion map[string][]int         `json:"released-scenery-packs-with-version"`
}

var GlobalConfig Config

func LoadConfig(file string) (Config, error) {
	var config Config

	configFile, err := os.Open(file)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}

// SaveConfig Save the config into a file
func SaveConfig(file string) error {
	content, err := json.MarshalIndent(GlobalConfig, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, content, os.ModePerm)
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
		fmt.Printf("Could not find list of released scenery packs for version %s\n", c.XPlaneVersion)
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
