package app

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xEtarusx/xplane-gateway-downloader/config"
	"log"
	"os"
)

type Action struct {
}

func (a *Action) Process(c *cli.Context, callback func(c *cli.Context) error) error {
	configFileLocation := c.String("config")

	// Config file not found
	if _, err := os.Stat(configFileLocation); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Config file '%s' not found. Download the default config.json from the project and put it next to the binary.", configFileLocation)
		return nil
	}

	// Load config.json
	cfg, err := config.LoadConfig(configFileLocation)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Save config.json content into GlobalConfig
	config.GlobalConfig = cfg

	// Call the callback function for all logic
	err = callback(c)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Save GlobalConfig back into config.json
	err = config.SaveConfig(configFileLocation)
	if err != nil {
		log.Println(err)
		return nil
	}

	return nil
}
