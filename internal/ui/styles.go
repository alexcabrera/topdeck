package ui

import "github.com/charmbracelet/lipgloss"

// Color definitions using adaptive colors for light/dark mode support.
var (
	subtle   = lipgloss.AdaptiveColor{Light: "#666666", Dark: "#888888"}
	statusBg = lipgloss.AdaptiveColor{Light: "#E0E0E0", Dark: "#1A1A1A"}
)

// Style definitions.
var (
	slideStyle = lipgloss.NewStyle().
			Padding(2, 4)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(subtle).
			Background(statusBg)
)
