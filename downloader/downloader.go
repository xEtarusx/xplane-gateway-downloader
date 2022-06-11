package downloader

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xEtarusx/xplane-gateway-downloader/types"
	"io/ioutil"
	"net/http"
	"strconv"
)

var airportCache = map[string]types.Airport{}
var sceneryCache = map[int]types.Scenery{}

func GetAirportData(icao string) (types.Airport, error) {
	if icao == "" {
		fmt.Println("")
		return types.Airport{}, errors.New("icao code cannot be empty")
	}

	// return airport from cache
	if airport, ok := airportCache[icao]; ok {
		return airport, nil
	}

	url := "https://gateway.x-plane.com/apiv1/airport/" + icao

	response, err := http.Get(url)
	if err != nil {
		return types.Airport{}, err
	}
	defer response.Body.Close()

	// Parse the json
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return types.Airport{}, err
	}

	var airportResult types.AirportResult
	err = json.Unmarshal(body, &airportResult)
	if err != nil {
		return types.Airport{}, err
	}

	// cache airport for later requests
	airportCache[icao] = airportResult.Airport

	return airportResult.Airport, nil
}

func GetSceneryData(sceneryId int) (types.Scenery, error) {
	url := "https://gateway.x-plane.com/apiv1/scenery/" + strconv.Itoa(sceneryId)

	// return scenery from cache
	if scenery, ok := sceneryCache[sceneryId]; ok {
		return scenery, nil
	}

	response, err := http.Get(url)
	if err != nil {
		return types.Scenery{}, err
	}
	defer response.Body.Close()

	// Parse the json
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return types.Scenery{}, err
	}

	var sceneryResult types.SceneryResult
	err = json.Unmarshal(body, &sceneryResult)
	if err != nil {
		return types.Scenery{}, err
	}

	// cache scenery for later requests
	sceneryCache[sceneryId] = sceneryResult.Scenery

	return sceneryResult.Scenery, nil
}

func GetReleaseData() ([]types.Release, error) {
	url := "https://gateway.x-plane.com/apiv1/releases"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the json
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var releases []types.Release
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func GetReleaseSceneryData(release string) ([]int, error) {
	url := "https://gateway.x-plane.com/apiv1/release/" + release

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the json
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var sceneryPacks types.SceneryPacksResult
	err = json.Unmarshal(body, &sceneryPacks)
	if err != nil {
		return nil, err
	}

	return sceneryPacks.SceneryPacks, nil
}
