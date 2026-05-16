# ascii-color

A command-line tool that renders text as ASCII art with optional terminal color support. Built on top of the ascii-art project, with the ability to color an entire string or target a specific substring.

---

## Features

- Renders any printable ASCII text as banner art
- Supports 7 terminal colors via ANSI escape codes
- Color the entire output or a specific substring only
- Handles multi-line input using `\n`

---

## Usage

```bash
go run . [--color=COLOR] [SUBSTRING] "TEXT"
```

| Argument | Required | Description |
|---|---|---|
| `--color=COLOR` | No | Color to apply |
| `SUBSTRING` | No | Part of TEXT to color |
| `TEXT` | Yes | The string to render as ASCII art |

### Examples

```bash
# No color
go run . "Hello"

# Color the entire string
go run . --color=red "Hello"

# Color only a substring
go run . --color=green "He" "Hello"

# Multi-line input
go run . --color=blue "Hello\nWorld"
```

---

## Supported Colors

| Name | Preview |
|---|---|
| `red` | `\033[31m` |
| `green` | `\033[32m` |
| `yellow` | `\033[33m` |
| `blue` | `\033[34m` |
| `magenta` | `\033[35m` |
| `cyan` | `\033[36m` |
| `white` | `\033[37m` |

---

## Project Structure

```
ascii-color/
├── main.go            # Entry point, argument parsing
├── build.go           # Orchestrates rendering across lines
├── render.go          # Renders each line with color logic
├── color.go           # ANSI color map and Colorize function
├── banner.go          # Loads and parses the banner file
├── splitvalidate.go   # Input splitting and validation
└── standard.txt       # ASCII art banner font file
```

---

## How It Works

1. `main.go` parses `--color`, the optional substring, and the input text from `os.Args`
2. `BuildArt` splits the input on `\n` and processes each line
3. `RenderLine` renders each character's 8-row banner representation, wrapping characters that fall within the substring range with ANSI color codes
4. `Colorize` applies the opening color code and closes with a reset `\033[0m`

---

## Requirements

- Go 1.18+
- A terminal that supports ANSI escape codes
- `standard.txt` banner file in the project root

---

## Author

Built as an extension of the ascii-art project, adding color support through ANSI escape codes and substring-aware rendering.
