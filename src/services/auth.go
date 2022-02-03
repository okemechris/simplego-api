package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"simplegoapi/src/errors"
	"strings"

	"github.com/Nerzal/gocloak/v10"
	_ "github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

var (
	clientId     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	realm        = os.Getenv("REALM")
	hostname     = os.Getenv("HOST")
)

var client gocloak.GoCloak

func InitializeOauthServer() {
	client = gocloak.NewClient(hostname)
}

func Protect(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if len(authHeader) < 1 {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(errors.UnauthorizedError())
			return
		}

		accessToken := strings.Split(authHeader, " ")[1]

		rptResult, err := client.RetrospectToken(r.Context(), accessToken, clientId, clientSecret, realm)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errors.BadRequestError(err.Error()))
			return
		}

		isTokenValid := *rptResult.Active

		if !isTokenValid {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(errors.UnauthorizedError())
			return
		}

		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}
