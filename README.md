# number-to-words

:smile: `number-to-words` converts a number to words

[![CircleCI](https://circleci.com/gh/moul/number-to-words.svg?style=shield)](https://circleci.com/gh/moul/number-to-words)
[![GoDoc](https://godoc.org/moul.io/number-to-words?status.svg)](https://godoc.org/moul.io/number-to-words)
[![License](https://img.shields.io/github/license/moul/number-to-words.svg)](https://github.com/moul/number-to-words/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/moul/number-to-words.svg)](https://github.com/moul/number-to-words/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/number-to-words)](https://goreportcard.com/report/moul.io/number-to-words)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/number-to-words/badge)](https://www.codefactor.io/repository/github/moul/number-to-words)
[![codecov](https://codecov.io/gh/moul/number-to-words/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/number-to-words)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/number-to-words.svg)](https://microbadger.com/images/moul/number-to-words)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

## Supported languages



| Code            | Flag | Language                         | Main Region | 42               |
| --------------- | ---- | -------------------------------- | ----------- | ---------------- |
| `en`, `en-us`   | 🇺🇸   | American English                 | USA         | forty-two        |
| `fr`, `fr-fr`   | 🇫🇷   | French, Français                 | France      | quarante-deux    |
| `it`, `it-it`   | 🇮🇹   | Italiano                         | Italy       | quarantadue      |
| `es`, `es-es`   | 🇪🇸   | European Spanish                 | Spain       | cuarenta y dos   |
| `dk`, `da-dk`   | 🇩🇰   | Danish                          | Denmark      | toogfyrre       |
| `se`, `sv-se`   | 🇸🇪   | Swedish                          | Sweden      | fyrtio-två       |
| `nl`, `nl-nl`   | 🇳🇱   | Dutch                            | Netherlands | tweeenveertig    |
| `tr`, `tr-tr`   | 🇹🇷   | Turkish                          | Turkey      | kırk iki         |
| `pt`, `pt-pt`   | 🇵🇹   | Portuguese                       | Portugal    | quarenta e dois  |
| `pl`, `pl-pl`   | 🇵🇱   | Polish                           | Poland      | czterdzieści dwa |
| `ru`, `ru-ru`   | 🇷🇺   | Russian                          | Russia      | сорок два        |
| `ir`, `ir-ir`   | 🇮🇷   | Iranian                          | Iran        | چهل و دو         |
| `id`, `id-id`   | 🇮🇩   | Indonesian                       | Indonesia   | empat puluh dua  |
| `jp`, `ja-jp`   | 🇯🇵   | Japanese                         | Japan       | 四十二           |
| `fr-be`         | 🇧🇪   | Belgian French, Français (Belge) | Belgium     | quarante-deux    |
| `uk`, `uk-ua`   | 🇺🇦   | Ukrainian                        | Ukraine     | сорок два        |
| `roman`         |      | Roman Numbers                    |             | XLII             |
| `roman-unicode` |      | Roman (with Unicode)             |             | ⅩⅬⅡ              |
| `aegean`        |      | Aegean Numerals                  |             | 𐄓𐄈               |

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

$ number-to-words --lang=dk 42
toogfyrre

$ number-to-words --lang=se 42
fyrtio-två

$ number-to-words --lang=nl 42
tweeenveertig

$ number-to-words --lang=pl 42
czterdzieści dwa

$ number-to-words --lang=tr 42
kırk iki

$ number-to-words --lang=pt-pt 42
quarenta e dois

$ number-to-words --lang=roman 42
XLII

$ number-to-words --lang=roman-unicode
ⅩⅬⅡ

$ number-to-words --lang=aegean 42
𐄓𐄈

$ number-to-words --lang=ir 42
چهل و دو

$ number-to-words --lang=id 42
empat puluh dua
```

The default language is English

```console
$ number-to-words 42
forty-two
```

Print different numbers in every supported language

```console
$ number-to-words --lang=all 42
forty-two
quarante-deux
quarante-deux
quarantadue
cuarenta y dos
toogfyrre
fyrtio-två
tweeenveertig
kırk iki
quarenta e dois
XLII
𐄓𐄈
چهل و دو
empat puluh dua

$ number-to-words --lang=all 1
one
un
un
uno
uno
en
en
één
bir
um
I
𐄇
یک
satu

$ number-to-words --lang=all 1337
one thousand three hundred thirty-seven
mille trois cent trente-sept
mille trois cent trente-sept
uno mille trecento trentasette
un mil trescientos treinta y siete
et tusind tre hundrede og syvogtredive
en tusen tre hundra trettio-sju
éénduizend driehonderdzevenendertig
bin üç yüz otuz yedi
mil trezentos e trinta e sete
MCCCXXXVII
𐄢𐄛𐄒𐄍
یک هزار سیصد سی و هفت
seribu tiga ratus tiga puluh tujuh

$ number-to-words --lang=all 1234567890
one billion two hundred thirty-four million five hundred sixty-seven thousand eight hundred ninety
un milliard deux cent trente-quatre millions cinq cent soixante-sept mille huit cent quatre-vingt-dix
un milliard deux cent trente-quatre millions cinq cent soixante-sept mille huit cent nonante
uno miliardo duecento trentaquattro milione cinquecento sessentasette mille ottocento novanta
un mil millones doscientos treinta y cuatro millones quinientos sesenta y siete mil ochocientos noventa
en milliard to hundrede og firetredive millioner fem hundrede og syvogtres tusinde ottehundrede og halvfems
en miljarder två hundra trettio-fyra miljoner fem hundra sextio-sju tusen åtta hundra nittio
één miljard tweehonderdvierendertig miljoen vijfhonderdzevenenzestigduizend achthonderdnegentig
bir milyar iki yüz otuz dört milyon beş yüz altmış yedi bin sekiz yüz doksan
mil milhões duzentos e trinta e quatro milhões quinhentos e sessenta e sete mil oitocentos e noventa
یک میلیارد دویست سی و چهار میلیون پانصد شصد و هفت هزار هشتصد نود
satu milyar dua ratus tiga puluh empat juta lima ratus enam puluh tujuh ribu delapan ratus sembilan puluh

$ number-to-words --lang=all 1000000000000
one trillion
un billion
uno triliardo
un billón
en billion
en biljoner
één biljoen
bir trilyon
um bilião
too big number
too big number
satu triliun
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
   Manfred Touron <https://moul.io/number-to-words>

COMMANDS:
GLOBAL OPTIONS:
   --lang value, -l value   Set language (default: "en") [$NTW_LANGUAGE]
   --help, -h               show help
   --version, -v            print the version

AVAILABLE LANGUAGES:
   European Spanish (es, es-es, es_ES, spanish) 🇪🇸
   Belgian French (fr-be, fr_BE, belgian) 🇧🇪
   French (fr, fr-fr, fr_FR, french) 🇫🇷
   Italian (it, it-it, it_IT, italian) 🇮🇹
   Roman Numbers (with Unicode) (roman-unicode)
   Danish (da-dk, da_DK, danish) 🇩🇰
   Swedish (sv-se, sv_SE, swedish) 🇸🇪
   Aegean (aegean)
   American English (en, en-us, es_US, american, english) 🇺🇸  *default*
   Dutch (nl, dutch, nl-nl, nl_NL) 🇳🇱
   Portuguese (Portugal) (pt, pt-pt, pt_PT, portuguese) 🇵🇹
   Polish (Poland) (pl, pl-pl, pl_PL, polish) 🇵🇱
   Iranian (Iran) (ir, ir-ir, ir_IR, Iran) 🇮🇷
   Indonesian (Indonesia) (id, id-id, id_ID, indonesian) 🇮🇩
   Roman Numbers (roman)
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
import ntw "moul.io/number-to-words"

fmt.Println(ntw.IntegerToFrFr(42)) // french
// Outputs: quarante-deux
```

```go
fmt.Println(ntw.IntegerToFrBe(92)) // belgian french
// Outputs: nonante-deux

fmt.Println(ntw.IntegerToEnUs(42)) // american english
// Outputs: forty-two

fmt.Println(ntw.IntegerToItIt(42)) // italian
// Outputs: quarantadue

fmt.Println(ntw.IntegerToEsEs(42)) // spanish
// Outputs: cuarenta y dos

fmt.Println(ntw.IntegerToDaDk(42)) // danish
// Outputs: toogfyrre

fmt.Println(ntw.IntegerToSvSe(42)) // swedish
// Outputs: fyrtio-två

fmt.Println(ntw.IntegerToNlNl(42)) // dutch
// Outputs: tweeenveertig

fmt.Println(ntw.IntegerToPlPl(42)) // polish
// Outputs: czterdzieści dwa

fmt.Println(ntw.IntegerToTrTr(42)) // turkish
// Outputs: kırk iki

fmt.Println(ntw.IntegerToPtPt(42)) // portuguese (portugal)
// Outputs: quarenta e dois

fmt.Println(ntw.IntegerToRoman(42)) // roman
// Outputs: XLII

fmt.Println(ntw.IntegerToRomanUnicode(42)) // roman (unicode)
// Outputs: ⅩⅬⅡ

fmt.Println(ntw.IntegerToAegean(42)) // aegean (unicode)
// Outputs: 𐄓𐄈

fmt.Println(ntw.IntegerToIrIr(42)) // iranian
// Outputs: چهل و دو

fmt.Println(ntw.IntegerToIDID(42)) // indonesian
// Outputs: empat puluh dua
```

## Install

#### Using Golang

1. install and configure Go on your host
2. get and build: `go get moul.io/number-to-words/cmd/number-to-words`
3. profit: `$GOPATH/bin/number-to-words 42`

#### Using Homebrew

1. install Homebrew
2. install number-to-words: `brew install moul/moul/number-to-words`
3. profit: `number-to-words 42`

#### Using Docker

1. install and configure Docker on your host
2. profit: `docker run --rm moul/number-to-words 42`

## Test

1. profit: `make`

## License

MIT
