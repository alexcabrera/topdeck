# AGENTS.md

This document provides instructions and context for AI coding agents working in this repository.

---

## Immutable Section

> **⚠️ WARNING: This section may NEVER be modified without explicit user request.**

### Core Principles

1. **Preferred Libraries Directory (`./.preferred-libraries/`)**
   - This directory is **READ-ONLY**
   - Consult during planning/reasoning phases as reference material
   - Use as inspiration for implementation patterns
   - **NEVER** modify, fix, or alter any files in this directory

2. **Language & Standards**
   - All code must be written in **idiomatic Go**
   - Follow Go best practices (effective Go, code review comments, proverbs)
   - Use standard library where possible before reaching for third-party dependencies

3. **Testing Requirements**
   - Every feature **MUST** have a comprehensive test suite
   - A request is **NEVER** complete unless the entire end-to-end test suite passes
   - Tests must cover happy paths, edge cases, and error conditions

4. **Documentation Requirements**
   - `./README.md` - Keep up-to-date with comprehensive end-to-end single-file documentation covering full functionality and usage
   - `./AGENTS.md` - Keep up-to-date with coding decisions and architectural choices

5. **Git Workflow**
   - Branch aggressively for new features/changes
   - Use verbose, descriptive commit messages
   - Merge back to the originating branch upon completion
   - **Always** inspect git history (`git log`, `git blame`, `git diff`) when gathering context for any request

6. **Semantic Versioning (SemVer)**
   - All version numbers **MUST** follow the `x.y.z` format (MAJOR.MINOR.PATCH)
   - The agent **MUST** assume responsibility for updating:
     - `z` (PATCH): Increment for bug fixes and backwards-compatible changes
     - `y` (MINOR): Increment for new features that are backwards-compatible (resets `z` to 0)
   - The agent **MUST NEVER** update `x` (MAJOR) unless explicitly directed by the user
   - Major version changes indicate breaking changes and require explicit user authorization

---

## Mutable Section

This section contains evolving decisions, patterns, and conventions. Update as the codebase grows.

### Project Overview

- **Name**: topdeck
- **Language**: Go
- **Description**: Terminal presentation tool inspired by presenterm
- **Philosophy**: Convention over configuration - beautiful defaults, zero config
- **Status**: Ready for release (v0.1.0)
- **Version**: 0.1.0

### Directory Structure

```
topdeck/
├── .goreleaser.yml         # Release automation
├── .preferred-libraries/   # READ-ONLY reference libraries
├── AGENTS.md               # This file - agent instructions
├── CHANGELOG.md            # Version history
├── LICENSE                 # MIT license
├── README.md               # Project documentation
├── go.mod                  # Go module definition
├── main.go                 # CLI entry point (fang + cobra)
├── internal/
│   ├── parser/             # Markdown/frontmatter parsing
│   │   ├── parser.go       # Parse function, Document/Slide types
│   │   └── parser_test.go
│   └── ui/                 # Bubbletea TUI
│       ├── model.go        # Model, Init, Update, View
│       ├── styles.go       # Lipgloss style definitions
│       └── model_test.go
├── examples/
│   └── full-example.md     # Kitchen sink demo
└── testdata/
    └── example.md          # Test presentation
```

### Commands

| Command | Description |
|---------|-------------|
| `go build ./...` | Build all packages |
| `go test ./...` | Run all tests |
| `go test -v ./...` | Run tests with verbose output |
| `go test -race ./...` | Run tests with race detector |
| `go vet ./...` | Run static analysis |
| `go fmt ./...` | Format all code |
| `go mod tidy` | Clean up dependencies |

### Code Conventions

- Follow standard Go formatting (`gofmt`)
- Use meaningful variable and function names
- Keep functions small and focused
- Handle all errors explicitly
- Use table-driven tests where appropriate
- Document exported functions and types
- **NEVER** use emojis in code, comments, commit messages, or documentation

### Architectural Decisions

1. **Convention over Configuration**: No flags, no config files, no themes. One beautiful default that adapts to terminal light/dark mode.

2. **Parser Design**: Manual frontmatter extraction followed by regex-based slide splitting on `---`. Simpler than using goldmark extensions.

3. **Title Slide Generation**: If frontmatter contains a title, automatically prepend a generated title slide with centered title and author.

