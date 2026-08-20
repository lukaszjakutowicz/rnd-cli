# rnd-cli

A small command-line tool for generating random things — UUIDs today, with random strings and passwords planned.

## Install

```sh
go install github.com/<owner>/rnd-cli@latest
```

Or build from source:

```sh
git clone https://github.com/<owner>/rnd-cli.git
cd rnd-cli
go build -o rnd-cli .
```

## Usage

```sh
rnd-cli <command> [flags]
```

Or, without building, run directly with `go run`:

```sh
go run . uuid
```

### Commands

| Command | Description |
|---|---|
| `uuid` | Generate a random UUID (v4) |

## Examples

Generate a UUID:

```sh
$ rnd-cli uuid
7d8c34e1-7527-45f9-9224-795190c5668b
```

Generate a UUID for use in a script:

```sh
$ ID=$(rnd-cli uuid)
$ echo "Created resource $ID"
Created resource 0f5cdf6b-92cf-4eba-bf61-9bd33ab821d9
```

Show available commands:

```sh
$ rnd-cli --help
```

## Development

```sh
go build ./...   # build
go vet ./...     # vet
go test ./...    # test
```

## License

TBD
