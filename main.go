// Package main is the entry point for the topdeck CLI.
package main

import (
	"context"
	"os"

	"github.com/alexcabrera/topdeck/internal/parser"
	"github.com/alexcabrera/topdeck/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "topdeck <file>",
		Short: "Terminal presentations",
		Long:  "A beautiful terminal presentation tool. Just markdown.",
		Example: `  topdeck slides.md
  topdeck presentation.md`,
		Args: cobra.ExactArgs(1),
		RunE: run,
	}

	if err := fang.Execute(context.Background(), cmd); err != nil {
		os.Exit(1)
	}
}

func run(_ *cobra.Command, args []string) error {
	content, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}

	doc, err := parser.Parse(content)
	if err != nil {
		return err
	}

	p := tea.NewProgram(
		ui.New(doc),
		tea.WithAltScreen(),
	)

	_, err = p.Run()
	return err
}
