package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/Sirupsen/logrus"
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
			Name:   "debug, D",
			EnvVar: "NTW_DEBUG",
			Usage:  "Enable debug mode",
		},
	}
	app.Action = convert
	app.Run(os.Args)
}

func convert(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if len(c.Args()) != 1 {
		return fmt.Errorf("usage: number-to-words <number>")
	}

	inputStr := c.Args()[0]
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		return err
	}

	var output string
	switch lang := c.String("lang"); lang {
	case "en", "english":
		output = ntw.IntegerToEnglish(input)
		break
	case "fr", "french":
		output = ntw.IntegerToFrench(input)
		break
	default:
		fmt.Fprintf(os.Stderr, "Unknown language: %s\n", lang)
		os.Exit(1)
		break
	}

	fmt.Println(output)
	return nil
}
