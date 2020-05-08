package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println("simple HTTP server that logs anything you throw at it ")
		flag.PrintDefaults()
	}
	addr := flag.String("a", ":8347", "address to listen on")
	flag.Parse()

	err := http.ListenAndServe(
		*addr,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			buf, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			fmt.Printf(
				`
--------------------------------------------------------------------------------
path:    %s
method:  %s
headers: %+v
body:    %s`,
				r.URL.Path,
				r.Method,
				r.Header,
				string(buf),
			)
		}),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
