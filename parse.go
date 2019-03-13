package romanus

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/derekparker/trie"
)

const catechismTar = "catechism.tar.gz"
const catechismFilename = "catechism.json"

func NewCatechism() *Catechism {
	_, currFile, _, _ := runtime.Caller(0)
	filename := fmt.Sprintf("%s/%s", path.Dir(currFile), catechismTar)

	c := &Catechism{
		searchTree: trie.New(),
	}

	f, _ := os.Open(filename)
	defer f.Close()

	gzf, _ := gzip.NewReader(f)
	defer gzf.Close()

	tr := tar.NewReader(gzf)

	// Build testaments, books, chapters, verses
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}

		switch h.Name {
		case catechismFilename:
			c.Parts = decode(tr)
		}
	}

	return c
}

func decode(r io.Reader) []Part {
	var parts []Part
	json.NewDecoder(r).Decode(&parts)
	return parts
}
