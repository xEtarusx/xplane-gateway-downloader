package app

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func GetSceneryPackFolderName(icao string) string {
	return icao + "_Scenery_Pack"
}

// WriteZipToFolder Write bytes from a zip archive into destination path
func WriteZipToFolder(zipContent []byte, destinationPath string) error {
	// create a zip reader
	zipReader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	if err != nil {
		return err
	}

	// loop through all the files of the zip archive
	for _, zipFile := range zipReader.File {
		fpath := filepath.Join(destinationPath, zipFile.Name)

		if zipFile.FileInfo().IsDir() {
			// Make Folder
			_ = os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
		if err != nil {
			return err
		}

		rc, err := zipFile.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		_ = outFile.Close()
		_ = rc.Close()
	}

	return nil
}
