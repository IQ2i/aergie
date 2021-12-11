package file

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"
)

// Exists checks if a given filepath is an existing file and not a directory
func Exists(filepath string) bool {

	fileinfo, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return false
	}

	return !fileinfo.IsDir()
}

func Download(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func Uncompress(source string, dest string) error {

	gzipfile, err := os.Open(source)

	if err != nil {
		return err
	}

	reader, err := gzip.NewReader(gzipfile)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dest)

	if err != nil {
		return err
	}

	defer writer.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}

	return nil
}

func Remove(source string) error {

	if err := os.Remove(source); err != nil {
		return err
	}

	return nil
}
