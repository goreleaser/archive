// Package tar implements the Archive interface providing tar.gz archiving
// and compression.
package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

// Archive as tar.gz
type Archive struct {
	gw *gzip.Writer
	tw *tar.Writer
}

// Close all closeables
func (a Archive) Close() error {
	if err := a.tw.Close(); err != nil {
		return err
	}
	return a.gw.Close()
}

// New tar.gz archive
func New(target *os.File) Archive {
	gw := gzip.NewWriter(target)
	tw := tar.NewWriter(gw)
	return Archive{
		gw: gw,
		tw: tw,
	}
}

// Add file to the archive
func (a Archive) Add(name, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close() // nolint: errcheck
	stat, err := file.Stat()
	if err != nil || stat.IsDir() {
		return err
	}
	header, err := tar.FileInfoHeader(stat, name)
	if err != nil {
		return err
	}
	if err := a.tw.WriteHeader(header); err != nil {
		return err
	}
	_, err = io.Copy(a.tw, file)
	return err
}
