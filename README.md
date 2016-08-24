# number-to-words

[![GoDoc](https://godoc.org/github.com/moul/number-to-words?status.svg)](https://godoc.org/github.com/moul/number-to-words)
[![Coverage Status](https://coveralls.io/repos/github/moul/number-to-words/badge.svg?branch=master)](https://coveralls.io/github/moul/number-to-words?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moul/number-to-words)](https://goreportcard.com/report/github.com/moul/number-to-words)

Convert numbers to words.

Converti les nombres en lettres.

## Supported languages / Langues supportées

* English [en]
* Français [fr]
* Italiano [it]
* Swedish [se]
* Roman numbers


## CLI usage

```console
$ number-to-words --lang=fr 42
quarante-deux

$ number-to-words --lang=it 42
quarantadue

$ number-to-words --lang=en 42
forty-two

$ number-to-words --lang=se 42
fyrtio-två

$ number-to-words --lang=romain 42
XLII

$ number-to-words 42
forty-two

$ number-to-words --lang=all 42
forty-two
quarante-deux
quarantadue
fyrtio-två
XLII

$ number-to-words --lang=all 1
one
un
uno
en
I

$ number-to-words --lang=all 1337
one thousand three hundred thirty-seven
mille trois cent trente-sept
uno mille trecento trentasette
en tusen tre hundra trettio-sju
MCCCXXXVII

$ number-to-words --lang=all 1234567890
one billion two hundred thirty-four million five hundred sixty-seven thousand eight hundred ninety
un milliard deux cent trente-quatre millions cinq cent soixante-sept mille huit cent quatre-vingt-dix
uno miliardo duecento trentaquattro milione cinquecento sessentasette mille ottocento novanta
en miljarder två hundra trettio-fyra miljoner fem hundra sextio-sju tusen åtta hundra nittio

$ number-to-words --lang=all 1000000000000
one quadrillion
un billiard
uno quadrillion
en biljoner
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

fmt.Println(ntw.IntegerToSwedish(42))
// Outputs: fyrtio-två

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
