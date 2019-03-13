package romanus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var c = NewCatechism()

func TestCatechism_GetPart(t *testing.T) {
	part, err := c.GetPart(1)
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), part.PartNumber)
}

func TestCatechism_GetPart_Invalid(t *testing.T) {
	_, err := c.GetPart(0)
	assert.Equal(t, "invalid part number", err.Error())
}

func TestCatechism_GetArticle(t *testing.T) {
	a, err := c.GetArticle(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), a.ArticleNumber)
}

func TestCatechism_GetArticle_InvalidPart(t *testing.T) {
	_, err := c.GetArticle(0, 100)
	assert.Equal(t, "invalid part number", err.Error())
}

func TestCatechism_GetArticle_InvalidArticle(t *testing.T) {
	_, err := c.GetArticle(1, 100)
	assert.Equal(t, "invalid article number", err.Error())
}

func TestCatechism_GetSection(t *testing.T) {
	s, err := c.GetSection(1, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), s.SectionNumber)
}

func TestCatechism_GetSection_InvalidPart(t *testing.T) {
	_, err := c.GetSection(0, 100, 100)
	assert.Equal(t, "invalid part number", err.Error())
}

func TestCatechism_GetSection_InvalidArticle(t *testing.T) {
	_, err := c.GetSection(1, 100, 100)
	assert.Equal(t, "invalid article number", err.Error())
}

func TestCatechism_GetSection_InvalidSection(t *testing.T) {
	_, err := c.GetSection(1, 1, 100)
	assert.Equal(t, "invalid section number", err.Error())
}

func TestCatechism_GetParagraph(t *testing.T) {
	p, err := c.GetParagraph(1, 1, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), p.ParagraphNumber)
}

func TestCatechism_GetParagraph_InvalidPart(t *testing.T) {
	_, err := c.GetParagraph(0, 100, 100, 100)
	assert.Equal(t, "invalid part number", err.Error())
}

func TestCatechism_GetParagraph_InvalidArticle(t *testing.T) {
	_, err := c.GetParagraph(1, 100, 100, 100)
	assert.Equal(t, "invalid article number", err.Error())
}

func TestCatechism_GetParagraph_InvalidSection(t *testing.T) {
	_, err := c.GetParagraph(1, 1, 100, 100)
	assert.Equal(t, "invalid section number", err.Error())
}

func TestCatechism_GetParagraph_InvalidParagraph(t *testing.T) {
	_, err := c.GetParagraph(1, 1, 1, 100)
	assert.Equal(t, "invalid paragraph number", err.Error())
}
