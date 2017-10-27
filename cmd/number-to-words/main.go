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
		cli.BoolFlag{
			Name:   "unicode, u",
			EnvVar: "NTW_UNICODE",
			Usage:  "Use unicode characters when available",
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
	var outputs []string
	lang := c.String("lang")

	if lang == "en" || lang == "english" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToEnglish(input))
		found = true
	}
	if lang == "fr" || lang == "fr-fr" || lang == "french" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToFrench(input))
		found = true
	}
	if lang == "fr-be" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToFrBe(input))
		found = true
	}
	if lang == "it" || lang == "italian" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToItalian(input))
		found = true
	}
	if lang == "es" || lang == "spanish" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToSpanish(input))
		found = true
	}
	if lang == "se" || lang == "swedish" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToSwedish(input))
		found = true
	}
	if lang == "nl" || lang == "dutch" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToDutch(input))
		found = true
	}
	if lang == "tr" || lang == "turkish" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToTurkish(input))
		found = true
	}
	if lang == "pt-pt" || lang == "portuguesePT" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToPortuguesePT(input))
		found = true
	}
	if lang == "roman" || lang == "all" {
		if c.Bool("unicode") {
			outputs = append(outputs, ntw.IntegerToUnicodeRoman(input))
		} else {
			outputs = append(outputs, ntw.IntegerToRoman(input))
		}
		found = true
	}
	if lang == "aegean" || lang == "all" {
		outputs = append(outputs, ntw.IntegerToAegean(input))
		found = true
	}
	if !found {
		fmt.Fprintf(os.Stderr, "Unknown language: %s\n", lang)
		os.Exit(1)
	}

	fmt.Println(strings.Join(outputs, "\n"))
	return nil
}
