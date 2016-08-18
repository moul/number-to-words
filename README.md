# number-to-words

[![GoDoc](https://godoc.org/github.com/moul/number-to-words?status.svg)](https://godoc.org/github.com/moul/number-to-words)

## CLI usage

```console
$ number-to-words --lang=fr 42
quarante-deux
$ number-to-words --lang=en 42
forty-two
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
   --debug, -D              Enable debug mode [$NTW_DEBUG]
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
```

## License

MIT
