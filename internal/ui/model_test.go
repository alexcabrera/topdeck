package ui

import (
	"testing"

	"github.com/alexcabrera/topdeck/internal/parser"
	tea "github.com/charmbracelet/bubbletea"
)

func TestNew(t *testing.T) {
	doc := &parser.Document{
		Slides: []parser.Slide{
			{Content: "# Slide 1"},
			{Content: "# Slide 2"},
		},
	}

	m := New(doc)
	if m.Current() != 0 {
		t.Errorf("expected current to be 0, got %d", m.Current())
	}
}

func TestNavigation_Next(t *testing.T) {
	doc := &parser.Document{
		Slides: []parser.Slide{
			{Content: "# Slide 1"},
			{Content: "# Slide 2"},
			{Content: "# Slide 3"},
		},
	}

	m := New(doc)

	// Navigate forward
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRight})
	m = updated.(Model)
	if m.Current() != 1 {
		t.Errorf("expected current to be 1, got %d", m.Current())
	}

	// Navigate forward again
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyRight})
	m = updated.(Model)
	if m.Current() != 2 {
		t.Errorf("expected current to be 2, got %d", m.Current())
	}

	// Should not go past the end
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyRight})
	m = updated.(Model)
	if m.Current() != 2 {
		t.Errorf("expected current to remain 2, got %d", m.Current())
	}
}

func TestNavigation_Prev(t *testing.T) {
	doc := &parser.Document{
		Slides: []parser.Slide{
			{Content: "# Slide 1"},
			{Content: "# Slide 2"},
		},
	}

	m := New(doc)

	// Move to second slide
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRight})
	m = updated.(Model)

	// Navigate back
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyLeft})
	m = updated.(Model)
	if m.Current() != 0 {
		t.Errorf("expected current to be 0, got %d", m.Current())
	}

	// Should not go before the start
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyLeft})
	m = updated.(Model)
	if m.Current() != 0 {
		t.Errorf("expected current to remain 0, got %d", m.Current())
	}
}

func TestNavigation_HomeEnd(t *testing.T) {
	doc := &parser.Document{
		Slides: []parser.Slide{
			{Content: "# Slide 1"},
			{Content: "# Slide 2"},
			{Content: "# Slide 3"},
			{Content: "# Slide 4"},
		},
	}

	m := New(doc)

	// Go to end
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnd})
	m = updated.(Model)
	if m.Current() != 3 {
		t.Errorf("expected current to be 3, got %d", m.Current())
	}

	// Go to start
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyHome})
	m = updated.(Model)
	if m.Current() != 0 {
		t.Errorf("expected current to be 0, got %d", m.Current())
	}
}

func TestNavigation_VimKeys(t *testing.T) {
	doc := &parser.Document{
		Slides: []parser.Slide{
			{Content: "# Slide 1"},
			{Content: "# Slide 2"},
		},
	}

	m := New(doc)

	// Test 'l' for next
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'l'}})
	m = updated.(Model)
	if m.Current() != 1 {
		t.Errorf("expected current to be 1 after 'l', got %d", m.Current())
	}

	// Test 'h' for previous
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'h'}})
	m = updated.(Model)
	if m.Current() != 0 {
		t.Errorf("expected current to be 0 after 'h', got %d", m.Current())
	}
}

func TestWindowResize(t *testing.T) {
	doc := &parser.Document{
		Slides: []parser.Slide{{Content: "# Test"}},
	}

	m := New(doc)

	// Simulate window resize
	updated, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = updated.(Model)

	if m.width != 80 {
		t.Errorf("expected width 80, got %d", m.width)
	}
	if m.height != 24 {
		t.Errorf("expected height 24, got %d", m.height)
	}
}
