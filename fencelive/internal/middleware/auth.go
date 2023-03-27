package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func VerifyJWT() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header["Authorization"])
		if len(r.Header["Authorization"]) < 1 {
			fmt.Println("No token provided")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("")), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		roles := claims["roles"]

		fmt.Printf("ROLES: %s\n", roles)

	})
}
// shit ahh code