4. **Rendering Pipeline**: Slide content -> Glamour (markdown to ANSI) -> Lipgloss (padding/layout) -> Status bar composition.

5. **State Management**: Bubbletea model holds current slide index, terminal dimensions. Pre-renders all slides on startup/resize for instant navigation.

6. **Performance**: Glamour renderer created once per resize, all slides pre-rendered and cached. View() only applies lipgloss styling to cached content.

7. **Styling**: Uses `lipgloss.AdaptiveColor` throughout for automatic light/dark mode support. No user-configurable themes.

### Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/charmbracelet/bubbletea` | TUI framework |
| `github.com/charmbracelet/glamour` | Markdown rendering |
| `github.com/charmbracelet/lipgloss` | Terminal styling |
| `github.com/charmbracelet/fang` | CLI framework |
| `github.com/spf13/cobra` | Command parsing |
| `gopkg.in/yaml.v3` | Frontmatter parsing |

---

## Preferred Libraries Guide

The `.preferred-libraries/` directory contains reference implementations for libraries this project should use. This section documents when and how to use each library.

### Libraries

#### fantasy (`charm.land/fantasy`)

**Purpose**: Build AI agents with Go. Multi-provider, multi-model, one API.

**When to use**:
- Building AI-powered features that need LLM integration
- Implementing tool-calling agents
- Supporting multiple AI providers (OpenAI, Anthropic, Google, Azure, Bedrock, OpenRouter)
- Structured output generation from LLMs

**Key patterns**:
```go
import (
    "charm.land/fantasy"
    "charm.land/fantasy/providers/openrouter"
)

// 1. Create a provider
provider, err := openrouter.New(openrouter.WithAPIKey(apiKey))

// 2. Get a language model
model, err := provider.LanguageModel(ctx, "model-name")

// 3. Define tools with typed inputs
type QueryInput struct {
    Field string `json:"field" description:"Field description"`
}
tool := fantasy.NewAgentTool("tool_name", "Tool description", handlerFunc)

// 4. Create and run agent
agent := fantasy.NewAgent(model,
    fantasy.WithSystemPrompt("System prompt"),
    fantasy.WithTools(tool),
)
result, err := agent.Generate(ctx, fantasy.AgentCall{Prompt: prompt})
```

**Available providers**: `providers/anthropic`, `providers/openai`, `providers/google`, `providers/azure`, `providers/bedrock`, `providers/openrouter`, `providers/openaicompat`

**Reference**: `.preferred-libraries/fantasy/` (see `examples/` for usage patterns)

---

#### bubbletea (`github.com/charmbracelet/bubbletea`)

**Purpose**: Build terminal user interfaces using The Elm Architecture.

**When to use**:
- Building interactive terminal applications
- Creating TUIs with keyboard/mouse input
- Full-screen or inline terminal interfaces
- Applications requiring real-time updates

**Key patterns**:
```go
import tea "github.com/charmbracelet/bubbletea"

// 1. Define model (application state)
type model struct {
    cursor   int
    choices  []string
    selected map[int]struct{}
}

// 2. Implement tea.Model interface
func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "up", "k":
            m.cursor--
        case "down", "j":
            m.cursor++
        }
    }
    return m, nil
}

func (m model) View() string {
    return "Render your UI as a string"
}

// 3. Run the program
p := tea.NewProgram(initialModel())
_, err := p.Run()
```

