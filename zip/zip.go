// Package zip implements the Archive interface providing zip archiving
// and compression.
package zip

import (
	"archive/zip"
	"io"
	"os"
)

// Archive zip struct
type Archive struct {
	z *zip.Writer
}

// Close all closeables
func (a Archive) Close() error {
	return a.z.Close()
}

// New zip archive
func New(target *os.File) Archive {
	return Archive{
		z: zip.NewWriter(target),
	}
}

// Add a file to the zip archive
func (a Archive) Add(name, path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close() // nolint: errcheck
	stat, err := file.Stat()
	if err != nil {
		return
	}
	if stat.IsDir() {
		return
	}
	w, err := a.z.Create(name)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, file)
	return err
}
