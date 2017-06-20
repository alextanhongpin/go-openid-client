package authsvc

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

// SetupRouter initialize the route for the authoriztion service
func SetupRouter(r *httprouter.Router) {
	r.GET("/authorize", getAuthorizeEndpoint)
	r.GET("/authorize/callback", getAuthorizeCallbackEndpoint)
	r.POST("/authorize", postAuthorizeEndpoint)
}

// getAuthorizeEndpoint displays the authorization endpoint and gives user an option to connect to third party provider
func getAuthorizeEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Load the connect screen
}

// postAuthorizeEndpoint will redirect the user to the consent screen
func postAuthorizeEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Make a call to the endpoint to get the authorization code
	http.Redirect(w, r, "http://localhost:8080/authorize?response_type=code&scope=openid profile email&client_id=&state=&redirect_uri=http://locahost:4000/authorize/callback", http.StatusPermanentRedirect)
	return
}

// getAuthorizeCallbackEndpoint will receive the response access token from the endpoint
func getAuthorizeCallbackEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	code := r.FormValue("code")
	state := r.FormValue("state")
	// Get the authorization code and exchange it with the access token
	// Make a call to the token endpoint to get the access token
	resp, err := http.PostForm("http://localhost:8080/token", url.Values{"code": {code}, "state": {state}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(resp)
	// Carry out type assertion and get the access token
	// Load a view, inject it to local storage and redirect immediately
	return
}
