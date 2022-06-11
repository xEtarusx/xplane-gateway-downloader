package types

import "time"

type Airport struct {
	ICAO                 string    `json:"icao"`
	Name                 string    `json:"airportName"`
	RecommendedSceneryId int       `json:"recommendedSceneryId"`
	SceneryApprovedDate  time.Time `json:"sceneryApprovedDate"`
}
