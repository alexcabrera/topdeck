package parser

import (
	"testing"
)

func TestParse_FrontmatterAndSlides(t *testing.T) {
	input := `---
title: Test Presentation
author: Test Author
---

# First Slide

Some content here.

---

# Second Slide

More content.
`
	doc, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if doc.Meta.Title != "Test Presentation" {
		t.Errorf("expected title 'Test Presentation', got %q", doc.Meta.Title)
	}
	if doc.Meta.Author != "Test Author" {
		t.Errorf("expected author 'Test Author', got %q", doc.Meta.Author)
	}

	// Should have 3 slides: title slide + 2 content slides
	if len(doc.Slides) != 3 {
		t.Fatalf("expected 3 slides, got %d", len(doc.Slides))
	}

	// First slide should be the generated title slide
	if doc.Slides[0].Content != "# Test Presentation\n\n*Test Author*" {
		t.Errorf("unexpected title slide content: %q", doc.Slides[0].Content)
	}
}

func TestParse_NoFrontmatter(t *testing.T) {
	input := `# First Slide

Content here.

---

# Second Slide

More content.
`
	doc, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if doc.Meta.Title != "" {
		t.Errorf("expected empty title, got %q", doc.Meta.Title)
	}

	// Should have 2 slides (no title slide without frontmatter)
	if len(doc.Slides) != 2 {
		t.Fatalf("expected 2 slides, got %d", len(doc.Slides))
	}
}

func TestParse_SingleSlide(t *testing.T) {
	input := `# Only Slide

Some content.
`
	doc, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(doc.Slides) != 1 {
		t.Fatalf("expected 1 slide, got %d", len(doc.Slides))
	}
}

func TestParse_TitleOnlyFrontmatter(t *testing.T) {
	input := `---
title: Just Title
---

# Content
`
	doc, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Title slide should not have author line
	if doc.Slides[0].Content != "# Just Title" {
		t.Errorf("unexpected title slide content: %q", doc.Slides[0].Content)
	}
}

func TestParse_EmptyDocument(t *testing.T) {
	input := ``
	_, err := Parse([]byte(input))
	if err == nil {
		t.Error("expected error for empty document, got nil")
	}
}

func TestParse_MultipleDashes(t *testing.T) {
	input := `# Slide 1

-----

# Slide 2

----------

# Slide 3
`
	doc, err := Parse([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(doc.Slides) != 3 {
		t.Fatalf("expected 3 slides, got %d", len(doc.Slides))
	}
}
