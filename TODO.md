# TODO

Findings from a project review (2026-11-14).

## Code (`sumdir.go`)

- [x] **Errors go to stdout and exit 0.** `sumdir -dir /nope` prints the error and exits with status 0. Print to stderr and `os.Exit(1)`.
- [x] **Extensionless files get a bare `:` header.** e.g. `Makefile`. Use a label such as `(none)`.
- [x] **Dotfiles are grouped by their name.** `filepath.Ext(".bashrc")` is `.bashrc`, so `-hidden` puts it under `bashrc:`. Strip the leading dot before taking the extension.
- [x] **Symlinks to directories are listed as files.** `DirEntry.IsDir()` doesn't follow symlinks; `os.Stat` entries with `ModeSymlink`.
- [x] **Extensions are case-sensitive.** `B.GO` and `a.go` land in separate groups. Lowercase the extension.
- [x] Use `slices.Sorted(maps.Keys(...))` instead of the manual key loop, avoid shadowing `files`, add a package doc comment (revive `package-comments`).
- [x] **No tests.** Add table tests for `listFilesByExtension` using `t.TempDir()`.

## Tooling

- [x] **`.golangci.yml` is invalid for v2** (`golangci-lint config verify` fails): `version` must be the string `"2"`; `disable-all` → `default: none`; `linters-settings` → `linters.settings` / `formatters.settings`; `issues.exclude-dirs` → `linters.exclusions.paths`.
- [x] **`.goreleaser.yaml` fails `goreleaser check`**: add `version: 2`; `snapshot.name_template` → `version_template`; drop the no-op `go generate` hook; zip archives for Windows.
- [x] **CI**: `go.mod` says `go 1.21` but CI tests 1.23–1.25 (bump to 1.23); checkout before `setup-go` so caching works; add a lint job; bump action versions.
- [x] **README** is one line: add install, usage, and example output.
