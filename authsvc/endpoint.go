package authsvc

import "github.com/julienschmidt/httprouter"
import "net/http"

// SetupRouter initialize the route for the authoriztion service
func SetupRouter(r *httprouter.Router) {
	r.GET("/authorize", getAuthorizeEndpoint)
	r.GET("/authorize/callback", getAuthorizeCallbackEndpoint)
	r.POST("/authorize", postAuthorizeEndpoint)
}

// getAuthorizeEndpoint displays the authorization endpoint and gives user an option to connect to third party provider
func getAuthorizeEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return
}

// postAuthorizeEndpoint will submit an oauth request to oauth server
func postAuthorizeEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Make a call to the endpoint to get the authorization code
	return
}

// getAuthorizeCallbackEndpoint will receive the response access token from the endpoint
func getAuthorizeCallbackEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get the authorization code and exchange it with the access token
	// Make a call to the token endpoint to get the access token
	return
}
