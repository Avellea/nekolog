# nekolog

A small, dependency-free logging package for Go providing leveled console output with optional ANSI coloring.

## Features

- Leveled logs: `Trace`, `Debug`, `Info`, `Warn`, `Error`, `Fatal`.
- Formatted variants: `Tracef`, `Debugf`, `Infof`, `Warnf`, `Errorf`, `Fatalf`.
- Thread-safe (mutex-protected).
- Timestamps included in log lines.
- Lightweight and easy to import.

## Installation

Add the module to your project (Go modules enabled):

    import "github.com/avellea/nekolog"

Or, to explicitly fetch the module:

    go get github.com/avellea/nekolog

## Usage

Create a logger and use the provided methods:

    // example (in a .go file)
    import "github.com/avellea/nekolog"

    func main() {
        log := nekolog.NewLogger()

        log.Info("Application starting...")
        log.Infof("Listening on port %d", 8080)
        log.Warn("Cache is almost full")
        log.Error("failed to read config")
    }

Methods are safe for concurrent use.

## Project layout

- `log.go` — core logger implementation and public API.
- `internal/utils` — helper utilities (time formatting, colors, level types).
- `cmd/main.go` — minimal example demonstrating usage.

## Contributing

Contributions are welcome. Keep changes small and focused; add tests or examples when relevant.

## License

Choose and add a license file as appropriate for your project.
