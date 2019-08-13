package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/urfave/cli"
	ntw "moul.io/number-to-words"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://moul.io/number-to-words"
	app.Version = ntw.Version
	app.Usage = "number to number"
	// FIXME: enable autocomplete

	// append language list to help
	cli.AppHelpTemplate = fmt.Sprintf("%s\nAVAILABLE LANGUAGES:\n", cli.AppHelpTemplate)
	for _, lang := range ntw.Languages {
		cli.AppHelpTemplate = fmt.Sprintf("%s   %s\n", cli.AppHelpTemplate, lang.HelpText())
	}
	cli.AppHelpTemplate += "\n"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang, l",
			EnvVar: "NTW_LANGUAGE",
			Usage:  "Set language",
			Value:  "en",
		},
	}
	app.Action = convert
	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}

func convert(c *cli.Context) error {
	if len(c.Args()) != 1 {
		return fmt.Errorf("usage: number-to-words [--lang=<lang>] <number>")
	}

	inputStr := c.Args()[0]
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		return err
	}

	var outputs []string
	lang := c.String("lang")

	if lang == "all" {
		for _, language := range ntw.Languages {
			outputs = append(outputs, language.IntegerToWords(input))
		}
	} else {
		language := ntw.Languages.Lookup(lang)
		if language == nil {
			return fmt.Errorf("unknown language: %s", lang)
		}
		outputs = append(outputs, language.IntegerToWords(input))
	}

	fmt.Println(strings.Join(outputs, "\n"))
	return nil
}
