// Package archive provides tar.gz and zip archiving
package archive

import (
	"os"

	"github.com/goreleaser/archive/tar"
	"github.com/goreleaser/archive/zip"
)

// Archive represents a compression archive files from disk can be written to.
type Archive interface {
	Close() error
	Add(name, path string) error
}

// NewZip returns an archive instance capable of compressing in zip format
func NewZip(file *os.File) Archive {
	return zip.New(file)
}

// NewTargz returns an archive instance capable of compressing in tar.gz format
func NewTargz(file *os.File) Archive {
	return tar.New(file)
}
