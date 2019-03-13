package romanus

import (
	"errors"
	"github.com/derekparker/trie"
)

// Catechism contains a list of parts
type Catechism struct {
	Parts       []Part        `json:"-"`
	partSummary []PartSummary `json:"parts"`
	searchTree  *trie.Trie
}

// Part contains a list of articles
type Part struct {
	Title      string    `json:"title"`
	PartNumber uint8     `json:"partNumber"`
	Articles   []Article `json:"articles"`
}

// Article contains a list of sections
type Article struct {
	Title         string    `json:"title"`
	ArticleNumber uint8     `json:"articleNumber"`
	Sections      []Section `json:"sections"`
}

// Section contains a list of paragraphs
type Section struct {
	Title         string      `json:"title"`
	SectionNumber uint8       `json:"sectionNumber"`
	Paragraphs    []Paragraph `json:"paragraphs"`
}

// Paragraph contains text
type Paragraph struct {
	ParagraphNumber uint8  `json:"paragraphNumber"`
	Text            string `json:"text"`
}

// PartSummary contains metadata for a part
type PartSummary struct {
	Title      string           `json:"title"`
	PartNumber uint8            `json:"partNumber"`
	Articles   []ArticleSummary `json:"articles"`
}

// ArticleSummary contains metadata for an article
type ArticleSummary struct {
	Title         string `json:"title"`
	ArticleNumber uint8  `json:"articleNumber"`
}

// GetPart obtains a part within the catechism by its number
func (c *Catechism) GetPart(partNumber int) (*Part, error) {
	idx := partNumber - 1
	if idx < 0 || idx >= len(c.Parts) {
		return nil, errors.New("invalid part number")
	}

	return &c.Parts[idx], nil
}

// GetArticle obtains an article within the catechism by its number
func (c *Catechism) GetArticle(partNumber, articleNumber int) (*Article, error) {
	p, err := c.GetPart(partNumber)
	if err != nil {
		return nil, err
	}

	idx := articleNumber - 1
	if idx < 0 || idx >= len(p.Articles) {
		return nil, errors.New("invalid article number")
	}

	return &p.Articles[idx], nil
}

// GetSection obtains a section within the catechism by its number
func (c *Catechism) GetSection(partNumber, articleNumber, sectionNumber int) (*Section, error) {
	a, err := c.GetArticle(partNumber, articleNumber)
	if err != nil {
		return nil, err
	}

	idx := sectionNumber - 1
	if idx < 0 || idx >= len(a.Sections) {
		return nil, errors.New("invalid section number")
	}

	return &a.Sections[idx], nil
}

// GetParagraph obtains a paragraph within the catechism by its number
func (c *Catechism) GetParagraph(partNumber, articleNumber, sectionNumber, paragraphNumber int) (*Paragraph, error) {
	s, err := c.GetSection(partNumber, articleNumber, sectionNumber)
	if err != nil {
		return nil, err
	}

	idx := paragraphNumber - 1
	if idx < 0 || idx >= len(s.Paragraphs) {
		return nil, errors.New("invalid paragraph number")
	}

	return &s.Paragraphs[idx], nil
}
