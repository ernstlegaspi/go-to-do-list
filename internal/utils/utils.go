package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func HasJWT(r *http.Request) (jwt.MapClaims, error) {
	cookie, cookieError := r.Cookie("session_token")

	if cookieError != nil {
		fmt.Println(cookieError)
		fmt.Println("Unauthorized.")
		return nil, cookieError
	}

	token, tokenError := ParseJWT(cookie.Value)

	if tokenError != nil {
		fmt.Println(tokenError)
		fmt.Println("Invalid token")
		return nil, cookieError
	}

	if !token.Valid {
		fmt.Println("Unathorized")
		return nil, nil
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

func CreateJWT(id int, name string) (string, error) {
	expiration := time.Second * time.Duration(3600*24*7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        id,
		"name":      name,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		fmt.Println(err)
		fmt.Println("Token error")
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func SetCookies(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     "session_token",
		HttpOnly: true,
		Value:    token,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
}
