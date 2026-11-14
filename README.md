# sumdir

Show files in a directory grouped by extension, written in Go.

## Install

```sh
go install github.com/mikemadden42/sumdir@latest
```

## Usage

```sh
sumdir [-dir path] [-hidden]
```

| Flag      | Default | Description                         |
|-----------|---------|-------------------------------------|
| `-dir`    | `.`     | The directory to list files from    |
| `-hidden` | `false` | Include hidden (dot) files          |

Subdirectories (including symlinks to directories) are skipped. Extensions
are compared case-insensitively, and files with no extension (such as
`Makefile` or `.bashrc`) are grouped under `(none)`.

## Example

```console
$ sumdir -dir ~/project -hidden
(none):
- .gitignore
- Makefile

go:
- main.go
- main_test.go

md:
- README.md
```
