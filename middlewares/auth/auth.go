package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type userNameKey struct{}

const tokenPrefix = "T"

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			t := req.Header.Get("Authorization")
			if t == "" {
				h.ServeHTTP(w, req)
				return
			}

			u, err := validate(t)
			if err != nil {
				log.Println(err)
				http.Error(w, `{"message": "invalid token"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(req.Context(), userNameKey{}, u)
			h.ServeHTTP(w, req.WithContext(ctx))
		})
}

func validate(t string) (string, error) {
	elements := strings.SplitN(t, "_", 2)
	if len(elements) < 2 {
		return "", errors.New("invalid token: elements length is less than 2")
	}

	tType, user := elements[0], elements[1]
	if tType != tokenPrefix {
		return "", errors.New(fmt.Sprintf("invalid token: token type is not T. tType is %s", tType))
	}

	return user, nil
}