**Companion libraries**:
- [Bubbles](https://github.com/charmbracelet/bubbles) - Common UI components (text inputs, spinners, lists, tables)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling and layout

**Reference**: `.preferred-libraries/bubbletea/` (see `tutorials/` and `examples/` for patterns)

---

#### cobra (`github.com/spf13/cobra`)

**Purpose**: Create powerful CLI applications with subcommands, flags, and completions.

**When to use**:
- Building CLI tools with subcommands (`app server`, `app fetch`)
- Applications requiring POSIX-compliant flags
- CLIs needing shell completions (bash, zsh, fish, powershell)
- Generating man pages or documentation

**Key patterns**:
```go
import "github.com/spf13/cobra"

// Define root command
var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "A brief description",
    Long:  "A longer description",
}

// Add subcommands
var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Start the server",
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation
    },
}

func init() {
    rootCmd.AddCommand(serveCmd)
    serveCmd.Flags().IntP("port", "p", 8080, "Port to listen on")
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

**Reference**: `.preferred-libraries/cobra/` (see `site/content/user_guide.md` for documentation)

---

#### fang (`github.com/charmbracelet/fang`)

**Purpose**: CLI starter kit with batteries-included Cobra enhancements.

**When to use**:
- Building CLIs that need styled help/usage output
- Applications requiring automatic `--version` flags
- CLIs that should generate manpages
- Wanting consistent, attractive error formatting

**Key patterns**:
```go
import (
    "github.com/charmbracelet/fang"
    "github.com/spf13/cobra"
)

func main() {
    cmd := &cobra.Command{
        Use:   "myapp",
        Short: "A styled CLI application",
    }
    if err := fang.Execute(context.Background(), cmd); err != nil {
        os.Exit(1)
    }
}
```

**Features**:
- Styled help and usage pages
- Styled error output
- Automatic `--version` flag (from build info or custom)
- Hidden `man` command for manpage generation
- `completion` command for shell completions
- Themeable output

**When to choose fang over raw cobra**: Use fang when you want attractive, styled CLI output with minimal setup. Use raw cobra for maximum control or when output styling is not needed.

**Reference**: `.preferred-libraries/fang/`

---

#### log (`github.com/charmbracelet/log`)

**Purpose**: Minimal, colorful, structured logging.

**When to use**:
- Application logging with colored output
- Structured logging with key-value pairs
- When you need leveled logging (Debug, Info, Warn, Error, Fatal)
- Integration with `log/slog` or standard library `log`

**Key patterns**:
```go
import "github.com/charmbracelet/log"

// Global logger (timestamps enabled, info level default)
log.Info("Starting application")
log.Debug("Debug info")  // Won't print unless level set
log.Error("Operation failed", "err", err, "user", userID)

// Create custom logger
logger := log.NewWithOptions(os.Stderr, log.Options{
    ReportCaller:    true,
    ReportTimestamp: true,
    TimeFormat:      time.Kitchen,
    Prefix:          "myapp ",
    Level:           log.DebugLevel,
})

// Sub-loggers with context
reqLogger := logger.With("request_id", reqID)
reqLogger.Info("Processing request")

// Formatters: TextFormatter (default), JSONFormatter, LogfmtFormatter
logger.SetFormatter(log.JSONFormatter)

// Use as slog handler
slogger := slog.New(logger)
```

**Reference**: `.preferred-libraries/log/` (see `examples/` for patterns)

---

### Usage Examples (Not Libraries)

#### catwalk

**Purpose**: A database/service for Crush-compatible AI models. This is NOT a library to import but rather a reference for how model provider configurations are structured.

**When to consult**:
- Understanding how AI provider configurations are structured
- Adding new AI provider support
- Debugging model compatibility issues

**Reference**: `.preferred-libraries/catwalk/` (see `internal/providers/configs/` for provider JSON schemas)

---

### Library Selection Decision Tree

```
Need AI/LLM integration?
  -> Use fantasy

Need interactive terminal UI?
  -> Use bubbletea (+ bubbles for components)

Need CLI with subcommands?
  -> Need styled output? -> Use fang (wraps cobra)
  -> Need maximum control? -> Use cobra directly

Need logging?
  -> Use log (charmbracelet/log)
```

---

## Agent Workflow

When working on a request:

1. **Gather Context**
   - Read this file (`AGENTS.md`)
   - Check `README.md` for current functionality
   - Inspect git history for relevant context
   - Consult `.preferred-libraries/` for implementation patterns (read-only)

2. **Plan**
   - Understand the full scope of the request
   - Identify all files that need modification
   - Design the test strategy

3. **Implement**
   - Create a feature branch
   - Write tests first or alongside implementation
   - Make atomic commits with verbose messages
   - Follow Go idioms and project conventions

4. **Verify**
   - Run `go build ./...` - must pass
   - Run `go test ./...` - must pass
   - Run `go vet ./...` - must pass
   - Verify no regressions

5. **Document**
   - Update `README.md` with new functionality
   - Update `AGENTS.md` mutable section with decisions
   - Write clear commit message

6. **Complete**
   - Merge branch back to origin branch
   - Confirm all tests still pass
