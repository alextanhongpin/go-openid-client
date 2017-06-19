package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	var (
		port = flag.Int("port", 4000, "The server port")
	)

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "hello")
	})

	fmt.Printf("\nListening to port *:%d. Press ctrl + c to cancel.", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
