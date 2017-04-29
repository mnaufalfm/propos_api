package main

import (
	"net/http"

	"io"

	"./models"
)

type UserHandler int

func (u UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//io.WriteString(res,req.RequestURI)
	io.WriteString(res, user.UserController(req.RequestURI, res, req))
}

func main() {
	var pengg UserHandler

	mux := http.NewServeMux()
	mux.Handle("/user/", pengg)

	http.ListenAndServe("localhost:9000", mux)
}
