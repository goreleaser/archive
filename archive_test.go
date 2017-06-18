package archive

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArchive(t *testing.T) {
	var assert = assert.New(t)

	folder, err := ioutil.TempDir("", "archivetest")
	assert.NoError(err)

	empty, err := os.Create(folder + "/empty.txt")
	assert.NoError(err)

	assert.NoError(os.Mkdir(folder+"/folder-inside", 0755))

	for _, archive := range []Archive{
		newTar(folder, t),
		newZip(folder, t),
	} {
		assert.NoError(archive.Add("empty.txt", empty.Name()))
		assert.NoError(archive.Add("empty.txt", folder+"/folder-inside"))
		assert.Error(archive.Add("dont.txt", empty.Name()+"_nope"))
		assert.NoError(archive.Close())
	}
}

func newZip(folder string, t *testing.T) Archive {
	file, err := os.Create(folder + "/folder.zip")
	assert.NoError(t, err)
	return NewZip(file)
}

func newTar(folder string, t *testing.T) Archive {
	file, err := os.Create(folder + "/folder.tar.gz")
	assert.NoError(t, err)
	return NewTargz(file)
}
