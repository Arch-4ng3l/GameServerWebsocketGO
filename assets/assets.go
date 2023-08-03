package assets

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Init(homeDir string) error {

	archive, err := os.Create("assets.zip")

	if err != nil {
		return err
	}
	path := homeDir + "/Hololens/assets/"
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(archive)

	defer archive.Close()
	defer zipWriter.Close()
	for _, file := range files {
		name := file.Name()

		if strings.HasSuffix(name, ".go") {
			continue
		}

		w, err := zipWriter.Create("assets/" + name)
		if err != nil {
			return err
		}

		f, err := os.Open(path + name)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}

	return nil
}
