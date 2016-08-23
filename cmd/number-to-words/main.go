package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/moul/number-to-words"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/number-to-words"
	app.Version = ntw.Version
	app.Usage = "number to number"
	// FIXME: enable autocomplete

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang, l",
			EnvVar: "NTW_LANGUAGE",
			Usage:  "Set language",
			Value:  "en",
		},
	}
	app.Action = convert
	app.Run(os.Args)
}

func convert(c *cli.Context) error {
	if len(c.Args()) != 1 {
		return fmt.Errorf("usage: number-to-words <number>")
	}

	inputStr := c.Args()[0]
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		return err
	}

	found := false
	var output []string
	lang := c.String("lang")

	if lang == "en" || lang == "english" || lang == "all" {
		output = append(output, ntw.IntegerToEnglish(input))
		found = true
	}
	if lang == "fr" || lang == "french" || lang == "all" {
		output = append(output, ntw.IntegerToFrench(input))
		found = true
	}
	if lang == "it" || lang == "italian" || lang == "all" {
		output = append(output, ntw.IntegerToItalian(input))
		found = true
	}
	if lang == "roman" || lang == "all" {
		output = append(output, ntw.IntegerToRoman(input))
		found = true
	}
	if !found {
		fmt.Fprintf(os.Stderr, "Unknown language: %s\n", lang)
		os.Exit(1)
	}

	fmt.Println(strings.Join(output, "\n"))
	return nil
}
