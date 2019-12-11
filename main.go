package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"github.com/koding/websocketproxy"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
	%s [-l [host]:port]

Options:
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	var (
		listenhp string
	)
	flag.StringVar(&listenhp, "l", ":13000", "listen host and port")

	flag.Parse()

	if strings.Count(listenhp, ":") > 0 {
		if listenhp[len(listenhp)-1] == ':' {
			listenhp = listenhp + "13000"
		}
	} else {
		listenhp = listenhp + ":13000"
	}

	u, err := url.Parse("ws://localhost:3000/")
	if err != nil {
		log.Fatalln(err)
	}

	err = http.ListenAndServe(listenhp, websocketproxy.NewProxy(u))
	if err != nil {
		log.Fatalln(err)
	}
}
