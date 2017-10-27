package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/moul/number-to-words"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/number-to-words"
	app.Version = ntw.Version
	app.Usage = "number to number web API"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind, b",
			Usage: "HTTP bind address",
			Value: ":8000",
		},
	}
	app.Action = server
	app.Run(os.Args)
}

func server(c *cli.Context) error {
	r := mux.NewRouter()
	r.HandleFunc("/{lang}/{number}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		number, err := strconv.Atoi(vars["number"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid input number %q: %v\n", vars["number"], err)
			return
		}

		var output string
		switch vars["lang"] {
		case "en", "english":
			output = ntw.IntegerToEnglish(number)
		case "fr", "fr-fr", "french":
			output = ntw.IntegerToFrench(number)
		case "fr-be":
			output = ntw.IntegerToFrBe(number)
		case "it", "italian":
			output = ntw.IntegerToItalian(number)
		case "es", "spanish":
			output = ntw.IntegerToSpanish(number)
		case "se", "swedish":
			output = ntw.IntegerToSwedish(number)
		case "nl", "dutch":
			output = ntw.IntegerToDutch(number)
		case "tr", "turkish":
			output = ntw.IntegerToTurkish(number)
		case "pt-pt", "portuguesePT":
			output = ntw.IntegerToPortuguesePT(number)
		case "roman":
			output = ntw.IntegerToRoman(number)
		case "roman-unicode":
			output = ntw.IntegerToUnicodeRoman(number)
		case "aegean":
			output = ntw.IntegerToAegean(number)
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such language %q\n", vars["lang"])
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", output)
	})
	log.Printf("Listening to %s", c.String("bind"))
	return http.ListenAndServe(c.String("bind"), r)
}
