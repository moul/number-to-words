package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kulaginds/number-to-words/pkg/ntw"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/number-to-words"
	app.Version = ntw.Version
	app.Usage = "number to number web API"

	defaultListen := ":" + os.Getenv("PORT")
	if defaultListen == ":" {
		defaultListen = ":8000"
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind, b",
			Usage: "HTTP bind address",
			Value: defaultListen,
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

		language := ntw.Languages.Lookup(vars["lang"])
		if language == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such language %q\n", vars["lang"])
			return
		}

		output := language.IntegerToWords(number)
		if output == "" {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "number not supported for this language\n")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", output)
	})
	log.Printf("Listening to %s", c.String("bind"))
	return http.ListenAndServe(c.String("bind"), r)
}
