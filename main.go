package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(
		":8347",
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
