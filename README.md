# PokeDex CLI

An interactive command-line Pokedex built in Go. Explore Pokemon locations, encounter and catch Pokemon using different Pokeballs, and build your personal Pokedex — all from the terminal.

Data is sourced from the public [PokéAPI](https://pokeapi.co/api/v2) with an in-memory cache to avoid redundant requests.

## Requirements

- Go 1.25.5 or later

## Build & Run

```bash
go build -o pokedexcli
./pokedexcli
```

The REPL starts with a `Pokedex >` prompt. Type `help` to see all available commands.

## Commands

| Command | Description |
|---------|-------------|
| `help` | List all available commands |
| `exit` | Exit the application |
| `map` | Show the next page of Pokemon location areas |
| `mapb` | Show the previous page of Pokemon location areas |
| `explore <location>` | List all Pokemon that can be encountered in a location |
| `catch <pokemon> [ball]` | Attempt to catch a Pokemon |
| `inspect <pokemon>` | View stats for a Pokemon you've caught |
| `pokedex` | List all Pokemon in your Pokedex |

### Catching Pokemon

Catch probability is based on the official Pokemon capture formula, using each Pokemon's base capture rate. You can optionally specify a Pokeball type to improve your odds:

| Ball | Alias | Multiplier |
|------|-------|------------|
| (default) | — | 1.0x |
| `greatball` | `great` | 1.5x |
| `ultraball` | `ultra` | 2.0x |
| `masterball` | `master` | guaranteed |

Example:

```
Pokedex > catch pikachu ultraball
```

## Tests

```bash
go test ./...
```

## Project Structure

```
pokedexcli/
├── main.go                   # Entry point and client setup
├── repl.go                   # REPL loop and command dispatcher
├── command_*.go              # Command implementations
├── catch_test.go             # Catch probability unit tests
├── repl_test.go              # Input parsing unit tests
└── internal/
    ├── pokeapi/              # PokéAPI HTTP client
    └── pokecache/            # In-memory TTL cache
```

> Caught Pokemon are stored in memory for the current session only.
