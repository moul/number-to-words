# number-to-words

[![GoDoc](https://godoc.org/github.com/moul/number-to-words?status.svg)](https://godoc.org/github.com/moul/number-to-words)
[![Coverage Status](https://coveralls.io/repos/github/moul/number-to-words/badge.svg?branch=master)](https://coveralls.io/github/moul/number-to-words?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moul/number-to-words)](https://goreportcard.com/report/github.com/moul/number-to-words)

## CLI usage

```console
$ number-to-words --lang=fr 42
quarante-deux

$ number-to-words --lang=it 42
quarantadue

$ number-to-words --lang=en 42
forty-two

$ number-to-words --lang=romain 42
XLII

$ number-to-words 42
forty-two
```

```console
$ number-to-words -h
NAME:
   number-to-words - number to number

USAGE:
   number-to-words [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR(S):
   Manfred Touron <https://github.com/moul/number-to-words>

COMMANDS:
GLOBAL OPTIONS:
   --lang value, -l value   Set language (default: "en") [$NTW_LANGUAGE]
   --help, -h               show help
   --version, -v            print the version
```

## API usage

```go
import "github.com/moul/number-to-words"

fmt.Println(ntw.IntegerToFrench(42))
// Outputs: quarante-deux

fmt.Println(ntw.IntegerToEnglish(42))
// Outputs: forty-two

fmt.Println(ntw.IntegerToItalian(42))
// Outputs: quarantadue

fmt.Println(ntw.IntegerToRoman(42))
// Outputs: XLII
```

## Install

#### Using Golang

1. install and configure go on your host
2. get and build: `go get github.com/moul/number-to-words/cmd/number-to-words`
3. profit: `$GOPATH/bin/number-to-words 42`

#### Using Homebrew

coming soon

#### Using Docker

1. install and configure docker on your host
2. profit: `docker run --rm moul/number-to-words 42`

## License

MIT
