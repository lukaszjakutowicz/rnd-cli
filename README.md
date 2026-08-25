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
| `uuid` | Generate a random UUID (v4) | `--items` (default `1`, max `20`) |
| `string` | Generate a random string | `--length` (default `16`), `--lowercase` - only lowercase characters, `--uppercase` - only uppercase characters, `--items` (default `1`, max `20`) |
| `integer` | Generate a random integer in a range | `--min` (default `0`) - minimal value,<br> `--max` (default `100`) - maximum value, `--items` (default `1`, max `20`) - how many items should be generated |
| `password` | Generate a random password | `--length` (default `16`),<br>`--mixed-letters` - return only letters [a-zA-Z],<br>`--only-digits` - return only digits,<br>`--no-special` - generate without special characters,<br>`--items` (default `1`, max `20`) - how many items should be generated |


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

Generate a password (16 mixed-case letters, digits, and special characters by default):

```sh
$ rnd-cli password
:E(#KtA8U,p3;a%B
```

Generate a digits-only PIN of a given length:

```sh
$ rnd-cli password --length 6 --only-digits
073139
```

Generate a password without special characters:

```sh
$ rnd-cli password --no-special
WI3aMkPKU2Nx8rIm
```

Generate multiple values at once with `--items`:

```sh
$ rnd-cli string --items 3 --length 8
nffIPokV
DnwwLqAV
HkHkAKfb
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

MIT License
