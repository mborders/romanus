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
	"strings"

	"github.com/derekparker/trie"
)

type SearchNode struct {
	Part      *Part
	Article   *Article
	Section   *Section
	Paragraph *Paragraph
}

// String creates a string representation for the SearchNode,
// ex. Part 1, Article 1, Section 1, Paragraph 1
func (s *SearchNode) String() string {
	return fmt.Sprintf("Part %d, Article %d, Section %d, Paragraph %d",
		s.Part.PartNumber,
		s.Article.ArticleNumber,
		s.Section.SectionNumber,
		s.Paragraph.ParagraphNumber);
}

const catechismTar = "catechism.tar.gz"
const catechismFilename = "catechism.json"

func NewCatechism() *Catechism {
	_, currFile, _, _ := runtime.Caller(0)
	filename := fmt.Sprintf("%s/%s", path.Dir(currFile), catechismTar)

	catechism := &Catechism{
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
			catechism.Parts = decode(tr)
		}
	}

	// Build search tree
	for i := range catechism.Parts {
		p := catechism.Parts[i]
		for j := range p.Articles {
			a := p.Articles[j]
			for k := range a.Sections {
				s := a.Sections[k]
				for l := range s.Paragraphs {
					par := s.Paragraphs[l]
					t := strings.ToLower(par.Text)
					catechism.searchTree.Add(t, SearchNode{
						Part:      &p,
						Article:   &a,
						Section:   &s,
						Paragraph: &par,
					})
				}
			}
		}
	}

	return catechism
}

func decode(r io.Reader) []Part {
	var parts []Part
	json.NewDecoder(r).Decode(&parts)
	return parts
}

// Search finds top matching verses based on the given query.
// The number of search results are restricted by maxResults
func (c *Catechism) Search(query string, maxResults int) []SearchNode {
	t := c.searchTree
	keys := t.FuzzySearch(strings.ToLower(query))
	var nodes []SearchNode

	for k := range keys {
		res, _ := t.Find(keys[k])
		nodes = append(nodes, res.Meta().(SearchNode))

		if len(nodes) >= maxResults {
			break
		}
	}

	return nodes
}
