# number-to-words

[![GoDoc](https://godoc.org/github.com/moul/number-to-words?status.svg)](https://godoc.org/github.com/moul/number-to-words)
[![Build Status](https://travis-ci.org/moul/number-to-words.svg?branch=master)](https://travis-ci.org/moul/number-to-words)
[![Coverage Status](https://coveralls.io/repos/github/moul/number-to-words/badge.svg?branch=master)](https://coveralls.io/github/moul/number-to-words?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moul/number-to-words)](https://goreportcard.com/report/github.com/moul/number-to-words)

Convert numbers to words.

## Supported languages

* American English / United States of America [en] 🇺🇸
* Français / France [fr] 🇫🇷
* Français (Belge) / Belgium [fr-be] 🇧🇪
* Italiano / Italy [it] 🇮🇹
* Spanish / Spain [es] 🇪🇸
* Swedish / Sweden [se] 🇸🇪
* Dutch / Netherlands [nl] 🇳🇱
* Turkish / Turkey [tr] 🇹🇷
* Portuguese / Portugal [pt-pt] 🇵🇹
* Roman numbers Ⅷ  (with `--unicode` support)
* Aegean numerals


## CLI usage

```console
$ number-to-words --lang=fr 42
quarante-deux

$ number-to-words --lang=fr-be 92
nonante-deux

$ number-to-words --lang=it 42
quarantadue

$ number-to-words --lang=es 42
cuarenta y dos

$ number-to-words --lang=en 42
forty-two

$ number-to-words --lang=se 42
fyrtio-två

$ number-to-words --lang=nl 42
tweeenveertig

$ number-to-words --lang=tr 42
kırk iki

$ number-to-words --lang=pt-pt 42
quarenta e dois

$ number-to-words --lang=roman 42
XLII

$ number-to-words --lang=roman --unicode 42
ⅩⅬⅡ

$ number-to-words --lang=aegean 42
𐄓𐄈
```

default language is english

```console
$ number-to-words 42
forty-two
```

print every supported language at once

```console
$ number-to-words --lang=all 42
forty-two
quarante-deux
quarante-deux
quarantadue
cuarenta y dos
fyrtio-två
tweeenveertig
kırk iki
quarenta e dois
XLII
𐄓𐄈

$ number-to-words --lang=all 1
one
un
un
uno
uno
en
één
bir
um
I
𐄇

$ number-to-words --lang=all 1337
one thousand three hundred thirty-seven
mille trois cent trente-sept
mille trois cent trente-sept
uno mille trecento trentasette
un mil trescientos treinta y siete
en tusen tre hundra trettio-sju
éénduizend driehonderdzevenendertig
bin üç yüz otuz yedi
mil trezentos e trinta e sete
MCCCXXXVII
𐄢𐄛𐄒𐄍

$ number-to-words --lang=all 1234567890
one billion two hundred thirty-four million five hundred sixty-seven thousand eight hundred ninety
un milliard deux cent trente-quatre millions cinq cent soixante-sept mille huit cent quatre-vingt-dix
un milliard deux cent trente-quatre millions cinq cent soixante-sept mille huit cent nonante
uno miliardo duecento trentaquattro milione cinquecento sessentasette mille ottocento novanta
un mil millones doscientos treinta y cuatro millones quinientos sesenta y siete mil ochocientos noventa
en miljarder två hundra trettio-fyra miljoner fem hundra sextio-sju tusen åtta hundra nittio
één miljard tweehonderdvierendertig miljoen vijfhonderdzevenenzestigduizend achthonderdnegentig
bir milyar iki yüz otuz dört milyon beş yüz altmış yedi bin sekiz yüz doksan
mil milhões duzentos e trinta e quatro milhões quinhentos e sessenta e sete mil oitocentos e noventa

$ number-to-words --lang=all 1000000000000
one trillion
un billion
uno triliardo
un billón
en biljoner
één biljoen
bir trilyon
um bilião
too big number
too big number
```

### `--help`

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
   --unicode, -u            Use unicode characters when available [$NTW_UNICODE]
   --version, -v            print the version
```

### Unicode support

Roman support the `--unicode` option.

```console
$ for i in {1..20}; do ./number-to-words -l roman -u $i; done
Ⅰ
Ⅱ
Ⅲ
Ⅳ
Ⅴ
Ⅵ
Ⅶ
Ⅷ
Ⅸ
Ⅹ
Ⅺ
Ⅻ
ⅩⅢ
ⅩⅣ
ⅩⅤ
ⅩⅥ
ⅩⅦ
ⅩⅧ
ⅩⅨ
ⅩⅩ
```

## API usage

```go
import "github.com/moul/number-to-words"

fmt.Println(ntw.IntegerToFrench(42))
// Outputs: quarante-deux
```

```go
fmt.Println(ntw.IntegerToFrBe(92))
// Outputs: nonante-deux

fmt.Println(ntw.IntegerToEnglish(42))
// Outputs: forty-two

fmt.Println(ntw.IntegerToItalian(42))
// Outputs: quarantadue

fmt.Println(ntw.IntegerToSpanish(42))
// Outputs: cuarenta y dos

fmt.Println(ntw.IntegerToSwedish(42))
// Outputs: fyrtio-två

fmt.Println(ntw.IntegerToDutch(42))
// Outputs: tweeenveertig

fmt.Println(ntw.IntegerToTurkish(42))
// Outputs: kırk iki

fmt.Println(ntw.IntegerToPortuguesePT(42))
// Outputs: quarenta e dois

fmt.Println(ntw.IntegerToRoman(42))
// Outputs: XLII

fmt.Println(ntw.IntegerToUnicodeRoman(42))
// Outputs: ⅩⅬⅡ

fmt.Println(ntw.IntegerToAegean(42))
// Outputs: 𐄓𐄈
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
