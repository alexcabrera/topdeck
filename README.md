# Topdeck

A beautiful terminal presentation tool. Just markdown.

## Installation

### Homebrew

```bash
brew tap alexcabrera/tap
brew install topdeck
```

### Go

```bash
go install github.com/alexcabrera/topdeck@latest
```

### From source

```bash
git clone https://github.com/alexcabrera/topdeck.git
cd topdeck
go build .
```

### Packages

Download a package for your system from the [releases page](https://github.com/alexcabrera/topdeck/releases).

## Usage

```bash
topdeck presentation.md
```

That's it. No flags, no configuration. Just present.

## Writing Presentations

Create a markdown file with slides separated by horizontal rules (`---`):

```markdown
---
title: My Presentation
author: Your Name
---

# First Slide

This is the first slide content.

---

## Second Slide

- Point one
- Point two  
- Point three

---

## Code Example

```go
func main() {
    fmt.Println("Hello, World!")
}
```

---

# The End

Thanks for watching!
```

### Frontmatter

Optional YAML frontmatter at the top of your file:

```yaml
---
title: Presentation Title
author: Author Name
---
```

If you include a title, topdeck automatically generates a title slide.

### Slide Delimiters

Slides are separated by horizontal rules (three or more dashes on their own line):

```markdown
# Slide 1

Content here.

---

# Slide 2

More content.
```

## Keyboard Controls

| Key | Action |
|-----|--------|
| `→` `l` `n` `Space` `Enter` | Next slide |
| `←` `h` `p` `Backspace` | Previous slide |
| `g` `Home` | First slide |
| `G` `End` | Last slide |
| `q` `Esc` `Ctrl+C` | Quit |

## Supported Markdown

Topdeck renders standard GitHub Flavored Markdown:

- **Headings** (`#`, `##`, etc.)
- **Emphasis** (`*italic*`, `**bold**`)
- **Lists** (ordered and unordered)
- **Code blocks** with syntax highlighting
- **Tables**
- **Blockquotes**
- **Links**
- **Inline code**
- **Emoji** (`:rocket:`, `:sparkles:`, etc.)

## Philosophy

Topdeck follows **convention over configuration**:

- One beautiful default theme that adapts to your terminal
- No configuration files
- No command-line flags (except `--help` and `--version`)
- Just markdown, just presentations

