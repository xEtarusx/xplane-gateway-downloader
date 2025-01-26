package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xEtarusx/xplane-gateway-downloader/config"
	"github.com/xEtarusx/xplane-gateway-downloader/downloader"
	"os"
)

func ActionConfig(c *cli.Context) error {
	customSceneryFolder := c.String("custom-scenery-folder")
	xPlaneVersion := c.String("x-plane-version")

	if customSceneryFolder != "" {
		handleCustomSceneryFolder(customSceneryFolder)
	}

	if xPlaneVersion != "" {
		handleXPlaneVersion(xPlaneVersion)
	}

	return nil
}

func ConfigChangePresent(c *cli.Context) bool {
	if c.String("custom-scenery-folder") != "" {
		return true
	}

	if c.String("x-plane-version") != "" {
		return true
	}

	return false
}

func handleCustomSceneryFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/custom-scenery-folder does not exist
		fmt.Println("The path does not exist. Please make sure that it's pointing to your CustomScenery folder in the X-Plane game folder and ends with a slash")
		return
	}

	// Persist the CustomScenery folder path in config.json
	config.GlobalConfig.CustomSceneryFolder = path
}

func handleXPlaneVersion(version string) {
	releases, err := downloader.GetReleaseData()
	if err != nil {
		fmt.Println("Could not download release list from gateway")
		fmt.Println(err)
		return
	}

	if len(releases) == 0 {
		fmt.Println("No releases found on gateway")
		return
	}

	for _, release := range releases {
		if release.Version != version {
			continue
		}

		// Persist the X-Plane version in config.json
		config.GlobalConfig.XPlaneVersion = version
		fmt.Printf("Version X-Plane %s successfully saved\n", version)

		// Download list of all sceneries released with version
		fmt.Print("Downloading list of released sceneries ... ")
		sceneriesList, err := downloader.GetReleaseSceneryData(version)
		if err != nil {
			fmt.Println("failed")
			fmt.Printf("Could not download list of released sceneries from gateway for version %s\n", version)
			fmt.Println(err)
			return
		}

		config.GlobalConfig.ReleasedSceneryPacksWithVersion[version] = sceneriesList
		fmt.Println("done")

		return
	}

	fmt.Printf("The version %s is invalid\n\n", version)
	fmt.Println("Valid versions:")

	for _, release := range releases {
		fmt.Printf("%s (released %s)\n", release.Version, release.Date.Format("2006-01-02"))
	}
}
