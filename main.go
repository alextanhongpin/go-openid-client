package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/alextanhongpin/go-openid-client/app"
)

func main() {
	var (
		port = flag.Int("port", 4000, "The server port")
	)

	templates := app.NewTemplate()
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("hello")
		t := templates["home"]
		fmt.Println(t)
		t.Execute(w, nil)
	})


	fmt.Printf("\nListening to port *:%d. Press ctrl + c to cancel.", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
