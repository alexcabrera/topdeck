// Package parser handles parsing markdown presentation files into slides.
package parser

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Metadata from YAML frontmatter.
type Metadata struct {
	Title  string `yaml:"title"`
	Author string `yaml:"author"`
}

// Slide represents a single slide in the presentation.
type Slide struct {
	Content string
}

// Document is a parsed presentation.
type Document struct {
	Meta   Metadata
	Slides []Slide
}

// delimiterPattern matches horizontal rules (3+ dashes on their own line).
var delimiterPattern = regexp.MustCompile(`(?m)^-{3,}\s*$`)

// Parse parses a markdown presentation file into a Document.
func Parse(content []byte) (*Document, error) {
	var meta Metadata
	body := content

	// Check for frontmatter (starts with "---\n")
	if bytes.HasPrefix(content, []byte("---\n")) {
		rest := content[4:]
		// Find closing delimiter
		idx := bytes.Index(rest, []byte("\n---"))
		if idx != -1 {
			frontmatter := rest[:idx]
			if err := yaml.Unmarshal(frontmatter, &meta); err != nil {
				return nil, fmt.Errorf("invalid frontmatter: %w", err)
			}
			// Skip past the closing delimiter and any trailing newlines
			body = rest[idx+4:]
			// Trim leading newline after closing ---
			body = bytes.TrimPrefix(body, []byte("\n"))
		}
	}

	// Split on horizontal rules
	slides := splitSlides(body)

	// If we have a title in frontmatter, prepend a title slide
	if meta.Title != "" {
		titleSlide := generateTitleSlide(meta)
		slides = append([]Slide{titleSlide}, slides...)
	}

	if len(slides) == 0 {
		return nil, fmt.Errorf("no slides found in document")
	}

	return &Document{
		Meta:   meta,
		Slides: slides,
	}, nil
}

// splitSlides splits content on horizontal rule delimiters.
func splitSlides(content []byte) []Slide {
	parts := delimiterPattern.Split(string(content), -1)

	var slides []Slide
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		slides = append(slides, Slide{Content: trimmed})
	}

	return slides
}

// generateTitleSlide creates a centered title slide from metadata.
func generateTitleSlide(meta Metadata) Slide {
	var content strings.Builder
	content.WriteString("# ")
	content.WriteString(meta.Title)
	if meta.Author != "" {
		content.WriteString("\n\n")
		content.WriteString("*")
		content.WriteString(meta.Author)
		content.WriteString("*")
	}
	return Slide{Content: content.String()}
}
