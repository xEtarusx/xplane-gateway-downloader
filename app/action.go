package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xEtarusx/xplane-gateway-downloader/config"
)

type Action struct {
}

func (a *Action) Process(c *cli.Context, callback func(c *cli.Context) error) error {
	configFileLocation := c.String("config")

	// Load config.json
	err := config.LoadConfig(configFileLocation)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Validate config if the user does not want to change the config
	if !ConfigChangePresent(c) {
		if err := config.GlobalConfig.Valid(); err != nil {
			fmt.Println(err)
			return nil
		}
	}

	// Call the callback function for all logic
	err = callback(c)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Save GlobalConfig back into config.json
	err = config.SaveConfig(configFileLocation)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
