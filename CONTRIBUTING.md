# Contributing

## Development

```sh
go build ./...   # build
go vet ./...     # vet
go test ./...    # test
```

Run `gofmt -l .` before committing and fix any files it lists.

## Adding a new generator command

- Core generation logic goes in `internal/generator/<name>.go`, kept free of CLI concerns so it can be unit tested directly.
- Wire it up as a subcommand in `cmd/<name>.go`, registered via `register(command{...})` in an `init()` function.
- Add tests in `internal/generator/<name>_test.go` covering the normal case, edge cases, and error cases.
- Use `crypto/rand` for anything security-sensitive (passwords, tokens) — never `math/rand`.

## Commits and branches

- Branch from `develop`.
- Keep commit messages short and descriptive (e.g. `Add password command`).
- Open a pull request into `develop`; CI must pass (build, vet, test) before merging.
