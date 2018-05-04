package zip

import (
	"archive/zip"
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
	f, err := os.Create(filepath.Join(tmp, "test.zip"))
	assert.NoError(err)
	defer f.Close()
	archive := New(f)

	assert.Error(archive.Add("nope.txt", "../testdata/nope.txt"))
	assert.NoError(archive.Add("foo.txt", "../testdata/foo.txt"))
	assert.NoError(archive.Add("sub1", "../testdata/sub1"))
	assert.NoError(archive.Add("sub1/bar.txt", "../testdata/sub1/bar.txt"))
	assert.NoError(archive.Add("sub1/executable", "../testdata/sub1/executable"))
	assert.NoError(archive.Add("sub1/sub2", "../testdata/sub1/sub2"))
	assert.NoError(archive.Add("sub1/sub2/subfoo.txt", "../testdata/sub1/sub2/subfoo.txt"))

	assert.NoError(archive.Close())
	assert.Error(archive.Add("tar.go", "tar.go"))
	assert.NoError(f.Close())

	t.Log(f.Name())
	f, err = os.Open(f.Name())
	assert.NoError(err)
	info, err := f.Stat()
	assert.NoError(err)
	assert.Truef(info.Size() < 600, "archived file should be smaller than %d", info.Size())
	r, err := zip.NewReader(f, info.Size())
	assert.NoError(err)
	var paths []string
	for _, zf := range r.File {
		paths = append(paths, zf.Name)
		t.Logf("%s: %v", zf.Name, zf.Mode())
		if zf.Name == "sub1/executable" {
			assert.Equal("-rwxr-xr-x", zf.Mode().String())
		}
	}
	assert.Equal([]string{
		"foo.txt",
		"sub1/bar.txt",
		"sub1/executable",
		"sub1/sub2/subfoo.txt",
	}, paths)
}
