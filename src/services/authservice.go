package services

import (
	"github.com/Nerzal/gocloak/v8"
	_ "github.com/gorilla/mux"
	"net/http"
	"strings"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

const clientId = "sidekiq-client"
const clientSecret = "350c66b6-6762-493c-837e-4675e33adf66"
const  realm = "Sidekiq"
const hostname = "https://keycloak.sidekik.ng"

var client gocloak.GoCloak

func InitKeycloakClient(){
	client = gocloak.NewClient(hostname)
}




func Protect(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessToken := strings.Split(r.Header.Get("Authorization")," ")[1]

		if len(accessToken) < 1 {
			w.WriteHeader(401)
			return
		}

		rptResult, err := client.RetrospectToken(r.Context(), accessToken, clientId, clientSecret, realm)

		if err != nil{
			panic(err)
		}


		isTokenValid := *rptResult.Active


		if !isTokenValid {

			w.WriteHeader(401)
			return
		}

		permissions := *rptResult.Permissions
		println(permissions)

		if len(role) > 0 {

		}


		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}

