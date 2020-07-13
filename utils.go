package main

import (
	"bytes"
	"fmt"
	"net/http"
)

/*
	==================================================================================
	THIS FILE WAS COPIED FROM https://github.com/tonyghita/graphql-go-example
	==================================================================================
*/

func errorJSON(msg string) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `{"error": "%s"}`, msg)
	return buf.Bytes()
}

func respond(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}
