package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/widadfjry/cashier-app/app"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/web"
	"net/http"
)

var jwtSecret = []byte("sjhdjhihdoahssjcbabduiaghdwuiuah928319038jasfhji1289y39jkashdkj")
var publicAPIs = []string{
	"/api/v1/users/login",
	"/api/v1/users/register",
	"/api/v1/users/verification",
	"/api/v1/users/new-password",
}

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	url := request.URL.Path
	fmt.Printf("\n")

	for _, publicAPI := range publicAPIs {
		if publicAPI == url {
			middleware.Handler.ServeHTTP(writer, request)
			return
		}
	}

	token := request.Header.Get("Authorization")
	if token == "" {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponseError := web.ResponseError{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:  "token not defined",
		}

		helper.WriteToResponseBody(writer, webResponseError)
		return
	}
	claims, err := verifyToken(request.Header.Get("Authorization"))
	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponseError := web.ResponseError{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:  err.Error(),
		}

		helper.WriteToResponseBody(writer, webResponseError)
		return
	}

	db := app.NewDB()

	var blockedTokens []string
	SQL := "SELECT token FROM blockedusers"
	rows, err := db.Query(SQL)
	defer rows.Close()

	for rows.Next() {
		var blockedToken string
		err = rows.Scan(&blockedToken)
		helper.PanicIfError(err)

		blockedTokens = append(blockedTokens, blockedToken)
	}

	for _, encryptedToken := range blockedTokens {
		decryptedToken := helper.DecryptData(encryptedToken)
		if decryptedToken == token {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)

			webResponseError := web.ResponseError{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Error:  "you're logged out",
			}

			helper.WriteToResponseBody(writer, webResponseError)
			return
		}
	}

	SQL = "SELECT username FROM users WHERE username = ?"
	username := claims["username"]

	rows, err = db.Query(SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	var userInDB string
	if rows.Next() {
		err := rows.Scan(&userInDB)
		if err != nil {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusInternalServerError)

			webResponseError := web.ResponseError{
				Code:   http.StatusInternalServerError,
				Status: "INTERNAL SERVER ERROR",
				Error:  err.Error(),
			}

			helper.WriteToResponseBody(writer, webResponseError)
			return
		}
	}

	SQL = "SELECT role FROM users WHERE username = ?"
	rows, err = db.Query(SQL, username)
	helper.PanicIfError(err)

	var role string
	if rows.Next() {
		err = rows.Scan(&role)
		helper.PanicIfError(err)
	}

	if role == "super_admin" {
		request.Header.Set("isAdmin", "true")
	} else {
		request.Header.Set("isAdmin", "false")
	}

	if userInDB != username {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponseError := web.ResponseError{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:  "unregistered",
		}
		helper.WriteToResponseBody(writer, webResponseError)
	} else {
		request.Header.Set("username", username.(string))
		middleware.Handler.ServeHTTP(writer, request)
	}
}

func verifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
