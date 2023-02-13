package types

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"time"
)

type Scenery struct {
	Id            int             `json:"sceneryId"`
	AirportName   string          `json:"aptName"`
	DateApproved  time.Time       `json:"dateApproved"`
	Type          string          `json:"type"`
	Status        string          `json:"status"`
	MasterZipBlob string          `json:"masterZipBlob"`
	Metadata      SceneryMetadata `json:"additionalMetadata"`
}

func (s Scenery) decodeMasterZipBlob() ([]byte, error) {
	return base64.StdEncoding.DecodeString(s.MasterZipBlob)
}

// ExtractSceneryPackZip Extract bytes from {icao}_Scenery_Pack.zip from inside another zip file
func (s Scenery) ExtractSceneryPackZip(icao string) ([]byte, error) {
	blob, err := s.decodeMasterZipBlob()
	if err != nil {
		return nil, err
	}

	// create a zip reader
	zipReader, err := zip.NewReader(bytes.NewReader(blob), int64(len(blob)))
	if err != nil {
		return []byte{}, err
	}

	// zip structure:
	// |- {icao}.dat
	// |- {icao}.txt
	// |- {icao}_Scenery_Pack.zip

	// loop through all the files from the zip archive
	for _, zipFile := range zipReader.File {
		// we look for the {icao}_Scenery_Pack.zip inside the zip file, other files can be ignored
		if zipFile.Name != icao+"_Scenery_Pack.zip" {
			continue
		}

		return readZipFile(zipFile)
	}

	return []byte{}, errors.New("zip file was empty")
}

// ReadZipFile Source: https://stackoverflow.com/questions/50539118/golang-unzip-response-body
func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
