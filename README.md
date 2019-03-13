[![GoDoc](http://godoc.org/github.com/borderstech/romanus?status.png)](http://godoc.org/github.com/borderstech/romanus)
[![Build Status](https://travis-ci.org/borderstech/romanus.svg?branch=master)](https://travis-ci.org/borderstech/romanus)
[![Go Report Card](https://goreportcard.com/badge/github.com/borderstech/romanus)](https://goreportcard.com/report/github.com/borderstech/romanus)
[![codecov](https://codecov.io/gh/borderstech/romanus/branch/master/graph/badge.svg)](https://codecov.io/gh/borderstech/romanus)

# romanus

Golang library containing the entire Roman Catechism of the Council of Trent.

Documentation here: https://godoc.org/github.com/borderstech/romanus

## Example Usage

```go
// Create a new catechism instance
c := romanus.NewCatechism()

// Get Part 1
part, err := c.GetPart(1)

// Get Part 1, Article 1
a, err := c.GetArticle(1, 1)

// Get Part 1, Article 1, Section 1
s, err := c.GetSection(1, 1, 1)

// Get Part 1, Article 1, Section 1, Paragraph 1
p, err := c.GetParagraph(1, 1, 1, 1)
fmt.Print(p.Text) // English

// Search for paragraphs
r := c.Search("I believe", 10)
fmt.Print(r[0].Part.PartNumber)
fmt.Print(r[0].Article.ArticleNumber)
fmt.Print(r[0].Section.SectionNumber)
fmt.Print(r[0].Paragraph.ParagraphNumber)
fmt.Print(r[0].Paragraph.Text)
fmt.Print(r[0].String()) // Part 4, Article 3, Section 4, Paragraph 1
```
