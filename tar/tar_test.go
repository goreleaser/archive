package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTarGzFile(t *testing.T) {
	var assert = assert.New(t)
	tmp, err := ioutil.TempDir("", "")
	assert.NoError(err)
	f, err := os.Create(filepath.Join(tmp, "test.tar.gz"))
	assert.NoError(err)
	defer f.Close()
	archive := New(f)

	assert.Error(archive.Add("nope.txt", "../testdata/nope.txt"))
	assert.NoError(archive.Add("foo.txt", "../testdata/foo.txt"))
	assert.NoError(archive.Add("sub1", "../testdata/sub1"))
	assert.NoError(archive.Add("sub1/bar.txt", "../testdata/sub1/bar.txt"))
	assert.NoError(archive.Add("sub1/sub2", "../testdata/sub1/sub2"))
	assert.NoError(archive.Add("sub1/sub2/subfoo.txt", "../testdata/sub1/sub2/subfoo.txt"))

	assert.NoError(archive.Close())
	assert.Error(archive.Add("tar.go", "tar.go"))
	assert.NoError(f.Close())

	t.Log(f.Name())
	f, err = os.Open(f.Name())
	assert.NoError(err)
	gzf, err := gzip.NewReader(f)
	assert.NoError(err)
	var paths []string
	r := tar.NewReader(gzf)
	for {
		next, err := r.Next()
		if err == io.EOF {
			break
		}
		assert.NoError(err)
		paths = append(paths, next.Name)
	}
	assert.Equal([]string{
		"foo.txt",
		"sub1",
		"sub1/bar.txt",
		"sub1/sub2",
		"sub1/sub2/subfoo.txt",
	}, paths)
}
