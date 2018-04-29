package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gobuffalo/buffalo"
	jose "gopkg.in/square/go-jose.v2"
)

// Authenticate checks for a valid uuid to determine if a user is logged in or not
func Authenticate(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		// do some work before calling the next handler
		err := checkJwt(c.Response(), c.Request())
		if err == nil {
			err := next(c)
			// do some work after calling the next handler
			return err
		}

		return err
	}
}

func checkJwt(w http.ResponseWriter, r *http.Request) error {
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JwksURI}, nil)
	audience := Auth0APIAudience

	configuration := auth0.NewConfiguration(client, audience, Auth0APIIssuer, jose.RS256)
	validator := auth0.NewValidator(configuration, nil)

	_, err := validator.ValidateRequest(r)

	if err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Missing or invalid token")
		return err
	}

	return nil
}
