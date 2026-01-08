// Package ui provides the bubbletea TUI for presenting slides.
package ui

import (
	"fmt"
	"strings"

	"github.com/alexcabrera/topdeck/internal/parser"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// Model is the bubbletea model for the presentation.
type Model struct {
	doc     *parser.Document
	current int
	width   int
	height  int
	ready   bool

	// Cached rendering - avoids creating renderer on every View()
	renderedSlides []string
}

// New creates a new presentation model.
func New(doc *parser.Document) Model {
	return Model{
		doc:     doc,
		current: 0,
	}
}

// Init initializes the model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "right", "l", "n", " ", "enter":
			if m.current < len(m.doc.Slides)-1 {
				m.current++
			}
		case "left", "h", "p", "backspace":
			if m.current > 0 {
				m.current--
			}
		case "g", "home":
			m.current = 0
		case "G", "end":
			m.current = len(m.doc.Slides) - 1
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		widthChanged := m.width != msg.Width
		m.width = msg.Width
		m.height = msg.Height

		if !m.ready || widthChanged {
			m.renderedSlides = m.renderAllSlides()
			m.ready = true
		}
	}
	return m, nil
}

// contentWidth returns the available width for slide content.
func (m Model) contentWidth() int {
	w := m.width - 8 // 4 padding on each side
	if w < 20 {
		return 20
	}
	return w
}

// renderAllSlides pre-renders all slides and returns them.
func (m Model) renderAllSlides() []string {
	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(m.contentWidth()),
		glamour.WithEmoji(),
	)
	if err != nil {
		// Fallback: use raw content
		slides := make([]string, len(m.doc.Slides))
		for i, slide := range m.doc.Slides {
			slides[i] = slide.Content
		}
		return slides
	}

	slides := make([]string, len(m.doc.Slides))
	for i, slide := range m.doc.Slides {
		rendered, err := renderer.Render(slide.Content)
		if err != nil {
			slides[i] = slide.Content
		} else {
			slides[i] = strings.TrimRight(rendered, "\n")
		}
	}
	return slides
}

// View renders the current slide.
func (m Model) View() string {
	if !m.ready {
		return "Loading..."
	}

	slide := m.renderSlide()
	statusBar := m.renderStatusBar()

	return lipgloss.JoinVertical(lipgloss.Left, slide, statusBar)
}

// renderSlide applies styling to the cached slide content.
func (m Model) renderSlide() string {
	if m.current >= len(m.renderedSlides) {
		return ""
	}

	content := m.renderedSlides[m.current]
	contentHeight := m.height - 1 // reserve 1 line for status bar

	styled := slideStyle.
		Width(m.width).
		Height(contentHeight).
		Render(content)

	return styled
}

// renderStatusBar renders the status bar at the bottom.
func (m Model) renderStatusBar() string {
	title := m.doc.Meta.Title
	if title == "" {
		title = "Presentation"
	}

	counter := fmt.Sprintf(" %d / %d ", m.current+1, len(m.doc.Slides))

	titleWidth := lipgloss.Width(title)
	counterWidth := lipgloss.Width(counter)
	fillWidth := m.width - titleWidth - counterWidth - 1
	if fillWidth < 0 {
		fillWidth = 0
	}
	fill := strings.Repeat(" ", fillWidth)

	bar := " " + title + fill + counter

	return statusBarStyle.Width(m.width).Render(bar)
}

// Current returns the current slide index (for testing).
func (m Model) Current() int {
	return m.current
}
