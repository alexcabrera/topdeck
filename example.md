---
title: Kitchen Sink Demo
author: Topdeck Team
---

# Headings

## Level 2 Heading

### Level 3 Heading

#### Level 4 Heading

##### Level 5 Heading

###### Level 6 Heading

---

# Text Formatting

This is **bold text** and this is *italic text*.

You can also do ***bold and italic*** together.

Here's some `inline code` in a sentence.

And ~~strikethrough~~ for deleted content.

---

# Unordered Lists

- First item
- Second item
  - Nested item
  - Another nested item
    - Even deeper
- Back to top level

---

# Ordered Lists

1. First step
2. Second step
   1. Sub-step A
   2. Sub-step B
3. Third step
4. Fourth step

---

# Mixed Lists

1. Install dependencies
   - Go 1.21+
   - Git
2. Clone the repo
   - Use SSH if you have keys
   - Use HTTPS otherwise
3. Build and run

---

# Code Blocks

## Go

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    name := os.Getenv("USER")
    if name == "" {
        name = "World"
    }
    fmt.Printf("Hello, %s!\n", name)
}
```

---

# More Code

## Python

```python
def fibonacci(n: int) -> list[int]:
    """Generate Fibonacci sequence."""
    if n <= 0:
        return []
    elif n == 1:
        return [0]
    
    fib = [0, 1]
    for _ in range(2, n):
        fib.append(fib[-1] + fib[-2])
    return fib

print(fibonacci(10))
```

---

# Even More Code

## JavaScript

```javascript
const fetchUser = async (id) => {
  try {
    const response = await fetch(`/api/users/${id}`);
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error('Failed to fetch user:', error);
    return null;
  }
};
```

---

# Shell Commands

```bash
# Clone and build topdeck
git clone https://github.com/alexcabrera/topdeck.git
cd topdeck
go build .

# Run a presentation
./topdeck examples/full-example.md
```

---

# Tables

| Language   | Typing     | Paradigm      | Year |
|------------|------------|---------------|------|
| Go         | Static     | Imperative    | 2009 |
| Python     | Dynamic    | Multi         | 1991 |
| Rust       | Static     | Multi         | 2010 |
| JavaScript | Dynamic    | Multi         | 1995 |
| Haskell    | Static     | Functional    | 1990 |

---

# Blockquotes

> "Simplicity is the ultimate sophistication."
>
> — Leonardo da Vinci

---

# Nested Blockquotes

> This is a blockquote.
>
> > This is nested inside the blockquote.
> >
> > It can have multiple paragraphs.
>
> Back to the outer level.

---

# Links

Here are some useful links:

- [Go Documentation](https://go.dev/doc/)
- [Charm Libraries](https://charm.sh/)
- [Bubbletea](https://github.com/charmbracelet/bubbletea)
- [Glamour](https://github.com/charmbracelet/glamour)

---

# Task Lists

- [x] Parse markdown files
- [x] Extract frontmatter
- [x] Split on horizontal rules
- [x] Render with Glamour
- [x] Style with Lipgloss
- [x] Keyboard navigation
- [ ] Add image support
- [ ] Add speaker notes

---

# Horizontal Rules in Content

You can use horizontal rules within slides for visual separation.

* * *

Like this one above (using `* * *`).

___

Or this one (using `___`).

Note: Only `---` on its own line creates a new slide!

---

# Emoji Support

Topdeck supports emoji! :rocket: :sparkles: :tada:

Some favorites:
- :thumbsup: Thumbs up
- :heart: Heart
- :fire: Fire
- :star: Star
- :coffee: Coffee

---

# Long Code Block

```go
// A more complex example with structs and methods
type Presentation struct {
    Title    string
    Author   string
    Slides   []Slide
    Current  int
}

type Slide struct {
    Content string
    Notes   string
}

func (p *Presentation) Next() bool {
    if p.Current < len(p.Slides)-1 {
        p.Current++
        return true
    }
    return false
}

func (p *Presentation) Prev() bool {
    if p.Current > 0 {
        p.Current--
        return true
    }
    return false
}

func (p *Presentation) Progress() string {
    return fmt.Sprintf("%d / %d", p.Current+1, len(p.Slides))
}
```

---

# Definition Lists

Term 1
: Definition for term 1

Term 2
: Definition for term 2
: Can have multiple definitions

Complex Term
: A longer definition that explains
  the concept in more detail.

---

# Math-like Content

While we don't render LaTeX, you can still show formulas:

```
E = mc²

a² + b² = c²

∑(i=1 to n) i = n(n+1)/2
```

---

# ASCII Art

```
    ╔═══════════════════════════════╗
    ║     Welcome to Topdeck!       ║
    ║                               ║
    ║   ┌─────────┐                 ║
    ║   │  Slide  │ ──► Terminal    ║
    ║   └─────────┘                 ║
    ║                               ║
    ╚═══════════════════════════════╝
```

---

# Long Paragraph

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

---

# Multiple Paragraphs

First paragraph with some introductory text that sets up the context for what follows.

Second paragraph that builds on the first one and provides more details about the topic at hand.

Third paragraph that wraps things up and provides a conclusion or summary of the key points.

---

# Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `→` `l` `n` `Space` | Next slide |
| `←` `h` `p` | Previous slide |
| `g` `Home` | First slide |
| `G` `End` | Last slide |
| `q` `Esc` | Quit |

---

# Thank You!

That's all folks!

Press `q` to exit.

---

# Hidden Bonus Slide

You found the secret slide! :eyes:

```
  _____ _                 _        
 |_   _| |__   __ _ _ __ | | _____ 
   | | | '_ \ / _` | '_ \| |/ / __|
   | | | | | | (_| | | | |   <\__ \
   |_| |_| |_|\__,_|_| |_|_|\_\___/
                                   
```
