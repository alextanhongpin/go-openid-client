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
	// 	 POST /token HTTP/1.1
	//   Host: server.example.com
	//   Content-Type: application/x-www-form-urlencoded
	//   Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW

	//   grant_type=authorization_code&code=SplxlOBeZQQYbYS6WxSbIA
	//     &redirect_uri=https%3A%2F%2Fclient.example.org%2Fcb
	fmt.Println(resp)
	// Carry out type assertion and get the access token
	// Load a view, inject it to local storage and redirect immediately
	return
}

//   HTTP/1.1 200 OK
//   Content-Type: application/json
//   Cache-Control: no-store
//   Pragma: no-cache

//   {
//    "access_token": "SlAV32hkKG",
//    "token_type": "Bearer",
//    "refresh_token": "8xLOxBtZp8",
//    "expires_in": 3600,
//    "id_token": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFlOWdkazcifQ.ewogImlzc
//      yI6ICJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwKICJzdWIiOiAiMjQ4Mjg5
//      NzYxMDAxIiwKICJhdWQiOiAiczZCaGRSa3F0MyIsCiAibm9uY2UiOiAibi0wUzZ
//      fV3pBMk1qIiwKICJleHAiOiAxMzExMjgxOTcwLAogImlhdCI6IDEzMTEyODA5Nz
//      AKfQ.ggW8hZ1EuVLuxNuuIJKX_V8a_OMXzR0EHR9R6jgdqrOOF4daGU96Sr_P6q
//      Jp6IcmD3HP99Obi1PRs-cwh3LO-p146waJ8IhehcwL7F09JdijmBqkvPeB2T9CJ
//      NqeGpe-gccMg4vfKjkM8FcGvnzZUN4_KSP0aAp1tOJ1zZwgjxqGByKHiOtX7Tpd
//      QyHE5lcMiKPXfEIQILVq0pc_E2DzL7emopWoaoZTF_m0_N0YzFC6g6EJbOEoRoS
//      K5hoDalrcvRYLSrQAZZKflyuVCyixEoV9GfNQC3_osjzw2PAithfubEEBLuVVk4
//      XUVrWOLrLl0nx7RkKU8NXNHq-rvKMzqg"
//   }

//   HTTP/1.1 400 Bad Request
//   Content-Type: application/json
//   Cache-Control: no-store
//   Pragma: no-cache

//   {
//    "error": "invalid_request"
//   }
