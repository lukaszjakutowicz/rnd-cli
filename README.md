# rnd-cli

A small command-line tool for generating random things: UUIDs, strings, integers, and passwords.

## Install

```sh
go install github.com/lukaszjakutowicz/rnd-cli@latest
```

Or build from source:

```sh
git clone https://github.com/lukaszjakutowicz/rnd-cli.git
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

| Command | Description | Flags |
|---|---|---|
| `uuid` | Generate a random UUID (v4) | — |
| `string` | Generate a random string | `--length` (default `16`), `--lowercase`, `--uppercase` |
| `integer` | Generate a random integer in a range | `--min` (default `0`), `--max` (default `100`) |
| `password` | Generate a random password | `--length` (default `16`), `--mixed-letters`, `--only-digits` |

`string` draws from mixed-case letters by default; `--lowercase`/`--uppercase` narrow it to one case (mutually exclusive). `password` draws from mixed-case letters and digits by default; `--mixed-letters` narrows it to letters only, `--only-digits` to digits only (also mutually exclusive).

## Examples

Generate a UUID:

```sh
$ rnd-cli uuid
5cf07879-f13b-41a4-b943-12d1ed70ff9c
```

Generate a random string (16 mixed-case letters by default):

```sh
$ rnd-cli string
wdlXapGqQfnloxHK
```

Generate a lowercase-only string of a given length:

```sh
$ rnd-cli string --length 12 --lowercase
njndcphognik
```

Generate a random integer between 0 and 100 (the default range):

```sh
$ rnd-cli integer
19
```

Generate a random integer in a custom range, e.g. a dice roll:

```sh
$ rnd-cli integer --min 1 --max 6
2
```

Generate a password (16 mixed-case letters and digits by default):

```sh
$ rnd-cli password
tkmO4MWFE8FGk9aJ
```

Generate a digits-only PIN of a given length:

```sh
$ rnd-cli password --length 6 --only-digits
073139
```

Use a generated value in a script:

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
