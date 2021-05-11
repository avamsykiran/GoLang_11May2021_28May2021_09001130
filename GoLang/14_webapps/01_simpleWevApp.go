package main

import "fmt"
import "net/http"

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "<h1>This is a simpel golang web</h1>")
	})

	//set up the server and lsitin at a specific port 4545
	err := http.ListenAndServe(":4545", nil)
	if err != nil {
		fmt.Println(err)
	}
}
