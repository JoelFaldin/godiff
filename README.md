# godiff

A very simple [Go](https://go.dev/) based CLI tool. Alternative to `git diff`.

<div align="center">
  <img width="558" height="315" alt="image" src="https://github.com/user-attachments/assets/2715981c-3146-4c3e-965d-904f7d1a9559" />
</div>

All features subject to change.

## Project Structure

```
godiff/
├── cmd/
|   ├── root.go      # Base command
|   ├── run.go       # Main subcommand
├── internal/        # App internal logic
|   ├── errors/
|   ├── parser/
|   ├── renderer/
|   ├── runner/
├── .gitignore
├── go.mod
├── main.go
```

## Prerequisites

- Go 1.21 or higher [(download here)](https://go.dev/dl/)
- Git installed on your machine

## Installation

1. Clone the repo on your machine.
2. Run `go install .`
3. Navigate to any repository.
4. Run `godiff run .` for the whole repo or `godiff run README.md` for an specific file.
5. You can also add the `--staged` flag to see diff on files already added with `git add`.
6. Example: `godiff run --staged .\cmd\root.go`.

## Thanks for visiting!
