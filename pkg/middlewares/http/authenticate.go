package http

import (
	"errors"
	"fmt"
	"github.com/EugeneNail/motivatr-lib-common/pkg/authentication"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
)

func Authenticate(jwtSalt string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			parts := strings.Split(request.Header.Get("Authorization"), " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims := jwt.RegisteredClaims{}
			_, err := jwt.ParseWithClaims(parts[1], &claims, func(token *jwt.Token) (any, error) {
				return []byte(jwtSalt), nil
			})

			if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
				writer.WriteHeader(http.StatusUnauthorized)
				if _, writeErr := writer.Write([]byte("token expired")); writeErr != nil {
					http.Error(writer, fmt.Sprintf("writing data: %v (original error: %v", writeErr, err), http.StatusInternalServerError)
				}
				return
			}

			if err != nil {
				err = fmt.Errorf("parsing a token: %w", err)
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			stringUserId, err := claims.GetSubject()
			if err != nil {
				err = fmt.Errorf("extracting user's id: %w", err)
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			userId, err := strconv.ParseInt(stringUserId, 10, 64)
			if err != nil {
				err = fmt.Errorf("converting user id to int64: %w", err)
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			next.ServeHTTP(writer, request.WithContext(authentication.InjectHttpUserId(userId, request.Context())))
		}
	}
}
